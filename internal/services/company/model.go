// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package company

import (
	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/apijson"
	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CompanyModel struct {
	CompanyID     types.String                                            `tfsdk:"company_id" path:"company_id,optional"`
	WorkspaceID   types.String                                            `tfsdk:"workspace_id" json:"workspace_id,required,no_refresh"`
	Company       *CompanyCompanyModel                                    `tfsdk:"company" json:"company,required"`
	SourceID      types.String                                            `tfsdk:"source_id" json:"source_id,optional"`
	TechnicalData *CompanyTechnicalDataModel                              `tfsdk:"technical_data" json:"technical_data,optional"`
	PortalURL     types.String                                            `tfsdk:"portal_url" json:"portal_url,computed"`
	WebviewURL    types.String                                            `tfsdk:"webview_url" json:"webview_url,computed"`
	AmlSuspicions customfield.NestedObjectList[CompanyAmlSuspicionsModel] `tfsdk:"aml_suspicions" json:"aml_suspicions,computed"`
	Certificat    customfield.NestedObject[CompanyCertificatModel]        `tfsdk:"certificat" json:"certificat,computed"`
	Checks        customfield.NestedObjectList[CompanyChecksModel]        `tfsdk:"checks" json:"checks,computed"`
	Documents     customfield.NestedObjectList[CompanyDocumentsModel]     `tfsdk:"documents" json:"documents,computed"`
	Members       customfield.NestedObjectList[CompanyMembersModel]       `tfsdk:"members" json:"members,computed"`
	Properties    customfield.NestedObjectList[CompanyPropertiesModel]    `tfsdk:"properties" json:"properties,computed"`
	Risk          customfield.NestedObject[CompanyRiskModel]              `tfsdk:"risk" json:"risk,computed"`
}

func (m CompanyModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CompanyModel) MarshalJSONForUpdate(state CompanyModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type CompanyCompanyModel struct {
	Name                         types.String `tfsdk:"name" json:"name,required"`
	Address                      types.String `tfsdk:"address" json:"address,optional"`
	CommercialName               types.String `tfsdk:"commercial_name" json:"commercial_name,optional"`
	Country                      types.String `tfsdk:"country" json:"country,optional"`
	Email                        types.String `tfsdk:"email" json:"email,optional"`
	EmployerIdentificationNumber types.String `tfsdk:"employer_identification_number" json:"employer_identification_number,optional"`
	LegalForm                    types.String `tfsdk:"legal_form" json:"legal_form,optional"`
	PhoneNumber                  types.String `tfsdk:"phone_number" json:"phone_number,optional"`
	RegistrationDate             types.String `tfsdk:"registration_date" json:"registration_date,optional,no_refresh"`
	RegistrationID               types.String `tfsdk:"registration_id" json:"registration_id,optional"`
	ShareCapital                 types.String `tfsdk:"share_capital" json:"share_capital,optional"`
	Status                       types.String `tfsdk:"status" json:"status,optional"`
	TaxIdentificationNumber      types.String `tfsdk:"tax_identification_number" json:"tax_identification_number,optional"`
	Type                         types.String `tfsdk:"type" json:"type,optional"`
	WebsiteURL                   types.String `tfsdk:"website_url" json:"website_url,optional"`
}

type CompanyTechnicalDataModel struct {
	ActiveAmlSuspicions         types.Bool    `tfsdk:"active_aml_suspicions" json:"active_aml_suspicions,optional"`
	CallbackURL                 types.String  `tfsdk:"callback_url" json:"callback_url,optional"`
	CallbackURLNotification     types.String  `tfsdk:"callback_url_notification" json:"callback_url_notification,optional"`
	FilteringScoreAmlSuspicions types.Float64 `tfsdk:"filtering_score_aml_suspicions" json:"filtering_score_aml_suspicions,optional"`
	Language                    types.String  `tfsdk:"language" json:"language,optional"`
	RawData                     types.Bool    `tfsdk:"raw_data" json:"raw_data,optional"`
}

type CompanyAmlSuspicionsModel struct {
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

type CompanyCertificatModel struct {
	ID        types.String      `tfsdk:"id" json:"id,computed"`
	CreatedAt timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Filename  types.String      `tfsdk:"filename" json:"filename,computed"`
}

type CompanyChecksModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type CompanyDocumentsModel struct {
	ID           types.String                                              `tfsdk:"id" json:"id,computed"`
	Checks       customfield.NestedObjectList[CompanyDocumentsChecksModel] `tfsdk:"checks" json:"checks,computed"`
	CreatedAt    timetypes.RFC3339                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DocumentType types.String                                              `tfsdk:"document_type" json:"document_type,computed"`
	Name         types.String                                              `tfsdk:"name" json:"name,computed"`
	SignedURL    types.String                                              `tfsdk:"signed_url" json:"signed_url,computed"`
	State        types.String                                              `tfsdk:"state" json:"state,computed"`
	Status       types.String                                              `tfsdk:"status" json:"status,computed"`
	Tables       customfield.NestedObjectList[CompanyDocumentsTablesModel] `tfsdk:"tables" json:"tables,computed"`
	Values       customfield.NestedObjectList[CompanyDocumentsValuesModel] `tfsdk:"values" json:"values,computed"`
}

type CompanyDocumentsChecksModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type CompanyDocumentsTablesModel struct {
	Operation customfield.List[jsontypes.Normalized] `tfsdk:"operation" json:"operation,computed"`
}

type CompanyDocumentsValuesModel struct {
	Confidence types.Float64                 `tfsdk:"confidence" json:"confidence,computed"`
	Name       types.String                  `tfsdk:"name" json:"name,computed"`
	Value      customfield.List[types.Int64] `tfsdk:"value" json:"value,computed"`
}

type CompanyMembersModel struct {
	ID                   types.String                                               `tfsdk:"id" json:"id,computed"`
	Address              types.String                                               `tfsdk:"address" json:"address,computed"`
	Birthday             timetypes.RFC3339                                          `tfsdk:"birthday" json:"birthday,computed" format:"date-time"`
	Birthplace           types.String                                               `tfsdk:"birthplace" json:"birthplace,computed"`
	Country              types.String                                               `tfsdk:"country" json:"country,computed"`
	Documents            customfield.NestedObjectList[CompanyMembersDocumentsModel] `tfsdk:"documents" json:"documents,computed"`
	Email                types.String                                               `tfsdk:"email" json:"email,computed"`
	FirstName            types.String                                               `tfsdk:"first_name" json:"first_name,computed"`
	IsBeneficialOwner    types.Bool                                                 `tfsdk:"is_beneficial_owner" json:"is_beneficial_owner,computed"`
	IsDelegator          types.Bool                                                 `tfsdk:"is_delegator" json:"is_delegator,computed"`
	LastName             types.String                                               `tfsdk:"last_name" json:"last_name,computed"`
	LivenessVerification types.Bool                                                 `tfsdk:"liveness_verification" json:"liveness_verification,computed"`
	Name                 types.String                                               `tfsdk:"name" json:"name,computed"`
	OwnershipPercentage  types.Int64                                                `tfsdk:"ownership_percentage" json:"ownership_percentage,computed"`
	PhoneNumber          types.String                                               `tfsdk:"phone_number" json:"phone_number,computed"`
	PostalCode           types.String                                               `tfsdk:"postal_code" json:"postal_code,computed"`
	RegistrationID       types.String                                               `tfsdk:"registration_id" json:"registration_id,computed"`
	Relation             types.String                                               `tfsdk:"relation" json:"relation,computed"`
	Roles                types.String                                               `tfsdk:"roles" json:"roles,computed"`
	Source               types.String                                               `tfsdk:"source" json:"source,computed"`
	State                types.String                                               `tfsdk:"state" json:"state,computed"`
	Status               types.String                                               `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                               `tfsdk:"type" json:"type,computed"`
	WorkspaceID          types.String                                               `tfsdk:"workspace_id" json:"workspace_id,computed"`
}

type CompanyMembersDocumentsModel struct {
	ID           types.String                                                     `tfsdk:"id" json:"id,computed"`
	Checks       customfield.NestedObjectList[CompanyMembersDocumentsChecksModel] `tfsdk:"checks" json:"checks,computed"`
	CreatedAt    timetypes.RFC3339                                                `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	DocumentType types.String                                                     `tfsdk:"document_type" json:"document_type,computed"`
	Name         types.String                                                     `tfsdk:"name" json:"name,computed"`
	SignedURL    types.String                                                     `tfsdk:"signed_url" json:"signed_url,computed"`
	State        types.String                                                     `tfsdk:"state" json:"state,computed"`
	Status       types.String                                                     `tfsdk:"status" json:"status,computed"`
	Tables       customfield.NestedObjectList[CompanyMembersDocumentsTablesModel] `tfsdk:"tables" json:"tables,computed"`
	Values       customfield.NestedObjectList[CompanyMembersDocumentsValuesModel] `tfsdk:"values" json:"values,computed"`
}

type CompanyMembersDocumentsChecksModel struct {
	Masked   types.Bool   `tfsdk:"masked" json:"masked,computed"`
	Message  types.String `tfsdk:"message" json:"message,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	Validate types.Bool   `tfsdk:"validate" json:"validate,computed"`
	Weight   types.Int64  `tfsdk:"weight" json:"weight,computed"`
}

type CompanyMembersDocumentsTablesModel struct {
	Operation customfield.List[jsontypes.Normalized] `tfsdk:"operation" json:"operation,computed"`
}

type CompanyMembersDocumentsValuesModel struct {
	Confidence types.Float64                 `tfsdk:"confidence" json:"confidence,computed"`
	Name       types.String                  `tfsdk:"name" json:"name,computed"`
	Value      customfield.List[types.Int64] `tfsdk:"value" json:"value,computed"`
}

type CompanyPropertiesModel struct {
	Name  types.String `tfsdk:"name" json:"name,computed"`
	Type  types.String `tfsdk:"type" json:"type,computed"`
	Value types.String `tfsdk:"value" json:"value,computed"`
}

type CompanyRiskModel struct {
	Code   types.String  `tfsdk:"code" json:"code,computed"`
	Reason types.String  `tfsdk:"reason" json:"reason,computed"`
	Score  types.Float64 `tfsdk:"score" json:"score,computed"`
}
