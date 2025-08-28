// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package individual

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stainless-sdks/dataleonlabs-terraform/internal/apijson"
	"github.com/stainless-sdks/dataleonlabs-terraform/internal/customfield"
)

type IndividualModel struct {
	ID            types.String                                               `tfsdk:"id" json:"id,computed"`
	WorkspaceID   types.String                                               `tfsdk:"workspace_id" json:"workspace_id,required"`
	SourceID      types.String                                               `tfsdk:"source_id" json:"source_id,optional"`
	Person        *IndividualPersonModel                                     `tfsdk:"person" json:"person,optional"`
	TechnicalData *IndividualTechnicalDataModel                              `tfsdk:"technical_data" json:"technical_data,optional"`
	AuthURL       types.String                                               `tfsdk:"auth_url" json:"auth_url,computed"`
	CreatedAt     timetypes.RFC3339                                          `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Number        types.Int64                                                `tfsdk:"number" json:"number,computed"`
	PortalURL     types.String                                               `tfsdk:"portal_url" json:"portal_url,computed"`
	State         types.String                                               `tfsdk:"state" json:"state,computed"`
	Status        types.String                                               `tfsdk:"status" json:"status,computed"`
	WebviewURL    types.String                                               `tfsdk:"webview_url" json:"webview_url,computed"`
	AmlSuspicions customfield.NestedObjectList[IndividualAmlSuspicionsModel] `tfsdk:"aml_suspicions" json:"aml_suspicions,computed"`
	Certificat    customfield.NestedObject[IndividualCertificatModel]        `tfsdk:"certificat" json:"certificat,computed"`
	Checks        customfield.NestedObjectList[IndividualChecksModel]        `tfsdk:"checks" json:"checks,computed"`
	Documents     customfield.NestedObjectList[IndividualDocumentsModel]     `tfsdk:"documents" json:"documents,computed"`
	IdentityCard  customfield.NestedObject[IndividualIdentityCardModel]      `tfsdk:"identity_card" json:"identity_card,computed"`
	Properties    customfield.NestedObjectList[IndividualPropertiesModel]    `tfsdk:"properties" json:"properties,computed"`
	Risk          customfield.NestedObject[IndividualRiskModel]              `tfsdk:"risk" json:"risk,computed"`
	Tags          customfield.NestedObjectList[IndividualTagsModel]          `tfsdk:"tags" json:"tags,computed"`
}

func (m IndividualModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m IndividualModel) MarshalJSONForUpdate(state IndividualModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type IndividualPersonModel struct {
	Birthday    types.String `tfsdk:"birthday" json:"birthday,optional"`
	Email       types.String `tfsdk:"email" json:"email,optional"`
	FirstName   types.String `tfsdk:"first_name" json:"first_name,optional"`
	Gender      types.String `tfsdk:"gender" json:"gender,optional"`
	LastName    types.String `tfsdk:"last_name" json:"last_name,optional"`
	MaidenName  types.String `tfsdk:"maiden_name" json:"maiden_name,optional"`
	PhoneNumber types.String `tfsdk:"phone_number" json:"phone_number,optional"`
}

type IndividualTechnicalDataModel struct {
	ActiveAmlSuspicions     types.Bool   `tfsdk:"active_aml_suspicions" json:"active_aml_suspicions,optional"`
	CallbackURL             types.String `tfsdk:"callback_url" json:"callback_url,optional"`
	CallbackURLNotification types.String `tfsdk:"callback_url_notification" json:"callback_url_notification,optional"`
	Language                types.String `tfsdk:"language" json:"language,optional"`
	RawData                 types.Bool   `tfsdk:"raw_data" json:"raw_data,optional"`
}

type IndividualAmlSuspicionsModel struct {
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

type IndividualCertificatModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Filename  types.String      `tfsdk:"filename" json:"filename,computed"`
}

type IndividualChecksModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type IndividualDocumentsModel struct {
	ID           types.String                                                 `tfsdk:"id" json:"id,computed"`
	Checks       customfield.NestedObjectList[IndividualDocumentsChecksModel] `tfsdk:"checks" json:"checks,computed"`
	CreatedAt    timetypes.RFC3339                                            `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DocumentType types.String                                                 `tfsdk:"document_type" json:"document_type,computed"`
	Name         types.String                                                 `tfsdk:"name" json:"name,computed"`
	SignedURL    types.String                                                 `tfsdk:"signed_url" json:"signed_url,computed"`
	State        types.String                                                 `tfsdk:"state" json:"state,computed"`
	Status       types.String                                                 `tfsdk:"status" json:"status,computed"`
	Tables       customfield.NestedObjectList[IndividualDocumentsTablesModel] `tfsdk:"tables" json:"tables,computed"`
	Values       customfield.NestedObjectList[IndividualDocumentsValuesModel] `tfsdk:"values" json:"values,computed"`
}

type IndividualDocumentsChecksModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type IndividualDocumentsTablesModel struct {
	Operation customfield.List[jsontypes.Normalized] `tfsdk:"operation" json:"operation,computed"`
}

type IndividualDocumentsValuesModel struct {
	Confidence types.Float64                 `tfsdk:"confidence" json:"confidence,computed"`
	Name       types.String                  `tfsdk:"name" json:"name,computed"`
	Value      customfield.List[types.Int64] `tfsdk:"value" json:"value,computed"`
}

type IndividualIdentityCardModel struct {
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

type IndividualPropertiesModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type IndividualRiskModel struct {
	Code   types.String  `tfsdk:"code" json:"code,computed"`
	Reason types.String  `tfsdk:"reason" json:"reason,computed"`
	Score  types.Float64 `tfsdk:"score" json:"score,computed"`
}

type IndividualTagsModel struct {
	Key     types.String `tfsdk:"key" json:"key,computed"`
	Private types.Bool   `tfsdk:"private" json:"private,computed"`
	Type    types.String `tfsdk:"type" json:"type,computed"`
	Value   types.String `tfsdk:"value" json:"value,computed"`
}
