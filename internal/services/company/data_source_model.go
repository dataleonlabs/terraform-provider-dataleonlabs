// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package company

import (
	"context"

	"github.com/dataleonlabs/dataleonlabs-go"
	"github.com/dataleonlabs/dataleonlabs-go/packages/param"
	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CompanyDataSourceModel struct {
	CompanyID     types.String                                                      `tfsdk:"company_id" path:"company_id,required"`
	Document      types.Bool                                                        `tfsdk:"document" query:"document,optional"`
	Scope         types.String                                                      `tfsdk:"scope" query:"scope,optional"`
	PortalURL     types.String                                                      `tfsdk:"portal_url" json:"portal_url,computed"`
	SourceID      types.String                                                      `tfsdk:"source_id" json:"source_id,computed"`
	WebviewURL    types.String                                                      `tfsdk:"webview_url" json:"webview_url,computed"`
	AmlSuspicions customfield.NestedObjectList[CompanyAmlSuspicionsDataSourceModel] `tfsdk:"aml_suspicions" json:"aml_suspicions,computed"`
	Certificat    customfield.NestedObject[CompanyCertificatDataSourceModel]        `tfsdk:"certificat" json:"certificat,computed"`
	Checks        customfield.NestedObjectList[CompanyChecksDataSourceModel]        `tfsdk:"checks" json:"checks,computed"`
	Company       customfield.NestedObject[CompanyCompanyDataSourceModel]           `tfsdk:"company" json:"company,computed"`
	Documents     customfield.NestedObjectList[CompanyDocumentsDataSourceModel]     `tfsdk:"documents" json:"documents,computed"`
	Members       customfield.NestedObjectList[CompanyMembersDataSourceModel]       `tfsdk:"members" json:"members,computed"`
	Properties    customfield.NestedObjectList[CompanyPropertiesDataSourceModel]    `tfsdk:"properties" json:"properties,computed"`
	Risk          customfield.NestedObject[CompanyRiskDataSourceModel]              `tfsdk:"risk" json:"risk,computed"`
	TechnicalData customfield.NestedObject[CompanyTechnicalDataDataSourceModel]     `tfsdk:"technical_data" json:"technical_data,computed"`
}

func (m *CompanyDataSourceModel) toReadParams(_ context.Context) (params dataleonlabs.CompanyGetParams, diags diag.Diagnostics) {
	params = dataleonlabs.CompanyGetParams{}

	if !m.Document.IsNull() {
		params.Document = param.NewOpt(m.Document.ValueBool())
	}
	if !m.Scope.IsNull() {
		params.Scope = param.NewOpt(m.Scope.ValueString())
	}

	return
}

type CompanyAmlSuspicionsDataSourceModel struct {
	Caption  types.String  `tfsdk:"caption" json:"caption,computed"`
	Country  types.String  `tfsdk:"country" json:"country,computed"`
	Gender   types.String  `tfsdk:"gender" json:"gender,computed"`
	Relation types.String  `tfsdk:"relation" json:"relation,computed"`
	Schema   types.String  `tfsdk:"schema" json:"schema,computed"`
	Score    types.Float64 `tfsdk:"score" json:"score,computed"`
	Source   types.String  `tfsdk:"source" json:"source,computed"`
	Status   types.String  `tfsdk:"status" json:"status,computed"`
	Type     types.String  `tfsdk:"type" json:"type,computed"`
}

type CompanyCertificatDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Filename  types.String      `tfsdk:"filename" json:"filename,computed"`
}

type CompanyChecksDataSourceModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type CompanyCompanyDataSourceModel struct {
	Address                      types.String                                                   `tfsdk:"address" json:"address,computed"`
	ClosureDate                  timetypes.RFC3339                                              `tfsdk:"closure_date" json:"closure_date,computed" format:"date"`
	CommercialName               types.String                                                   `tfsdk:"commercial_name" json:"commercial_name,computed"`
	Contact                      customfield.NestedObject[CompanyCompanyContactDataSourceModel] `tfsdk:"contact" json:"contact,computed"`
	Country                      types.String                                                   `tfsdk:"country" json:"country,computed"`
	Email                        types.String                                                   `tfsdk:"email" json:"email,computed"`
	Employees                    types.Int64                                                    `tfsdk:"employees" json:"employees,computed"`
	EmployerIdentificationNumber types.String                                                   `tfsdk:"employer_identification_number" json:"employer_identification_number,computed"`
	InsolvencyExists             types.Bool                                                     `tfsdk:"insolvency_exists" json:"insolvency_exists,computed"`
	InsolvencyOngoing            types.Bool                                                     `tfsdk:"insolvency_ongoing" json:"insolvency_ongoing,computed"`
	LegalForm                    types.String                                                   `tfsdk:"legal_form" json:"legal_form,computed"`
	Name                         types.String                                                   `tfsdk:"name" json:"name,computed"`
	PhoneNumber                  types.String                                                   `tfsdk:"phone_number" json:"phone_number,computed"`
	RegistrationDate             timetypes.RFC3339                                              `tfsdk:"registration_date" json:"registration_date,computed" format:"date"`
	RegistrationID               types.String                                                   `tfsdk:"registration_id" json:"registration_id,computed"`
	ShareCapital                 types.String                                                   `tfsdk:"share_capital" json:"share_capital,computed"`
	Status                       types.String                                                   `tfsdk:"status" json:"status,computed"`
	TaxIdentificationNumber      types.String                                                   `tfsdk:"tax_identification_number" json:"tax_identification_number,computed"`
	Type                         types.String                                                   `tfsdk:"type" json:"type,computed"`
	WebsiteURL                   types.String                                                   `tfsdk:"website_url" json:"website_url,computed"`
}

