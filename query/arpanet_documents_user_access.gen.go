// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/galexrt/arpanet/model"
)

func newDocumentUserAccess(db *gorm.DB, opts ...gen.DOOption) documentUserAccess {
	_documentUserAccess := documentUserAccess{}

	_documentUserAccess.documentUserAccessDo.UseDB(db, opts...)
	_documentUserAccess.documentUserAccessDo.UseModel(&model.DocumentUserAccess{})

	tableName := _documentUserAccess.documentUserAccessDo.TableName()
	_documentUserAccess.ALL = field.NewAsterisk(tableName)
	_documentUserAccess.ID = field.NewUint(tableName, "id")
	_documentUserAccess.CreatedAt = field.NewTime(tableName, "created_at")
	_documentUserAccess.UpdatedAt = field.NewTime(tableName, "updated_at")
	_documentUserAccess.DocumentID = field.NewUint(tableName, "document_id")
	_documentUserAccess.Identifier = field.NewString(tableName, "identifier")
	_documentUserAccess.Access = field.NewString(tableName, "access")

	_documentUserAccess.fillFieldMap()

	return _documentUserAccess
}

type documentUserAccess struct {
	documentUserAccessDo

	ALL        field.Asterisk
	ID         field.Uint
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DocumentID field.Uint
	Identifier field.String
	Access     field.String

	fieldMap map[string]field.Expr
}

func (d documentUserAccess) Table(newTableName string) *documentUserAccess {
	d.documentUserAccessDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d documentUserAccess) As(alias string) *documentUserAccess {
	d.documentUserAccessDo.DO = *(d.documentUserAccessDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *documentUserAccess) updateTableName(table string) *documentUserAccess {
	d.ALL = field.NewAsterisk(table)
	d.ID = field.NewUint(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DocumentID = field.NewUint(table, "document_id")
	d.Identifier = field.NewString(table, "identifier")
	d.Access = field.NewString(table, "access")

	d.fillFieldMap()

	return d
}

func (d *documentUserAccess) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *documentUserAccess) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 6)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["document_id"] = d.DocumentID
	d.fieldMap["identifier"] = d.Identifier
	d.fieldMap["access"] = d.Access
}

func (d documentUserAccess) clone(db *gorm.DB) documentUserAccess {
	d.documentUserAccessDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d documentUserAccess) replaceDB(db *gorm.DB) documentUserAccess {
	d.documentUserAccessDo.ReplaceDB(db)
	return d
}

type documentUserAccessDo struct{ gen.DO }

type IDocumentUserAccessDo interface {
	gen.SubQuery
	Debug() IDocumentUserAccessDo
	WithContext(ctx context.Context) IDocumentUserAccessDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDocumentUserAccessDo
	WriteDB() IDocumentUserAccessDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDocumentUserAccessDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDocumentUserAccessDo
	Not(conds ...gen.Condition) IDocumentUserAccessDo
	Or(conds ...gen.Condition) IDocumentUserAccessDo
	Select(conds ...field.Expr) IDocumentUserAccessDo
	Where(conds ...gen.Condition) IDocumentUserAccessDo
	Order(conds ...field.Expr) IDocumentUserAccessDo
	Distinct(cols ...field.Expr) IDocumentUserAccessDo
	Omit(cols ...field.Expr) IDocumentUserAccessDo
	Join(table schema.Tabler, on ...field.Expr) IDocumentUserAccessDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDocumentUserAccessDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDocumentUserAccessDo
	Group(cols ...field.Expr) IDocumentUserAccessDo
	Having(conds ...gen.Condition) IDocumentUserAccessDo
	Limit(limit int) IDocumentUserAccessDo
	Offset(offset int) IDocumentUserAccessDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDocumentUserAccessDo
	Unscoped() IDocumentUserAccessDo
	Create(values ...*model.DocumentUserAccess) error
	CreateInBatches(values []*model.DocumentUserAccess, batchSize int) error
	Save(values ...*model.DocumentUserAccess) error
	First() (*model.DocumentUserAccess, error)
	Take() (*model.DocumentUserAccess, error)
	Last() (*model.DocumentUserAccess, error)
	Find() ([]*model.DocumentUserAccess, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DocumentUserAccess, err error)
	FindInBatches(result *[]*model.DocumentUserAccess, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.DocumentUserAccess) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDocumentUserAccessDo
	Assign(attrs ...field.AssignExpr) IDocumentUserAccessDo
	Joins(fields ...field.RelationField) IDocumentUserAccessDo
	Preload(fields ...field.RelationField) IDocumentUserAccessDo
	FirstOrInit() (*model.DocumentUserAccess, error)
	FirstOrCreate() (*model.DocumentUserAccess, error)
	FindByPage(offset int, limit int) (result []*model.DocumentUserAccess, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDocumentUserAccessDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d documentUserAccessDo) Debug() IDocumentUserAccessDo {
	return d.withDO(d.DO.Debug())
}

func (d documentUserAccessDo) WithContext(ctx context.Context) IDocumentUserAccessDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d documentUserAccessDo) ReadDB() IDocumentUserAccessDo {
	return d.Clauses(dbresolver.Read)
}

