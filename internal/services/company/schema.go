// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package company

import (
	"context"

	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*CompanyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"company_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"workspace_id": schema.StringAttribute{
				Description: "Unique identifier of the workspace in which the company is being created.",
				Required:    true,
			},
			"company": schema.SingleNestedAttribute{
				Description: "Main information about the company being registered.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "Legal name of the company.",
						Required:    true,
					},
					"address": schema.StringAttribute{
						Description: "Registered address of the company.",
						Optional:    true,
					},
					"commercial_name": schema.StringAttribute{
						Description: "Commercial or trade name of the company, if different from the legal name.",
						Optional:    true,
					},
					"country": schema.StringAttribute{
						Description: `ISO 3166-1 alpha-2 country code of company registration (e.g., "FR" for France).`,
						Optional:    true,
					},
					"email": schema.StringAttribute{
						Description: "Contact email address for the company.",
						Optional:    true,
					},
					"employer_identification_number": schema.StringAttribute{
						Description: "Employer Identification Number (EIN) or equivalent.",
						Optional:    true,
					},
					"legal_form": schema.StringAttribute{
						Description: "Legal structure of the company (e.g., SARL, SAS).",
						Optional:    true,
					},
					"phone_number": schema.StringAttribute{
						Description: "Contact phone number for the company.",
						Optional:    true,
					},
					"registration_date": schema.StringAttribute{
						Description: "Date of official company registration in YYYY-MM-DD format.",
						Optional:    true,
					},
					"registration_id": schema.StringAttribute{
						Description: "Official company registration identifier.",
						Optional:    true,
					},
					"share_capital": schema.StringAttribute{
						Description: "Declared share capital of the company, usually in euros.",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "Current status of the company (e.g., active, inactive).",
						Optional:    true,
					},
					"tax_identification_number": schema.StringAttribute{
						Description: "National tax identifier (e.g., VAT or TIN).",
						Optional:    true,
					},
					"type": schema.StringAttribute{
						Description: `Type of company, such as "main" or "affiliated".`,
						Optional:    true,
					},
					"website_url": schema.StringAttribute{
						Description: "Company’s official website URL.",
						Optional:    true,
					},
				},
			},
			"source_id": schema.StringAttribute{
				Description: "Optional identifier to track the origin of the request or integration from your system.",
				Optional:    true,
			},
			"technical_data": schema.SingleNestedAttribute{
				Description: "Technical metadata and callback configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"active_aml_suspicions": schema.BoolAttribute{
						Description: "Flag indicating whether there are active research AML (Anti-Money Laundering) suspicions for the company when you apply for a new entry or get an existing one.",
						Optional:    true,
					},
					"callback_url": schema.StringAttribute{
						Description: "URL to receive a callback once the company is processed.",
						Optional:    true,
					},
					"callback_url_notification": schema.StringAttribute{
						Description: "URL to receive notifications about the processing state and status.",
						Optional:    true,
					},
					"language": schema.StringAttribute{
						Description: `Preferred language for responses or notifications (e.g., "eng", "fra").`,
						Optional:    true,
					},
					"raw_data": schema.BoolAttribute{
						Description: "Flag indicating whether to include raw data in the response.",
						Optional:    true,
					},
				},
			},
			"portal_url": schema.StringAttribute{
				Description: "Admin or internal portal URL for viewing the company's details, typically used by internal users.",
				Computed:    true,
			},
			"webview_url": schema.StringAttribute{
				Description: "Public-facing webview URL for the company’s identification process, allowing external access to the company data.",
				Computed:    true,
			},
			"aml_suspicions": schema.ListNestedAttribute{
				Description: "List of AML (Anti-Money Laundering) suspicion entries linked to the company, including their details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyAmlSuspicionsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"caption": schema.StringAttribute{
							Description: "Human-readable description or title for the suspicious finding.",
							Computed:    true,
						},
						"country": schema.StringAttribute{
							Description: "Country associated with the suspicion (ISO 3166-1 alpha-2 code).",
							Computed:    true,
						},
						"gender": schema.StringAttribute{
							Description: "Gender associated with the suspicion, if applicable.",
							Computed:    true,
						},
						"relation": schema.StringAttribute{
							Description: `Nature of the relationship between the entity and the suspicious activity (e.g., "linked", "associated").`,
							Computed:    true,
						},
						"schema": schema.StringAttribute{
							Description: "Version of the evaluation schema or rule engine used.",
							Computed:    true,
						},
						"score": schema.Float64Attribute{
							Description: "Risk score between 0.0 and 1 indicating the severity of the suspicion.",
							Computed:    true,
						},
						"source": schema.StringAttribute{
							Description: "Source system or service providing this suspicion.",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Status of the suspicion review process. Possible values: \"true_positive\", \"false_positive\", \"pending\".\nAvailable values: \"true_positive\", \"false_positive\", \"pending\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"true_positive",
									"false_positive",
									"pending",
								),
							},
						},
						"type": schema.StringAttribute{
							Description: "Category of the suspicion. Possible values: \"crime\", \"sanction\", \"pep\", \"adverse_news\", \"other\".\nAvailable values: \"crime\", \"sanction\", \"pep\", \"adverse_news\", \"other\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"crime",
									"sanction",
									"pep",
									"adverse_news",
									"other",
								),
							},
						},
					},
				},
			},
			"certificat": schema.SingleNestedAttribute{
				Description: "Digital certificate associated with the company, if any, including its creation timestamp and filename.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CompanyCertificatModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique identifier for the certificate.",
						Computed:    true,
					},
					"created_at": schema.StringAttribute{
						Description: "Timestamp when the certificate was created.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"filename": schema.StringAttribute{
						Description: "Name of the certificate file.",
						Computed:    true,
					},
				},
			},
			"checks": schema.ListNestedAttribute{
				Description: "List of verification or validation checks applied to the company, including their results and messages.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyChecksModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"masked": schema.BoolAttribute{
							Description: "Indicates whether the result or data is masked/hidden.",
							Computed:    true,
						},
						"message": schema.StringAttribute{
							Description: "Additional message or explanation about the check result.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name or type of the check performed.",
							Computed:    true,
						},
						"validate": schema.BoolAttribute{
							Description: "Result of the check, true if passed.",
							Computed:    true,
						},
						"weight": schema.Int64Attribute{
							Description: "Importance or weight of the check, often used in scoring.",
							Computed:    true,
						},
					},
				},
			},
			"documents": schema.ListNestedAttribute{
				Description: "All documents submitted or associated with the company, including their metadata and processing status.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier of the document.",
							Computed:    true,
						},
						"checks": schema.ListNestedAttribute{
							Description: "List of verification checks performed on the document.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsChecksModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"masked": schema.BoolAttribute{
										Description: "Indicates whether the result or data is masked/hidden.",
										Computed:    true,
									},
									"message": schema.StringAttribute{
										Description: "Additional message or explanation about the check result.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Name or type of the check performed.",
										Computed:    true,
									},
									"validate": schema.BoolAttribute{
										Description: "Result of the check, true if passed.",
										Computed:    true,
									},
									"weight": schema.Int64Attribute{
										Description: "Importance or weight of the check, often used in scoring.",
										Computed:    true,
									},
								},
							},
						},
						"created_at": schema.StringAttribute{
							Description: "Timestamp when the document was created or uploaded.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"document_type": schema.StringAttribute{
							Description: "Type/category of the document.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name or label for the document.",
							Computed:    true,
						},
						"signed_url": schema.StringAttribute{
							Description: "Signed URL for accessing the document file.",
							Computed:    true,
						},
						"state": schema.StringAttribute{
							Description: "Current processing state of the document (e.g., WAITING, PROCESSED).",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Status of the document reception or approval.",
							Computed:    true,
						},
						"tables": schema.ListNestedAttribute{
							Description: "List of tables extracted from the document, each containing operations.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsTablesModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"operation": schema.ListAttribute{
										Description: "List of operations or actions associated with the table.",
										Computed:    true,
										CustomType:  customfield.NewListType[jsontypes.Normalized](ctx),
										ElementType: jsontypes.NormalizedType{},
									},
								},
							},
						},
						"values": schema.ListNestedAttribute{
							Description: "Extracted key-value pairs from the document, including confidence scores.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsValuesModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"confidence": schema.Float64Attribute{
										Description: "Confidence score (between 0 and 1) for the extracted value.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Name or label of the extracted field.",
										Computed:    true,
									},
									"value": schema.ListAttribute{
										Description: "List of integer values related to the field (e.g., bounding box coordinates).",
										Computed:    true,
										CustomType:  customfield.NewListType[types.Int64](ctx),
										ElementType: types.Int64Type,
									},
								},
							},
						},
					},
				},
			},
			"members": schema.ListNestedAttribute{
				Description: "List of members or actors associated with the company, including personal and ownership information.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyMembersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"address": schema.StringAttribute{
							Description: "Address of the member, which may include street, city, postal code, and country.",
							Computed:    true,
						},
						"birthday": schema.StringAttribute{
							Description: "Birthday (available only if type = person)",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"birthplace": schema.StringAttribute{
							Description: "Birthplace (available only if type = person)",
							Computed:    true,
						},
						"country": schema.StringAttribute{
							Description: `ISO 3166-1 alpha-2 country code of the member's address (e.g., "FR" for France).`,
							Computed:    true,
						},
						"documents": schema.ListNestedAttribute{
							Description: "List of documents associated with the member, including their metadata and processing status.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Unique identifier of the document.",
										Computed:    true,
									},
									"checks": schema.ListNestedAttribute{
										Description: "List of verification checks performed on the document.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsChecksModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"masked": schema.BoolAttribute{
													Description: "Indicates whether the result or data is masked/hidden.",
													Computed:    true,
												},
												"message": schema.StringAttribute{
													Description: "Additional message or explanation about the check result.",
													Computed:    true,
												},
												"name": schema.StringAttribute{
													Description: "Name or type of the check performed.",
													Computed:    true,
												},
												"validate": schema.BoolAttribute{
													Description: "Result of the check, true if passed.",
													Computed:    true,
												},
												"weight": schema.Int64Attribute{
													Description: "Importance or weight of the check, often used in scoring.",
													Computed:    true,
												},
											},
										},
									},
									"created_at": schema.StringAttribute{
										Description: "Timestamp when the document was created or uploaded.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"document_type": schema.StringAttribute{
										Description: "Type/category of the document.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "Name or label for the document.",
										Computed:    true,
									},
									"signed_url": schema.StringAttribute{
										Description: "Signed URL for accessing the document file.",
										Computed:    true,
									},
									"state": schema.StringAttribute{
										Description: "Current processing state of the document (e.g., WAITING, PROCESSED).",
										Computed:    true,
									},
									"status": schema.StringAttribute{
										Description: "Status of the document reception or approval.",
										Computed:    true,
									},
									"tables": schema.ListNestedAttribute{
										Description: "List of tables extracted from the document, each containing operations.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsTablesModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"operation": schema.ListAttribute{
													Description: "List of operations or actions associated with the table.",
													Computed:    true,
													CustomType:  customfield.NewListType[jsontypes.Normalized](ctx),
													ElementType: jsontypes.NormalizedType{},
												},
											},
										},
									},
									"values": schema.ListNestedAttribute{
										Description: "Extracted key-value pairs from the document, including confidence scores.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsValuesModel](ctx),
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"confidence": schema.Float64Attribute{
													Description: "Confidence score (between 0 and 1) for the extracted value.",
													Computed:    true,
												},
												"name": schema.StringAttribute{
													Description: "Name or label of the extracted field.",
													Computed:    true,
												},
												"value": schema.ListAttribute{
													Description: "List of integer values related to the field (e.g., bounding box coordinates).",
													Computed:    true,
													CustomType:  customfield.NewListType[types.Int64](ctx),
													ElementType: types.Int64Type,
												},
											},
										},
									},
								},
							},
						},
						"email": schema.StringAttribute{
							Description: "Email address of the member, which may be used for communication or verification purposes.",
							Computed:    true,
						},
						"first_name": schema.StringAttribute{
							Description: "First name (available only if type = person)",
							Computed:    true,
						},
						"is_beneficial_owner": schema.BoolAttribute{
							Description: "Indicates whether the member is a beneficial owner of the company, meaning they have significant control or ownership.",
							Computed:    true,
						},
						"is_delegator": schema.BoolAttribute{
							Description: "Indicates whether the member is a delegator, meaning they have authority to act on behalf of the company.",
							Computed:    true,
						},
						"last_name": schema.StringAttribute{
							Description: "Last name (available only if type = person)",
							Computed:    true,
						},
						"liveness_verification": schema.BoolAttribute{
							Description: "Indicates whether liveness verification was performed for the member, typically in the context of identity checks.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Company name (available only if type = company)",
							Computed:    true,
						},
						"ownership_percentage": schema.Int64Attribute{
							Description: "Percentage of ownership the member has in the company, expressed as an integer between 0 and 100.",
							Computed:    true,
						},
						"phone_number": schema.StringAttribute{
							Description: "Contact phone number of the member, including country code and area code.",
							Computed:    true,
						},
						"postal_code": schema.StringAttribute{
							Description: "Postal code of the member's address, typically a numeric or alphanumeric code.",
							Computed:    true,
						},
						"registration_id": schema.StringAttribute{
							Description: "Official registration identifier of the member, such as a national ID or company registration number.",
							Computed:    true,
						},
						"relation": schema.StringAttribute{
							Description: `Type of relationship the member has with the company, such as "shareholder", "director", or "beneficial_owner".`,
							Computed:    true,
						},
						"roles": schema.StringAttribute{
							Description: `Role of the member within the company, such as "legal_representative", "director", or "manager".`,
							Computed:    true,
						},
						"source": schema.StringAttribute{
							Description: "Source of the data (e.g., government, user, company)\nAvailable values: \"gouve\", \"user\", \"company\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"gouve",
									"user",
									"company",
								),
							},
						},
						"state": schema.StringAttribute{
							Description: `Current state of the member in the workflow, such as "WAITING", "STARTED", "RUNNING", or "PROCESSED".`,
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: `Status of the member in the system, indicating whether they are approved, pending, or rejected. Possible values include "approved", "need_review", "rejected".`,
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Member type (person or company)\nAvailable values: \"person\", \"company\".",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("person", "company"),
							},
						},
						"workspace_id": schema.StringAttribute{
							Description: "Identifier of the workspace to which the member belongs, used for organizational purposes.",
							Computed:    true,
						},
					},
				},
			},
			"properties": schema.ListNestedAttribute{
				Description: "Custom key-value metadata fields associated with the company, allowing for flexible data storage.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyPropertiesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name/key of the property.",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: "Data type of the property value.",
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Value associated with the property name.",
							Computed:    true,
						},
					},
				},
			},
			"risk": schema.SingleNestedAttribute{
				Description: "Risk assessment associated with the company, including a risk code, reason, and confidence score.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CompanyRiskModel](ctx),
				Attributes: map[string]schema.Attribute{
					"code": schema.StringAttribute{
						Description: "Risk category or code identifier.",
						Computed:    true,
					},
					"reason": schema.StringAttribute{
						Description: "Explanation or justification for the assigned risk.",
						Computed:    true,
					},
					"score": schema.Float64Attribute{
						Description: "Numeric risk score between 0.0 and 1.0 indicating severity or confidence.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *CompanyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *CompanyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