type CompanyCompanyContactDataSourceModel struct {
	Department  types.String `tfsdk:"department" json:"department,computed"`
	Email       types.String `tfsdk:"email" json:"email,computed"`
	FirstName   types.String `tfsdk:"first_name" json:"first_name,computed"`
	LastName    types.String `tfsdk:"last_name" json:"last_name,computed"`
	PhoneNumber types.String `tfsdk:"phone_number" json:"phone_number,computed"`
}

type CompanyDocumentsDataSourceModel struct {
	ID           types.String                                                        `tfsdk:"id" json:"id,computed"`
	Checks       customfield.NestedObjectList[CompanyDocumentsChecksDataSourceModel] `tfsdk:"checks" json:"checks,computed"`
	CreatedAt    timetypes.RFC3339                                                   `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DocumentType types.String                                                        `tfsdk:"document_type" json:"document_type,computed"`
	Name         types.String                                                        `tfsdk:"name" json:"name,computed"`
	SignedURL    types.String                                                        `tfsdk:"signed_url" json:"signed_url,computed"`
	State        types.String                                                        `tfsdk:"state" json:"state,computed"`
	Status       types.String                                                        `tfsdk:"status" json:"status,computed"`
	Tables       customfield.NestedObjectList[CompanyDocumentsTablesDataSourceModel] `tfsdk:"tables" json:"tables,computed"`
	Values       customfield.NestedObjectList[CompanyDocumentsValuesDataSourceModel] `tfsdk:"values" json:"values,computed"`
}

type CompanyDocumentsChecksDataSourceModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type CompanyDocumentsTablesDataSourceModel struct {
	Operation customfield.List[jsontypes.Normalized] `tfsdk:"operation" json:"operation,computed"`
}

type CompanyDocumentsValuesDataSourceModel struct {
	Confidence types.Float64                 `tfsdk:"confidence" json:"confidence,computed"`
	Name       types.String                  `tfsdk:"name" json:"name,computed"`
	Value      customfield.List[types.Int64] `tfsdk:"value" json:"value,computed"`
}

type CompanyMembersDataSourceModel struct {
	ID                   types.String                                                         `tfsdk:"id" json:"id,computed"`
	Address              types.String                                                         `tfsdk:"address" json:"address,computed"`
	Birthday             timetypes.RFC3339                                                    `tfsdk:"birthday" json:"birthday,computed" format:"date-time"`
	Birthplace           types.String                                                         `tfsdk:"birthplace" json:"birthplace,computed"`
	Country              types.String                                                         `tfsdk:"country" json:"country,computed"`
	Documents            customfield.NestedObjectList[CompanyMembersDocumentsDataSourceModel] `tfsdk:"documents" json:"documents,computed"`
	Email                types.String                                                         `tfsdk:"email" json:"email,computed"`
	FirstName            types.String                                                         `tfsdk:"first_name" json:"first_name,computed"`
	IsBeneficialOwner    types.Bool                                                           `tfsdk:"is_beneficial_owner" json:"is_beneficial_owner,computed"`
	IsDelegator          types.Bool                                                           `tfsdk:"is_delegator" json:"is_delegator,computed"`
	LastName             types.String                                                         `tfsdk:"last_name" json:"last_name,computed"`
	LivenessVerification types.Bool                                                           `tfsdk:"liveness_verification" json:"liveness_verification,computed"`
	Name                 types.String                                                         `tfsdk:"name" json:"name,computed"`
	OwnershipPercentage  types.Int64                                                          `tfsdk:"ownership_percentage" json:"ownership_percentage,computed"`
	PhoneNumber          types.String                                                         `tfsdk:"phone_number" json:"phone_number,computed"`
	PostalCode           types.String                                                         `tfsdk:"postal_code" json:"postal_code,computed"`
	RegistrationID       types.String                                                         `tfsdk:"registration_id" json:"registration_id,computed"`
	Relation             types.String                                                         `tfsdk:"relation" json:"relation,computed"`
	Roles                types.String                                                         `tfsdk:"roles" json:"roles,computed"`
	Source               types.String                                                         `tfsdk:"source" json:"source,computed"`
	State                types.String                                                         `tfsdk:"state" json:"state,computed"`
	Status               types.String                                                         `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                                         `tfsdk:"type" json:"type,computed"`
	WorkspaceID          types.String                                                         `tfsdk:"workspace_id" json:"workspace_id,computed"`
}

