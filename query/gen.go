// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                  = new(Query)
	Account            *account
	Document           *document
	DocumentJobAccess  *documentJobAccess
	DocumentUserAccess *documentUserAccess
	Job                *job
	JobGrade           *jobGrade
	User               *user
	UserLicense        *userLicense
	UserProps          *userProps
	VpcL               *vpcL
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Account = &Q.Account
	Document = &Q.Document
	DocumentJobAccess = &Q.DocumentJobAccess
	DocumentUserAccess = &Q.DocumentUserAccess
	Job = &Q.Job
	JobGrade = &Q.JobGrade
	User = &Q.User
	UserLicense = &Q.UserLicense
	UserProps = &Q.UserProps
	VpcL = &Q.VpcL
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		Account:            newAccount(db, opts...),
		Document:           newDocument(db, opts...),
		DocumentJobAccess:  newDocumentJobAccess(db, opts...),
		DocumentUserAccess: newDocumentUserAccess(db, opts...),
		Job:                newJob(db, opts...),
		JobGrade:           newJobGrade(db, opts...),
		User:               newUser(db, opts...),
		UserLicense:        newUserLicense(db, opts...),
		UserProps:          newUserProps(db, opts...),
		VpcL:               newVpcL(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Account            account
	Document           document
	DocumentJobAccess  documentJobAccess
	DocumentUserAccess documentUserAccess
	Job                job
	JobGrade           jobGrade
	User               user
	UserLicense        userLicense
	UserProps          userProps
	VpcL               vpcL
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Account:            q.Account.clone(db),
		Document:           q.Document.clone(db),
		DocumentJobAccess:  q.DocumentJobAccess.clone(db),
		DocumentUserAccess: q.DocumentUserAccess.clone(db),
		Job:                q.Job.clone(db),
		JobGrade:           q.JobGrade.clone(db),
		User:               q.User.clone(db),
		UserLicense:        q.UserLicense.clone(db),
		UserProps:          q.UserProps.clone(db),
		VpcL:               q.VpcL.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		Account:            q.Account.replaceDB(db),
		Document:           q.Document.replaceDB(db),
		DocumentJobAccess:  q.DocumentJobAccess.replaceDB(db),
		DocumentUserAccess: q.DocumentUserAccess.replaceDB(db),
		Job:                q.Job.replaceDB(db),
		JobGrade:           q.JobGrade.replaceDB(db),
		User:               q.User.replaceDB(db),
		UserLicense:        q.UserLicense.replaceDB(db),
		UserProps:          q.UserProps.replaceDB(db),
		VpcL:               q.VpcL.replaceDB(db),
	}
}

type queryCtx struct {
	Account            IAccountDo
	Document           IDocumentDo
	DocumentJobAccess  IDocumentJobAccessDo
	DocumentUserAccess IDocumentUserAccessDo
	Job                IJobDo
	JobGrade           IJobGradeDo
	User               IUserDo
	UserLicense        IUserLicenseDo
	UserProps          IUserPropsDo
	VpcL               IVpcLDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Account:            q.Account.WithContext(ctx),
		Document:           q.Document.WithContext(ctx),
		DocumentJobAccess:  q.DocumentJobAccess.WithContext(ctx),
		DocumentUserAccess: q.DocumentUserAccess.WithContext(ctx),
		Job:                q.Job.WithContext(ctx),
		JobGrade:           q.JobGrade.WithContext(ctx),
		User:               q.User.WithContext(ctx),
		UserLicense:        q.UserLicense.WithContext(ctx),
		UserProps:          q.UserProps.WithContext(ctx),
		VpcL:               q.VpcL.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
