package servers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/galexrt/fivenet/internal/tests"
	"github.com/galexrt/fivenet/query"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"go.uber.org/zap"
)

var TestDBServer *dbServer

type dbServer struct {
	db       *sql.DB
	pool     *dockertest.Pool
	resource *dockertest.Resource

	stopped bool
}

func init() {
	TestDBServer = &dbServer{}
}

func (m *dbServer) Setup() {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	var err error
	m.pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %q", err)
	}

	// uses pool to try to connect to Docker
	err = m.pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %q", err)
	}

	// pulls an image, creates a container based on it and runs it
	m.resource, err = m.pool.RunWithOptions(
		&dockertest.RunOptions{
			Repository: "docker.io/library/mysql",
			Tag:        "8.1.0",
			Env: []string{
				"MYSQL_ROOT_PASSWORD=secret",
				"MYSQL_USER=fivenet",
				"MYSQL_PASSWORD=changeme",
				"MYSQL_DATABASE=fivenettest",
			},
			Cmd: []string{
				"mysqld",
				"--innodb-ft-min-token-size=2",
				"--innodb-ft-max-token-size=50",
				"--default-time-zone=Europe/Berlin",
			},
		},
		func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %q", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := m.pool.Retry(func() error {
		db, err := sql.Open("mysql", m.getDSN())
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %q", err)
	}

	m.prepareDBForFirstUse()

	m.LoadBaseData()

	m.db, err = sql.Open("mysql", m.getDSN())
	if err != nil {
		log.Fatalf("Could not connect to database after setup: %q", err)
	}
}

func (m *dbServer) DB() *sql.DB {
	if m.db == nil {
		log.Fatal("Test DB connection has not been established! You are accessing DB() method too early.")
	}

	return m.db
}

func (m *dbServer) getDSN() string {
	// Using `root` isn't cool, but a workaround for now to create triggers in the database
	return fmt.Sprintf("root:secret@(127.0.0.1:%s)/fivenettest?collation=utf8mb4_unicode_ci&parseTime=True&loc=Local", m.resource.GetPort("3306/tcp"))
}

func (m *dbServer) prepareDBForFirstUse() {
	// Load and apply premigrate.sql file
	m.loadSQLFile(filepath.Join(tests.TestDataSQLPath, "initial_esx.sql"))

	// Use DB migrations to handle the rest
	if err := query.MigrateDB(zap.NewNop(), m.getDSN()); err != nil {
		log.Fatalf("Failed to migrate test database: %q", err)
	}
}

func (m *dbServer) getMultiStatementDB() *sql.DB {
	// Open db connection with multiStatements param so we can apply sql files
	initDB, err := sql.Open("mysql", m.getDSN()+"&multiStatements=true")
	if err != nil {
		log.Fatalf("Failed to open test database connection for multi statement exec: %q", err)
	}
	return initDB
}

func (m *dbServer) loadSQLFile(file string) {
	initDB := m.getMultiStatementDB()

	c, ioErr := os.ReadFile(file)
	if ioErr != nil {
		log.Fatalf("Failed to read %s for tests: %s", file, ioErr)
	}
	sqlBase := string(c)
	if _, err := initDB.Exec(sqlBase); err != nil {
		log.Fatalf("Failed to apply %s for tests: %s", file, err)
	}
}

func (m *dbServer) LoadBaseData() {
	path := filepath.Join(tests.TestDataSQLPath, "base_*.sql")
	files, err := filepath.Glob(path)
	if err != nil {
		log.Fatalf("failed to find base data sql files (%s): %s", path, err)
	}
	// Sort the found files as they might not be in lexical order which we
	// need for this case https://github.com/golang/go/issues/17153
	sort.Strings(files)

	for _, file := range files {
		m.loadSQLFile(file)
	}
}

func (m *dbServer) Stop() {
	if m.stopped {
		return
	}
	m.stopped = true

	// You can't defer this because os.Exit doesn't care for defer
	if err := m.pool.Purge(m.resource); err != nil {
		log.Fatalf("Could not purge container resource: %q", err)
	}
}

// Reset truncates all `fivenet_*` tables and loads the base test data
func (m *dbServer) Reset() {
	initDB := m.getMultiStatementDB()

	rows, err := initDB.Query("SHOW TABLES LIKE 'fivenet_%';")
	if err != nil {
		log.Fatalf("Failed to list fivenet tables in test database: %q", err)
	}

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan table name to string: %q", err)
		}

		// Placeholders aren't supported for table names, see
		// https://github.com/go-sql-driver/mysql/issues/848#issuecomment-414910152`
		if _, err := initDB.Exec("SET FOREIGN_KEY_CHECKS = 0; TRUNCATE TABLE `" + tableName + "`; SET FOREIGN_KEY_CHECKS = 1;"); err != nil {
			log.Printf("Failed to truncate %s table: %s", tableName, err)
		}
	}

	// Load base test data after every reset
	m.LoadBaseData()
}