type CompanyMembersDocumentsDataSourceModel struct {
	ID           types.String                                                               `tfsdk:"id" json:"id,computed"`
	Checks       customfield.NestedObjectList[CompanyMembersDocumentsChecksDataSourceModel] `tfsdk:"checks" json:"checks,computed"`
	CreatedAt    timetypes.RFC3339                                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DocumentType types.String                                                               `tfsdk:"document_type" json:"document_type,computed"`
	Name         types.String                                                               `tfsdk:"name" json:"name,computed"`
	SignedURL    types.String                                                               `tfsdk:"signed_url" json:"signed_url,computed"`
	State        types.String                                                               `tfsdk:"state" json:"state,computed"`
	Status       types.String                                                               `tfsdk:"status" json:"status,computed"`
	Tables       customfield.NestedObjectList[CompanyMembersDocumentsTablesDataSourceModel] `tfsdk:"tables" json:"tables,computed"`
	Values       customfield.NestedObjectList[CompanyMembersDocumentsValuesDataSourceModel] `tfsdk:"values" json:"values,computed"`
}

type CompanyMembersDocumentsChecksDataSourceModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type CompanyMembersDocumentsTablesDataSourceModel struct {
	Operation customfield.List[jsontypes.Normalized] `tfsdk:"operation" json:"operation,computed"`
}

type CompanyMembersDocumentsValuesDataSourceModel struct {
	Confidence types.Float64                 `tfsdk:"confidence" json:"confidence,computed"`
	Name       types.String                  `tfsdk:"name" json:"name,computed"`
	Value      customfield.List[types.Int64] `tfsdk:"value" json:"value,computed"`
}

type CompanyPropertiesDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type CompanyRiskDataSourceModel struct {
	Code   types.String  `tfsdk:"code" json:"code,computed"`
	Reason types.String  `tfsdk:"reason" json:"reason,computed"`
	Score  types.Float64 `tfsdk:"score" json:"score,computed"`
}

type CompanyTechnicalDataDataSourceModel struct {
	ActiveAmlSuspicions      types.Bool        `tfsdk:"active_aml_suspicions" json:"active_aml_suspicions,computed"`
	APIVersion               types.Int64       `tfsdk:"api_version" json:"api_version,computed"`
	ApprovedAt               timetypes.RFC3339 `tfsdk:"approved_at" json:"approved_at,computed" format:"date-time"`
	CallbackURL              types.String      `tfsdk:"callback_url" json:"callback_url,computed"`
	CallbackURLNotification  types.String      `tfsdk:"callback_url_notification" json:"callback_url_notification,computed"`
	DisableNotification      types.Bool        `tfsdk:"disable_notification" json:"disable_notification,computed"`
	DisableNotificationDate  timetypes.RFC3339 `tfsdk:"disable_notification_date" json:"disable_notification_date,computed" format:"date-time"`
	ExportType               types.String      `tfsdk:"export_type" json:"export_type,computed"`
	FinishedAt               timetypes.RFC3339 `tfsdk:"finished_at" json:"finished_at,computed" format:"date-time"`
	IP                       types.String      `tfsdk:"ip" json:"ip,computed"`
	Language                 types.String      `tfsdk:"language" json:"language,computed"`
	LocationIP               types.String      `tfsdk:"location_ip" json:"location_ip,computed"`
	NeedReviewAt             timetypes.RFC3339 `tfsdk:"need_review_at" json:"need_review_at,computed" format:"date-time"`
	NotificationConfirmation types.Bool        `tfsdk:"notification_confirmation" json:"notification_confirmation,computed"`
	QrCode                   types.String      `tfsdk:"qr_code" json:"qr_code,computed"`
	RawData                  types.Bool        `tfsdk:"raw_data" json:"raw_data,computed"`
	RejectedAt               timetypes.RFC3339 `tfsdk:"rejected_at" json:"rejected_at,computed" format:"date-time"`
	SessionDuration          types.Int64       `tfsdk:"session_duration" json:"session_duration,computed"`
	StartedAt                timetypes.RFC3339 `tfsdk:"started_at" json:"started_at,computed" format:"date-time"`
	TransferAt               timetypes.RFC3339 `tfsdk:"transfer_at" json:"transfer_at,computed" format:"date-time"`
	TransferMode             types.String      `tfsdk:"transfer_mode" json:"transfer_mode,computed"`
}
