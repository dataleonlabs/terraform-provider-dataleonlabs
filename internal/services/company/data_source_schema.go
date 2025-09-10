// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package company

import (
	"context"

	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CompanyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"company_id": schema.StringAttribute{
				Required: true,
			},
			"document": schema.BoolAttribute{
				Description: "Include document signed url",
				Optional:    true,
			},
			"scope": schema.StringAttribute{
				Description: "Scope filter (id or scope)",
				Optional:    true,
			},
			"portal_url": schema.StringAttribute{
				Description: "Admin or internal portal URL for viewing the company's details, typically used by internal users.",
				Computed:    true,
			},
			"source_id": schema.StringAttribute{
				Description: "Optional identifier indicating the source of the company record, useful for tracking or integration purposes.",
				Computed:    true,
			},
			"webview_url": schema.StringAttribute{
				Description: "Public-facing webview URL for the company’s identification process, allowing external access to the company data.",
				Computed:    true,
			},
			"aml_suspicions": schema.ListNestedAttribute{
				Description: "List of AML (Anti-Money Laundering) suspicion entries linked to the company, including their details.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyAmlSuspicionsDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[CompanyCertificatDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectListType[CompanyChecksDataSourceModel](ctx),
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
			"company": schema.SingleNestedAttribute{
				Description: "Main information about the company being registered, including legal name, registration ID, and address.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CompanyCompanyDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description: "Full registered address of the company.",
						Computed:    true,
					},
					"closure_date": schema.StringAttribute{
						Description: "Closure date of the company, if applicable.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"commercial_name": schema.StringAttribute{
						Description: "Trade or commercial name of the company.",
						Computed:    true,
					},
					"contact": schema.SingleNestedAttribute{
						Description: "Contact information for the company, including email, phone number, and address.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[CompanyCompanyContactDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"department": schema.StringAttribute{
								Description: "Department of the contact person.",
								Computed:    true,
							},
							"email": schema.StringAttribute{
								Description: "Email address of the contact person.",
								Computed:    true,
							},
							"first_name": schema.StringAttribute{
								Description: "First name of the contact person.",
								Computed:    true,
							},
							"last_name": schema.StringAttribute{
								Description: "Last name of the contact person.",
								Computed:    true,
							},
							"phone_number": schema.StringAttribute{
								Description: "Phone number of the contact person.",
								Computed:    true,
							},
						},
					},
					"country": schema.StringAttribute{
						Description: "Country code where the company is registered.",
						Computed:    true,
					},
					"email": schema.StringAttribute{
						Description: "Contact email address for the company.",
						Computed:    true,
					},
					"employees": schema.Int64Attribute{
						Description: "Number of employees in the company.",
						Computed:    true,
					},
					"employer_identification_number": schema.StringAttribute{
						Description: "Employer Identification Number (EIN) or equivalent.",
						Computed:    true,
					},
					"insolvency_exists": schema.BoolAttribute{
						Description: "Indicates whether an insolvency procedure exists for the company.",
						Computed:    true,
					},
					"insolvency_ongoing": schema.BoolAttribute{
						Description: "Indicates whether an insolvency procedure is ongoing for the company.",
						Computed:    true,
					},
					"legal_form": schema.StringAttribute{
						Description: "Legal form or structure of the company (e.g., LLC, SARL).",
						Computed:    true,
					},
					"name": schema.StringAttribute{
						Description: "Legal registered name of the company.",
						Computed:    true,
					},
					"phone_number": schema.StringAttribute{
						Description: "Contact phone number for the company, including country code.",
						Computed:    true,
					},
					"registration_date": schema.StringAttribute{
						Description: "Date when the company was officially registered.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"registration_id": schema.StringAttribute{
						Description: "Official company registration number or ID.",
						Computed:    true,
					},
					"share_capital": schema.StringAttribute{
						Description: "Total share capital of the company, including currency.",
						Computed:    true,
					},
					"status": schema.StringAttribute{
						Description: "Current status of the company (e.g., active, inactive).",
						Computed:    true,
					},
					"tax_identification_number": schema.StringAttribute{
						Description: "Tax identification number for the company.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "Type of company within the workspace, e.g., main or affiliated.",
						Computed:    true,
					},
					"website_url": schema.StringAttribute{
						Description: "Official website URL of the company.",
						Computed:    true,
					},
				},
			},
			"documents": schema.ListNestedAttribute{
				Description: "All documents submitted or associated with the company, including their metadata and processing status.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier of the document.",
							Computed:    true,
						},
						"checks": schema.ListNestedAttribute{
							Description: "List of verification checks performed on the document.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsChecksDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsTablesDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[CompanyDocumentsValuesDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectListType[CompanyMembersDataSourceModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Unique identifier of the document.",
										Computed:    true,
									},
									"checks": schema.ListNestedAttribute{
										Description: "List of verification checks performed on the document.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsChecksDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsTablesDataSourceModel](ctx),
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
										CustomType:  customfield.NewNestedObjectListType[CompanyMembersDocumentsValuesDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectListType[CompanyPropertiesDataSourceModel](ctx),
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
				CustomType:  customfield.NewNestedObjectType[CompanyRiskDataSourceModel](ctx),
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
			"technical_data": schema.SingleNestedAttribute{
				Description: "Technical metadata related to the request, such as IP address, QR code settings, and callback URLs.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[CompanyTechnicalDataDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"active_aml_suspicions": schema.BoolAttribute{
						Description: "Flag indicating whether there are active research AML (Anti-Money Laundering) suspicions for the object when you apply for a new entry or get an existing one.",
						Computed:    true,
					},
					"api_version": schema.Int64Attribute{
						Description: "Version number of the API used.",
						Computed:    true,
					},
					"approved_at": schema.StringAttribute{
						Description: "Timestamp when the request or process was approved.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"callback_url": schema.StringAttribute{
						Description: "URL to receive callback data from the AML system.",
						Computed:    true,
					},
					"callback_url_notification": schema.StringAttribute{
						Description: "URL to receive notification updates about the processing status.",
						Computed:    true,
					},
					"disable_notification": schema.BoolAttribute{
						Description: "Flag to indicate if notifications are disabled.",
						Computed:    true,
					},
					"disable_notification_date": schema.StringAttribute{
						Description: "Timestamp when notifications were disabled; null if never disabled.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"export_type": schema.StringAttribute{
						Description: `Export format defined by the API (e.g., "json", "xml").`,
						Computed:    true,
					},
					"filtering_score_aml_suspicions": schema.Float64Attribute{
						Description: "Minimum filtering score (between 0 and 1) for AML suspicions to be considered.",
						Computed:    true,
					},
					"finished_at": schema.StringAttribute{
						Description: "Timestamp when the process finished.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"ip": schema.StringAttribute{
						Description: "IP address of the our system handling the request.",
						Computed:    true,
					},
					"language": schema.StringAttribute{
						Description: `Language preference used in the client workspace (e.g., "fra").`,
						Computed:    true,
					},
					"location_ip": schema.StringAttribute{
						Description: "IP address of the end client (final user) captured.",
						Computed:    true,
					},
					"need_review_at": schema.StringAttribute{
						Description: "Timestamp indicating when the request or process needs review; null if none.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"notification_confirmation": schema.BoolAttribute{
						Description: "Flag indicating if notification confirmation is required or received.",
						Computed:    true,
					},
					"qr_code": schema.StringAttribute{
						Description: `Indicates whether QR code is enabled ("true" or "false").`,
						Computed:    true,
					},
					"raw_data": schema.BoolAttribute{
						Description: "Flag indicating whether to include raw data in the response.",
						Computed:    true,
					},
					"rejected_at": schema.StringAttribute{
						Description: "Timestamp when the request or process was rejected; null if not rejected.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"session_duration": schema.Int64Attribute{
						Description: "Duration of the user session in seconds.",
						Computed:    true,
					},
					"started_at": schema.StringAttribute{
						Description: "Timestamp when the process started.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"transfer_at": schema.StringAttribute{
						Description: "Date/time of data transfer.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"transfer_mode": schema.StringAttribute{
						Description: "Mode of data transfer.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *CompanyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CompanyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
