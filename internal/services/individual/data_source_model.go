// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package individual

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

type IndividualDataSourceModel struct {
	IndividualID  types.String                                                         `tfsdk:"individual_id" path:"individual_id,required"`
	Document      types.Bool                                                           `tfsdk:"document" query:"document,optional"`
	Scope         types.String                                                         `tfsdk:"scope" query:"scope,optional"`
	AuthURL       types.String                                                         `tfsdk:"auth_url" json:"auth_url,computed"`
	CreatedAt     timetypes.RFC3339                                                    `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ID            types.String                                                         `tfsdk:"id" json:"id,computed"`
	Number        types.Int64                                                          `tfsdk:"number" json:"number,computed"`
	PortalURL     types.String                                                         `tfsdk:"portal_url" json:"portal_url,computed"`
	SourceID      types.String                                                         `tfsdk:"source_id" json:"source_id,computed"`
	State         types.String                                                         `tfsdk:"state" json:"state,computed"`
	Status        types.String                                                         `tfsdk:"status" json:"status,computed"`
	WebviewURL    types.String                                                         `tfsdk:"webview_url" json:"webview_url,computed"`
	WorkspaceID   types.String                                                         `tfsdk:"workspace_id" json:"workspace_id,computed"`
	AmlSuspicions customfield.NestedObjectList[IndividualAmlSuspicionsDataSourceModel] `tfsdk:"aml_suspicions" json:"aml_suspicions,computed"`
	Certificat    customfield.NestedObject[IndividualCertificatDataSourceModel]        `tfsdk:"certificat" json:"certificat,computed"`
	Checks        customfield.NestedObjectList[IndividualChecksDataSourceModel]        `tfsdk:"checks" json:"checks,computed"`
	Documents     customfield.NestedObjectList[IndividualDocumentsDataSourceModel]     `tfsdk:"documents" json:"documents,computed"`
	IdentityCard  customfield.NestedObject[IndividualIdentityCardDataSourceModel]      `tfsdk:"identity_card" json:"identity_card,computed"`
	Person        customfield.NestedObject[IndividualPersonDataSourceModel]            `tfsdk:"person" json:"person,computed"`
	Properties    customfield.NestedObjectList[IndividualPropertiesDataSourceModel]    `tfsdk:"properties" json:"properties,computed"`
	Risk          customfield.NestedObject[IndividualRiskDataSourceModel]              `tfsdk:"risk" json:"risk,computed"`
	Tags          customfield.NestedObjectList[IndividualTagsDataSourceModel]          `tfsdk:"tags" json:"tags,computed"`
	TechnicalData customfield.NestedObject[IndividualTechnicalDataDataSourceModel]     `tfsdk:"technical_data" json:"technical_data,computed"`
}

func (m *IndividualDataSourceModel) toReadParams(_ context.Context) (params dataleonlabs.IndividualGetParams, diags diag.Diagnostics) {
	params = dataleonlabs.IndividualGetParams{}

	if !m.Document.IsNull() {
		params.Document = param.NewOpt(m.Document.ValueBool())
	}
	if !m.Scope.IsNull() {
		params.Scope = param.NewOpt(m.Scope.ValueString())
	}

	return
}

type IndividualAmlSuspicionsDataSourceModel struct {
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

type IndividualCertificatDataSourceModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Filename  types.String      `tfsdk:"filename" json:"filename,computed"`
}

type IndividualChecksDataSourceModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type IndividualDocumentsDataSourceModel struct {
	ID           types.String                                                           `tfsdk:"id" json:"id,computed"`
	Checks       customfield.NestedObjectList[IndividualDocumentsChecksDataSourceModel] `tfsdk:"checks" json:"checks,computed"`
	CreatedAt    timetypes.RFC3339                                                      `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DocumentType types.String                                                           `tfsdk:"document_type" json:"document_type,computed"`
	Name         types.String                                                           `tfsdk:"name" json:"name,computed"`
	SignedURL    types.String                                                           `tfsdk:"signed_url" json:"signed_url,computed"`
	State        types.String                                                           `tfsdk:"state" json:"state,computed"`
	Status       types.String                                                           `tfsdk:"status" json:"status,computed"`
	Tables       customfield.NestedObjectList[IndividualDocumentsTablesDataSourceModel] `tfsdk:"tables" json:"tables,computed"`
	Values       customfield.NestedObjectList[IndividualDocumentsValuesDataSourceModel] `tfsdk:"values" json:"values,computed"`
}

type IndividualDocumentsChecksDataSourceModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type IndividualDocumentsTablesDataSourceModel struct {
	Operation customfield.List[jsontypes.Normalized] `tfsdk:"operation" json:"operation,computed"`
}

type IndividualDocumentsValuesDataSourceModel struct {
	Confidence types.Float64                 `tfsdk:"confidence" json:"confidence,computed"`
	Name       types.String                  `tfsdk:"name" json:"name,computed"`
	Value      customfield.List[types.Int64] `tfsdk:"value" json:"value,computed"`
}

