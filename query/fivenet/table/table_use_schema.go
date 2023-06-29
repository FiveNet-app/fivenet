//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	FivenetAccounts = FivenetAccounts.FromSchema(schema)
	FivenetAttrs = FivenetAttrs.FromSchema(schema)
	FivenetAuditLog = FivenetAuditLog.FromSchema(schema)
	FivenetCentrumUnits = FivenetCentrumUnits.FromSchema(schema)
	FivenetCentrumUnitsUsers = FivenetCentrumUnitsUsers.FromSchema(schema)
	FivenetDocuments = FivenetDocuments.FromSchema(schema)
	FivenetDocumentsCategories = FivenetDocumentsCategories.FromSchema(schema)
	FivenetDocumentsComments = FivenetDocumentsComments.FromSchema(schema)
	FivenetDocumentsJobAccess = FivenetDocumentsJobAccess.FromSchema(schema)
	FivenetDocumentsReferences = FivenetDocumentsReferences.FromSchema(schema)
	FivenetDocumentsRelations = FivenetDocumentsRelations.FromSchema(schema)
	FivenetDocumentsTemplates = FivenetDocumentsTemplates.FromSchema(schema)
	FivenetDocumentsTemplatesJobAccess = FivenetDocumentsTemplatesJobAccess.FromSchema(schema)
	FivenetDocumentsUserAccess = FivenetDocumentsUserAccess.FromSchema(schema)
	FivenetJobProps = FivenetJobProps.FromSchema(schema)
	FivenetLawbooks = FivenetLawbooks.FromSchema(schema)
	FivenetLawbooksLaws = FivenetLawbooksLaws.FromSchema(schema)
	FivenetNotifications = FivenetNotifications.FromSchema(schema)
	FivenetOauth2Accounts = FivenetOauth2Accounts.FromSchema(schema)
	FivenetPermissions = FivenetPermissions.FromSchema(schema)
	FivenetRoles = FivenetRoles.FromSchema(schema)
	FivenetRoleAttrs = FivenetRoleAttrs.FromSchema(schema)
	FivenetRolePermissions = FivenetRolePermissions.FromSchema(schema)
	FivenetUserActivity = FivenetUserActivity.FromSchema(schema)
	FivenetUserLocations = FivenetUserLocations.FromSchema(schema)
	FivenetUserProps = FivenetUserProps.FromSchema(schema)
	GksphoneJobMessage = GksphoneJobMessage.FromSchema(schema)
	Jobs = Jobs.FromSchema(schema)
	JobGrades = JobGrades.FromSchema(schema)
	Licenses = Licenses.FromSchema(schema)
	OwnedVehicles = OwnedVehicles.FromSchema(schema)
	Users = Users.FromSchema(schema)
	UserLicenses = UserLicenses.FromSchema(schema)
}