func (d documentUserAccessDo) WriteDB() IDocumentUserAccessDo {
	return d.Clauses(dbresolver.Write)
}

func (d documentUserAccessDo) Session(config *gorm.Session) IDocumentUserAccessDo {
	return d.withDO(d.DO.Session(config))
}

func (d documentUserAccessDo) Clauses(conds ...clause.Expression) IDocumentUserAccessDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d documentUserAccessDo) Returning(value interface{}, columns ...string) IDocumentUserAccessDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d documentUserAccessDo) Not(conds ...gen.Condition) IDocumentUserAccessDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d documentUserAccessDo) Or(conds ...gen.Condition) IDocumentUserAccessDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d documentUserAccessDo) Select(conds ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d documentUserAccessDo) Where(conds ...gen.Condition) IDocumentUserAccessDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d documentUserAccessDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDocumentUserAccessDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d documentUserAccessDo) Order(conds ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d documentUserAccessDo) Distinct(cols ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d documentUserAccessDo) Omit(cols ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d documentUserAccessDo) Join(table schema.Tabler, on ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d documentUserAccessDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d documentUserAccessDo) RightJoin(table schema.Tabler, on ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d documentUserAccessDo) Group(cols ...field.Expr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d documentUserAccessDo) Having(conds ...gen.Condition) IDocumentUserAccessDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d documentUserAccessDo) Limit(limit int) IDocumentUserAccessDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d documentUserAccessDo) Offset(offset int) IDocumentUserAccessDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d documentUserAccessDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDocumentUserAccessDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d documentUserAccessDo) Unscoped() IDocumentUserAccessDo {
	return d.withDO(d.DO.Unscoped())
}

func (d documentUserAccessDo) Create(values ...*model.DocumentUserAccess) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d documentUserAccessDo) CreateInBatches(values []*model.DocumentUserAccess, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d documentUserAccessDo) Save(values ...*model.DocumentUserAccess) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d documentUserAccessDo) First() (*model.DocumentUserAccess, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocumentUserAccess), nil
	}
}

func (d documentUserAccessDo) Take() (*model.DocumentUserAccess, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocumentUserAccess), nil
	}
}

func (d documentUserAccessDo) Last() (*model.DocumentUserAccess, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocumentUserAccess), nil
	}
}

func (d documentUserAccessDo) Find() ([]*model.DocumentUserAccess, error) {
	result, err := d.DO.Find()
	return result.([]*model.DocumentUserAccess), err
}

func (d documentUserAccessDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DocumentUserAccess, err error) {
	buf := make([]*model.DocumentUserAccess, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d documentUserAccessDo) FindInBatches(result *[]*model.DocumentUserAccess, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d documentUserAccessDo) Attrs(attrs ...field.AssignExpr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d documentUserAccessDo) Assign(attrs ...field.AssignExpr) IDocumentUserAccessDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d documentUserAccessDo) Joins(fields ...field.RelationField) IDocumentUserAccessDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d documentUserAccessDo) Preload(fields ...field.RelationField) IDocumentUserAccessDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d documentUserAccessDo) FirstOrInit() (*model.DocumentUserAccess, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocumentUserAccess), nil
	}
}

func (d documentUserAccessDo) FirstOrCreate() (*model.DocumentUserAccess, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DocumentUserAccess), nil
	}
}

func (d documentUserAccessDo) FindByPage(offset int, limit int) (result []*model.DocumentUserAccess, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d documentUserAccessDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d documentUserAccessDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d documentUserAccessDo) Delete(models ...*model.DocumentUserAccess) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *documentUserAccessDo) withDO(do gen.Dao) *documentUserAccessDo {
	d.DO = *do.(*gen.DO)
	return d
}
