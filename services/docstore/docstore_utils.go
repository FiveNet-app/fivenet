package docstore

import (
	context "context"
	"slices"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/documents"
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/users"
	permscitizenstore "github.com/fivenet-app/fivenet/gen/go/proto/services/citizenstore/perms"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth/userinfo"
	"github.com/fivenet-app/fivenet/pkg/grpc/errswrap"
	"github.com/fivenet-app/fivenet/pkg/perms"
	"github.com/fivenet-app/fivenet/pkg/utils/dbutils/tables"
	errorsdocstore "github.com/fivenet-app/fivenet/services/docstore/errors"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

func (s *Server) listDocumentsQuery(where jet.BoolExpression, onlyColumns jet.ProjectionList, userInfo *userinfo.UserInfo) jet.SelectStatement {
	tCreator := tables.Users().AS("creator")

	wheres := []jet.BoolExpression{}
	if !userInfo.SuperUser {
		wheres = []jet.BoolExpression{
			jet.AND(
				tDocumentShort.DeletedAt.IS_NULL(),
				jet.OR(
					tDocumentShort.Public.IS_TRUE(),
					jet.AND(
						tDocumentShort.CreatorID.EQ(jet.Int32(userInfo.UserId)),
						tDocumentShort.CreatorJob.EQ(jet.String(userInfo.Job)),
					),
					jet.OR(
						jet.AND(
							tDUserAccess.Access.IS_NOT_NULL(),
							tDUserAccess.Access.GT(jet.Int32(int32(documents.AccessLevel_ACCESS_LEVEL_BLOCKED))),
						),
						jet.AND(
							tDUserAccess.Access.IS_NULL(),
							tDJobAccess.Access.IS_NOT_NULL(),
							tDJobAccess.Access.GT(jet.Int32(int32(documents.AccessLevel_ACCESS_LEVEL_BLOCKED))),
						),
					),
				),
			),
		}
	}

	if where != nil {
		wheres = append(wheres, where)
	}

	var q jet.SelectStatement
	if onlyColumns != nil {
		q = tDocumentShort.
			SELECT(
				onlyColumns,
			)
	} else {
		columns := jet.ProjectionList{
			tDocumentShort.ID,
			tDocumentShort.CreatedAt,
			tDocumentShort.UpdatedAt,
			tDocumentShort.DeletedAt,
			tDocumentShort.CategoryID,
			tDCategory.ID,
			tDCategory.Name,
			tDCategory.Description,
			tDCategory.Job,
			tDCategory.Color,
			tDCategory.Icon,
			tDocumentShort.Title,
			tDocumentShort.ContentType,
			tDocumentShort.Summary.AS("documentshort.content"),
			tDocumentShort.CreatorID,
			tDocumentShort.TemplateID,
			tCreator.ID,
			tCreator.Job,
			tCreator.JobGrade,
			tCreator.Firstname,
			tCreator.Lastname,
			tCreator.Dateofbirth,
			tDocumentShort.CreatorJob,
			tDocumentShort.State,
			tDocumentShort.Closed,
			tDocumentShort.Public,
			tDocumentShort.TemplateID,
			tWorkflow.DocumentID,
			tWorkflow.AutoCloseTime,
			tWorkflow.NextReminderTime,
		}

		if userInfo.SuperUser {
			columns = append(columns, tDocumentShort.DeletedAt)
		}

		// Field Permission Check
		fieldsAttr, _ := s.ps.Attr(userInfo, permscitizenstore.CitizenStoreServicePerm, permscitizenstore.CitizenStoreServiceListCitizensPerm, permscitizenstore.CitizenStoreServiceListCitizensFieldsPermField)
		var fields perms.StringList
		if fieldsAttr != nil {
			fields = fieldsAttr.([]string)
		}

		if slices.Contains(fields, "PhoneNumber") {
			columns = append(columns, tCreator.PhoneNumber)
		}

		q = tDocumentShort.SELECT(columns[0], columns[1:])
	}

	var tables jet.ReadableTable
	if !userInfo.SuperUser {
		tables = tDocumentShort.
			LEFT_JOIN(tDUserAccess,
				tDUserAccess.DocumentID.EQ(tDocumentShort.ID).
					AND(tDUserAccess.UserID.EQ(jet.Int32(userInfo.UserId))),
			).
			LEFT_JOIN(tDJobAccess,
				tDJobAccess.DocumentID.EQ(tDocumentShort.ID).
					AND(tDJobAccess.Job.EQ(jet.String(userInfo.Job))).
					AND(tDJobAccess.MinimumGrade.LT_EQ(jet.Int32(userInfo.JobGrade))),
			).
			LEFT_JOIN(tDCategory,
				tDocumentShort.CategoryID.EQ(tDCategory.ID),
			).
			LEFT_JOIN(tCreator,
				tDocumentShort.CreatorID.EQ(tCreator.ID),
			).
			LEFT_JOIN(tWorkflow,
				tWorkflow.DocumentID.EQ(tDocumentShort.ID),
			)
	} else {
		tables = tDocumentShort.
			LEFT_JOIN(tDCategory,
				tDocumentShort.CategoryID.EQ(tDCategory.ID),
			).
			LEFT_JOIN(tCreator,
				tDocumentShort.CreatorID.EQ(tCreator.ID),
			).
			LEFT_JOIN(tWorkflow,
				tWorkflow.DocumentID.EQ(tDocumentShort.ID),
			)
	}

	return q.
		FROM(tables).
		WHERE(jet.AND(
			wheres...,
		))
}

func (s *Server) getDocumentQuery(where jet.BoolExpression, onlyColumns jet.ProjectionList, userInfo *userinfo.UserInfo, withContent bool) jet.SelectStatement {
	tCreator := tables.Users().AS("creator")

	var wheres []jet.BoolExpression
	if !userInfo.SuperUser {
		wheres = []jet.BoolExpression{
			jet.AND(
				tDocument.DeletedAt.IS_NULL(),
				jet.OR(
					jet.OR(
						tDocument.Public.IS_TRUE(),
						tDocument.CreatorID.EQ(jet.Int32(userInfo.UserId)),
					),
					jet.OR(
						jet.AND(
							tDUserAccess.Access.IS_NOT_NULL(),
							tDUserAccess.Access.GT(jet.Int32(int32(documents.AccessLevel_ACCESS_LEVEL_BLOCKED))),
						),
						jet.AND(
							tDUserAccess.Access.IS_NULL(),
							tDJobAccess.Access.IS_NOT_NULL(),
							tDJobAccess.Access.GT(jet.Int32(int32(documents.AccessLevel_ACCESS_LEVEL_BLOCKED))),
						),
					),
				),
			),
		}
	}

	if where != nil {
		wheres = append(wheres, where)
	}

	var q jet.SelectStatement
	if onlyColumns != nil {
		q = tDocument.
			SELECT(
				onlyColumns,
			)
	} else {
		columns := jet.ProjectionList{
			tDocument.ID,
			tDocument.CreatedAt,
			tDocument.UpdatedAt,
			tDocument.DeletedAt,
			tDocument.CategoryID,
			tDCategory.ID,
			tDCategory.Name,
			tDCategory.Description,
			tDCategory.Job,
			tDCategory.Color,
			tDCategory.Icon,
			tDocument.Title,
			tDocument.ContentType,
			tDocument.CreatorID,
			tCreator.ID,
			tCreator.Job,
			tCreator.JobGrade,
			tCreator.Firstname,
			tCreator.Lastname,
			tCreator.Dateofbirth,
			tDocument.CreatorJob,
			tDocument.State,
			tDocument.Closed,
			tDocument.Public,
			tDocument.TemplateID,
			tDPins.State.AS("document.pinned"),
			tWorkflow.DocumentID,
			tWorkflow.AutoCloseTime,
			tWorkflow.NextReminderTime,
			tUserWorkflow.DocumentID,
			tUserWorkflow.UserID,
			tUserWorkflow.ManualReminderTime,
			tUserWorkflow.ManualReminderMessage,
		}

		if withContent {
			columns = append(columns,
				tDocument.Data,
				tDocument.Content,
			)
		}

		if userInfo.SuperUser {
			columns = append(columns, tDocument.DeletedAt)
		}

		// Field Permission Check
		fieldsAttr, _ := s.ps.Attr(userInfo, permscitizenstore.CitizenStoreServicePerm, permscitizenstore.CitizenStoreServiceListCitizensPerm, permscitizenstore.CitizenStoreServiceListCitizensFieldsPermField)
		var fields perms.StringList
		if fieldsAttr != nil {
			fields = fieldsAttr.([]string)
		}

		if slices.Contains(fields, "PhoneNumber") {
			columns = append(columns, tCreator.PhoneNumber)
		}

		q = tDocument.SELECT(columns[0], columns[1:])
	}

	var tables jet.ReadableTable
	if !userInfo.SuperUser {
		tables = tDocument.
			LEFT_JOIN(tDUserAccess,
				tDUserAccess.DocumentID.EQ(tDocument.ID).
					AND(tDUserAccess.UserID.EQ(jet.Int32(userInfo.UserId))),
			).
			LEFT_JOIN(tDJobAccess,
				tDJobAccess.DocumentID.EQ(tDocument.ID).
					AND(tDJobAccess.Job.EQ(jet.String(userInfo.Job))).
					AND(tDJobAccess.MinimumGrade.LT_EQ(jet.Int32(userInfo.JobGrade))),
			).
			LEFT_JOIN(tDCategory,
				tDocument.CategoryID.EQ(tDCategory.ID),
			).
			LEFT_JOIN(tCreator,
				tDocument.CreatorID.EQ(tCreator.ID),
			).
			LEFT_JOIN(tDPins,
				tDPins.DocumentID.EQ(tDocument.ID),
			).
			LEFT_JOIN(tWorkflow,
				tWorkflow.DocumentID.EQ(tDocument.ID),
			).
			LEFT_JOIN(tUserWorkflow,
				tUserWorkflow.DocumentID.EQ(tDocument.ID).
					AND(tUserWorkflow.UserID.EQ(jet.Int32(userInfo.UserId))),
			)
	} else {
		tables = tDocument.
			LEFT_JOIN(tDCategory,
				tDocument.CategoryID.EQ(tDCategory.ID),
			).
			LEFT_JOIN(tCreator,
				tDocument.CreatorID.EQ(tCreator.ID),
			).
			LEFT_JOIN(tDPins,
				tDPins.DocumentID.EQ(tDocument.ID),
			).
			LEFT_JOIN(tWorkflow,
				tWorkflow.DocumentID.EQ(tDocument.ID),
			).
			LEFT_JOIN(tUserWorkflow,
				tUserWorkflow.DocumentID.EQ(tDocument.ID).
					AND(tUserWorkflow.UserID.EQ(jet.Int32(userInfo.UserId))),
			)
	}

	return q.
		FROM(tables).
		WHERE(jet.AND(
			wheres...,
		)).
		ORDER_BY(
			tDocument.CreatedAt.DESC(),
			tDocument.UpdatedAt.DESC(),
		)
}

func (s *Server) updateDocumentOwner(ctx context.Context, tx qrm.DB, documentId uint64, userInfo *userinfo.UserInfo, newOwner *users.UserShort) error {
	stmt := tDocument.
		UPDATE(
			tDocument.CreatorID,
		).
		SET(
			newOwner.UserId,
		).
		WHERE(
			tDocument.ID.EQ(jet.Uint64(documentId)),
		)

	if _, err := stmt.ExecContext(ctx, tx); err != nil {
		return errswrap.NewError(err, errorsdocstore.ErrFailedQuery)
	}

	if _, err := addDocumentActivity(ctx, tx, &documents.DocActivity{
		DocumentId:   documentId,
		ActivityType: documents.DocActivityType_DOC_ACTIVITY_TYPE_OWNER_CHANGED,
		CreatorId:    &userInfo.UserId,
		CreatorJob:   userInfo.Job,
		Data: &documents.DocActivityData{
			Data: &documents.DocActivityData_OwnerChanged{
				OwnerChanged: &documents.DocOwnerChanged{
					NewOwnerId: newOwner.UserId,
					NewOwner:   newOwner,
				},
			},
		},
	}); err != nil {
		return errswrap.NewError(err, errorsdocstore.ErrFailedQuery)
	}

	return nil
}