type IndividualIdentityCardDataSourceModel struct {
	ID                     types.String `tfsdk:"id" json:"id,computed"`
	BackDocumentSignedURL  types.String `tfsdk:"back_document_signed_url" json:"back_document_signed_url,computed"`
	BirthPlace             types.String `tfsdk:"birth_place" json:"birth_place,computed"`
	Birthday               types.String `tfsdk:"birthday" json:"birthday,computed"`
	Country                types.String `tfsdk:"country" json:"country,computed"`
	ExpirationDate         types.String `tfsdk:"expiration_date" json:"expiration_date,computed"`
	FirstName              types.String `tfsdk:"first_name" json:"first_name,computed"`
	FrontDocumentSignedURL types.String `tfsdk:"front_document_signed_url" json:"front_document_signed_url,computed"`
	Gender                 types.String `tfsdk:"gender" json:"gender,computed"`
	IssueDate              types.String `tfsdk:"issue_date" json:"issue_date,computed"`
	LastName               types.String `tfsdk:"last_name" json:"last_name,computed"`
	MrzLine1               types.String `tfsdk:"mrz_line_1" json:"mrz_line_1,computed"`
	MrzLine2               types.String `tfsdk:"mrz_line_2" json:"mrz_line_2,computed"`
	MrzLine3               types.String `tfsdk:"mrz_line_3" json:"mrz_line_3,computed"`
	Type                   types.String `tfsdk:"type" json:"type,computed"`
}

type IndividualPersonDataSourceModel struct {
	Birthday           types.String `tfsdk:"birthday" json:"birthday,computed"`
	Email              types.String `tfsdk:"email" json:"email,computed"`
	FaceImageSignedURL types.String `tfsdk:"face_image_signed_url" json:"face_image_signed_url,computed"`
	FirstName          types.String `tfsdk:"first_name" json:"first_name,computed"`
	FullName           types.String `tfsdk:"full_name" json:"full_name,computed"`
	Gender             types.String `tfsdk:"gender" json:"gender,computed"`
	LastName           types.String `tfsdk:"last_name" json:"last_name,computed"`
	MaidenName         types.String `tfsdk:"maiden_name" json:"maiden_name,computed"`
	Nationality        types.String `tfsdk:"nationality" json:"nationality,computed"`
	PhoneNumber        types.String `tfsdk:"phone_number" json:"phone_number,computed"`
}

type IndividualPropertiesDataSourceModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type IndividualRiskDataSourceModel struct {
	Code   types.String  `tfsdk:"code" json:"code,computed"`
	Reason types.String  `tfsdk:"reason" json:"reason,computed"`
	Score  types.Float64 `tfsdk:"score" json:"score,computed"`
}

type IndividualTagsDataSourceModel struct {
	Key     types.String `tfsdk:"key" json:"key,computed"`
	Private types.Bool   `tfsdk:"private" json:"private,computed"`
	Type    types.String `tfsdk:"type" json:"type,computed"`
	Value   types.String `tfsdk:"value" json:"value,computed"`
}

type IndividualTechnicalDataDataSourceModel struct {
	ActiveAmlSuspicions         types.Bool        `tfsdk:"active_aml_suspicions" json:"active_aml_suspicions,computed"`
	APIVersion                  types.Int64       `tfsdk:"api_version" json:"api_version,computed"`
	ApprovedAt                  timetypes.RFC3339 `tfsdk:"approved_at" json:"approved_at,computed" format:"date-time"`
	CallbackURL                 types.String      `tfsdk:"callback_url" json:"callback_url,computed"`
	CallbackURLNotification     types.String      `tfsdk:"callback_url_notification" json:"callback_url_notification,computed"`
	DisableNotification         types.Bool        `tfsdk:"disable_notification" json:"disable_notification,computed"`
	DisableNotificationDate     timetypes.RFC3339 `tfsdk:"disable_notification_date" json:"disable_notification_date,computed" format:"date-time"`
	ExportType                  types.String      `tfsdk:"export_type" json:"export_type,computed"`
	FilteringScoreAmlSuspicions types.Float64     `tfsdk:"filtering_score_aml_suspicions" json:"filtering_score_aml_suspicions,computed"`
	FinishedAt                  timetypes.RFC3339 `tfsdk:"finished_at" json:"finished_at,computed" format:"date-time"`
	IP                          types.String      `tfsdk:"ip" json:"ip,computed"`
	Language                    types.String      `tfsdk:"language" json:"language,computed"`
	LocationIP                  types.String      `tfsdk:"location_ip" json:"location_ip,computed"`
	NeedReviewAt                timetypes.RFC3339 `tfsdk:"need_review_at" json:"need_review_at,computed" format:"date-time"`
	NotificationConfirmation    types.Bool        `tfsdk:"notification_confirmation" json:"notification_confirmation,computed"`
	QrCode                      types.String      `tfsdk:"qr_code" json:"qr_code,computed"`
	RawData                     types.Bool        `tfsdk:"raw_data" json:"raw_data,computed"`
	RejectedAt                  timetypes.RFC3339 `tfsdk:"rejected_at" json:"rejected_at,computed" format:"date-time"`
	SessionDuration             types.Int64       `tfsdk:"session_duration" json:"session_duration,computed"`
	StartedAt                   timetypes.RFC3339 `tfsdk:"started_at" json:"started_at,computed" format:"date-time"`
	TransferAt                  timetypes.RFC3339 `tfsdk:"transfer_at" json:"transfer_at,computed" format:"date-time"`
	TransferMode                types.String      `tfsdk:"transfer_mode" json:"transfer_mode,computed"`
}
