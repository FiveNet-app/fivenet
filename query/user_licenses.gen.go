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

func newUserLicense(db *gorm.DB, opts ...gen.DOOption) userLicense {
	_userLicense := userLicense{}

	_userLicense.userLicenseDo.UseDB(db, opts...)
	_userLicense.userLicenseDo.UseModel(&model.UserLicense{})

	tableName := _userLicense.userLicenseDo.TableName()
	_userLicense.ALL = field.NewAsterisk(tableName)
	_userLicense.Type = field.NewField(tableName, "type")
	_userLicense.Owner = field.NewString(tableName, "owner")

	_userLicense.fillFieldMap()

	return _userLicense
}

type userLicense struct {
	userLicenseDo

	ALL   field.Asterisk
	Type  field.Field
	Owner field.String

	fieldMap map[string]field.Expr
}

func (u userLicense) Table(newTableName string) *userLicense {
	u.userLicenseDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userLicense) As(alias string) *userLicense {
	u.userLicenseDo.DO = *(u.userLicenseDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userLicense) updateTableName(table string) *userLicense {
	u.ALL = field.NewAsterisk(table)
	u.Type = field.NewField(table, "type")
	u.Owner = field.NewString(table, "owner")

	u.fillFieldMap()

	return u
}

func (u *userLicense) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userLicense) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 2)
	u.fieldMap["type"] = u.Type
	u.fieldMap["owner"] = u.Owner
}

func (u userLicense) clone(db *gorm.DB) userLicense {
	u.userLicenseDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userLicense) replaceDB(db *gorm.DB) userLicense {
	u.userLicenseDo.ReplaceDB(db)
	return u
}

type userLicenseDo struct{ gen.DO }

type IUserLicenseDo interface {
	gen.SubQuery
	Debug() IUserLicenseDo
	WithContext(ctx context.Context) IUserLicenseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserLicenseDo
	WriteDB() IUserLicenseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserLicenseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserLicenseDo
	Not(conds ...gen.Condition) IUserLicenseDo
	Or(conds ...gen.Condition) IUserLicenseDo
	Select(conds ...field.Expr) IUserLicenseDo
	Where(conds ...gen.Condition) IUserLicenseDo
	Order(conds ...field.Expr) IUserLicenseDo
	Distinct(cols ...field.Expr) IUserLicenseDo
	Omit(cols ...field.Expr) IUserLicenseDo
	Join(table schema.Tabler, on ...field.Expr) IUserLicenseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserLicenseDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserLicenseDo
	Group(cols ...field.Expr) IUserLicenseDo
	Having(conds ...gen.Condition) IUserLicenseDo
	Limit(limit int) IUserLicenseDo
	Offset(offset int) IUserLicenseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserLicenseDo
	Unscoped() IUserLicenseDo
	Create(values ...*model.UserLicense) error
	CreateInBatches(values []*model.UserLicense, batchSize int) error
	Save(values ...*model.UserLicense) error
	First() (*model.UserLicense, error)
	Take() (*model.UserLicense, error)
	Last() (*model.UserLicense, error)
	Find() ([]*model.UserLicense, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserLicense, err error)
	FindInBatches(result *[]*model.UserLicense, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserLicense) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserLicenseDo
	Assign(attrs ...field.AssignExpr) IUserLicenseDo
	Joins(fields ...field.RelationField) IUserLicenseDo
	Preload(fields ...field.RelationField) IUserLicenseDo
	FirstOrInit() (*model.UserLicense, error)
	FirstOrCreate() (*model.UserLicense, error)
	FindByPage(offset int, limit int) (result []*model.UserLicense, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserLicenseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userLicenseDo) Debug() IUserLicenseDo {
	return u.withDO(u.DO.Debug())
}

func (u userLicenseDo) WithContext(ctx context.Context) IUserLicenseDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userLicenseDo) ReadDB() IUserLicenseDo {
	return u.Clauses(dbresolver.Read)
}

func (u userLicenseDo) WriteDB() IUserLicenseDo {
	return u.Clauses(dbresolver.Write)
}

func (u userLicenseDo) Session(config *gorm.Session) IUserLicenseDo {
	return u.withDO(u.DO.Session(config))
}

func (u userLicenseDo) Clauses(conds ...clause.Expression) IUserLicenseDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userLicenseDo) Returning(value interface{}, columns ...string) IUserLicenseDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userLicenseDo) Not(conds ...gen.Condition) IUserLicenseDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userLicenseDo) Or(conds ...gen.Condition) IUserLicenseDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userLicenseDo) Select(conds ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userLicenseDo) Where(conds ...gen.Condition) IUserLicenseDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userLicenseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserLicenseDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userLicenseDo) Order(conds ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userLicenseDo) Distinct(cols ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userLicenseDo) Omit(cols ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userLicenseDo) Join(table schema.Tabler, on ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userLicenseDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userLicenseDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userLicenseDo) Group(cols ...field.Expr) IUserLicenseDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userLicenseDo) Having(conds ...gen.Condition) IUserLicenseDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userLicenseDo) Limit(limit int) IUserLicenseDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userLicenseDo) Offset(offset int) IUserLicenseDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userLicenseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserLicenseDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userLicenseDo) Unscoped() IUserLicenseDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userLicenseDo) Create(values ...*model.UserLicense) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userLicenseDo) CreateInBatches(values []*model.UserLicense, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userLicenseDo) Save(values ...*model.UserLicense) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userLicenseDo) First() (*model.UserLicense, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLicense), nil
	}
}

func (u userLicenseDo) Take() (*model.UserLicense, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLicense), nil
	}
}

func (u userLicenseDo) Last() (*model.UserLicense, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLicense), nil
	}
}

func (u userLicenseDo) Find() ([]*model.UserLicense, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserLicense), err
}

func (u userLicenseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserLicense, err error) {
	buf := make([]*model.UserLicense, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userLicenseDo) FindInBatches(result *[]*model.UserLicense, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userLicenseDo) Attrs(attrs ...field.AssignExpr) IUserLicenseDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userLicenseDo) Assign(attrs ...field.AssignExpr) IUserLicenseDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userLicenseDo) Joins(fields ...field.RelationField) IUserLicenseDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userLicenseDo) Preload(fields ...field.RelationField) IUserLicenseDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userLicenseDo) FirstOrInit() (*model.UserLicense, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLicense), nil
	}
}

func (u userLicenseDo) FirstOrCreate() (*model.UserLicense, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserLicense), nil
	}
}

func (u userLicenseDo) FindByPage(offset int, limit int) (result []*model.UserLicense, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userLicenseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userLicenseDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userLicenseDo) Delete(models ...*model.UserLicense) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userLicenseDo) withDO(do gen.Dao) *userLicenseDo {
	u.DO = *do.(*gen.DO)
	return u
}
