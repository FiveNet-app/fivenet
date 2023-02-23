// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"github.com/galexrt/arpanet/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Job{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Job{}) fail: %s", err)
	}
}

func Test_jobQuery(t *testing.T) {
	job := newJob(db)
	job = *job.As(job.TableName())
	_do := job.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(job.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <jobs> fail:", err)
		return
	}

	_, ok := job.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from job success")
	}

	err = _do.Create(&model.Job{})
	if err != nil {
		t.Error("create item in table <jobs> fail:", err)
	}

	err = _do.Save(&model.Job{})
	if err != nil {
		t.Error("create item in table <jobs> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Job{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <jobs> fail:", err)
	}

	_, err = _do.Select(job.ALL).Take()
	if err != nil {
		t.Error("Take() on table <jobs> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <jobs> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <jobs> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <jobs> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Job{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <jobs> fail:", err)
	}

	_, err = _do.Select(job.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <jobs> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <jobs> fail:", err)
	}

	_, err = _do.Select(job.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <jobs> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <jobs> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <jobs> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <jobs> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Job{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <jobs> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <jobs> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <jobs> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <jobs> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <jobs> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <jobs> fail:", err)
	}
}
