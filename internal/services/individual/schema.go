// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package individual

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

var _ resource.ResourceWithConfigValidators = (*IndividualResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Unique identifier of the individual.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"workspace_id": schema.StringAttribute{
				Description: "Unique identifier of the workspace where the individual is being registered.",
				Required:    true,
			},
			"source_id": schema.StringAttribute{
				Description: "Optional identifier for tracking the source system or integration from your system.",
				Optional:    true,
			},
			"person": schema.SingleNestedAttribute{
				Description: "Personal information about the individual.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"birthday": schema.StringAttribute{
						Description: "Date of birth in DD/MM/YYYY format.",
						Optional:    true,
					},
					"email": schema.StringAttribute{
						Description: "Email address of the individual.",
						Optional:    true,
					},
					"first_name": schema.StringAttribute{
						Description: "First name of the individual.",
						Optional:    true,
					},
					"gender": schema.StringAttribute{
						Description: "Gender of the individual (M for male, F for female).\nAvailable values: \"M\", \"F\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("M", "F"),
						},
					},
					"last_name": schema.StringAttribute{
						Description: "Last name (family name) of the individual.",
						Optional:    true,
					},
					"maiden_name": schema.StringAttribute{
						Description: "Maiden name, if applicable.",
						Optional:    true,
					},
					"nationality": schema.StringAttribute{
						Description: "Nationality of the individual (ISO 3166-1 alpha-3 country code).",
						Optional:    true,
					},
					"phone_number": schema.StringAttribute{
						Description: "Phone number of the individual.",
						Optional:    true,
					},
				},
			},
			"technical_data": schema.SingleNestedAttribute{
				Description: "Technical metadata related to the request or processing.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"active_aml_suspicions": schema.BoolAttribute{
						Description: "Flag indicating whether there are active research AML (Anti-Money Laundering) suspicions for the individual when you apply for a new entry or get an existing one.",
						Optional:    true,
					},
					"callback_url": schema.StringAttribute{
						Description: "URL to call back upon completion of processing.",
						Optional:    true,
					},
					"callback_url_notification": schema.StringAttribute{
						Description: "URL for receive notifications about the processing state or status.",
						Optional:    true,
					},
					"filtering_score_aml_suspicions": schema.Float64Attribute{
						Description: "Minimum filtering score (between 0 and 1) for AML suspicions to be considered.",
						Optional:    true,
					},
					"language": schema.StringAttribute{
						Description: `Preferred language for communication (e.g., "eng", "fra").`,
						Optional:    true,
					},
					"raw_data": schema.BoolAttribute{
						Description: "Flag indicating whether to include raw data in the response.",
						Optional:    true,
					},
				},
			},
			"auth_url": schema.StringAttribute{
				Description: "URL to authenticate the individual, usually for document signing or onboarding.",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Timestamp of the individual's creation in ISO 8601 format.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"number": schema.Int64Attribute{
				Description: "Internal sequential number or reference for the individual.",
				Computed:    true,
			},
			"portal_url": schema.StringAttribute{
				Description: "Admin or internal portal URL for viewing the individual's details.",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "Current operational state in the workflow (e.g., WAITING, IN_PROGRESS, COMPLETED).",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Overall processing status of the individual (e.g., rejected, need_review, approved).",
				Computed:    true,
			},
			"webview_url": schema.StringAttribute{
				Description: "Public-facing webview URL for the individual’s identification process.",
				Computed:    true,
			},
			"aml_suspicions": schema.ListNestedAttribute{
				Description: "List of AML (Anti-Money Laundering) suspicion entries linked to the individual.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[IndividualAmlSuspicionsModel](ctx),
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
				Description: "Digital certificate associated with the individual, if any.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[IndividualCertificatModel](ctx),
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
				Description: "List of verification or validation checks applied to the individual.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[IndividualChecksModel](ctx),
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
				Description: "All documents submitted or associated with the individual.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[IndividualDocumentsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier of the document.",
							Computed:    true,
						},
						"checks": schema.ListNestedAttribute{
							Description: "List of verification checks performed on the document.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[IndividualDocumentsChecksModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[IndividualDocumentsTablesModel](ctx),
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
							CustomType:  customfield.NewNestedObjectListType[IndividualDocumentsValuesModel](ctx),
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
			"identity_card": schema.SingleNestedAttribute{
				Description: "Reference to the individual's identity document.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[IndividualIdentityCardModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Unique identifier for the document.",
						Computed:    true,
					},
					"back_document_signed_url": schema.StringAttribute{
						Description: "Signed URL linking to the back image of the document.",
						Computed:    true,
					},
					"birth_place": schema.StringAttribute{
						Description: "Place of birth as indicated on the document.",
						Computed:    true,
					},
					"birthday": schema.StringAttribute{
						Description: "Date of birth in DD/MM/YYYY format as shown on the document.",
						Computed:    true,
					},
					"country": schema.StringAttribute{
						Description: "Country code issuing the document (ISO 3166-1 alpha-2).",
						Computed:    true,
					},
					"expiration_date": schema.StringAttribute{
						Description: "Expiration date of the document, in YYYY-MM-DD format.",
						Computed:    true,
					},
					"first_name": schema.StringAttribute{
						Description: "First name as shown on the document.",
						Computed:    true,
					},
					"front_document_signed_url": schema.StringAttribute{
						Description: "Signed URL linking to the front image of the document.",
						Computed:    true,
					},
					"gender": schema.StringAttribute{
						Description: `Gender indicated on the document (e.g., "M" or "F").`,
						Computed:    true,
					},
					"issue_date": schema.StringAttribute{
						Description: "Date when the document was issued, in YYYY-MM-DD format.",
						Computed:    true,
					},
					"last_name": schema.StringAttribute{
						Description: "Last name as shown on the document.",
						Computed:    true,
					},
					"mrz_line_1": schema.StringAttribute{
						Description: "First line of the Machine Readable Zone (MRZ) on the document.",
						Computed:    true,
					},
					"mrz_line_2": schema.StringAttribute{
						Description: "Second line of the MRZ on the document.",
						Computed:    true,
					},
					"mrz_line_3": schema.StringAttribute{
						Description: "Third line of the MRZ if applicable; otherwise null.",
						Computed:    true,
					},
					"type": schema.StringAttribute{
						Description: "Type of document (e.g., passport, identity card).",
						Computed:    true,
					},
				},
			},
			"properties": schema.ListNestedAttribute{
				Description: "Custom key-value metadata fields associated with the individual.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[IndividualPropertiesModel](ctx),
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
				Description: "Risk assessment associated with the individual.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[IndividualRiskModel](ctx),
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
			"tags": schema.ListNestedAttribute{
				Description: "List of tags assigned to the individual for categorization or metadata purposes.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[IndividualTagsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "Name of the tag used to identify the metadata field.",
							Computed:    true,
						},
						"private": schema.BoolAttribute{
							Description: "Indicates whether the tag is private (not visible to external users).",
							Computed:    true,
						},
						"type": schema.StringAttribute{
							Description: `Data type of the tag value (e.g., "string", "number", "boolean").`,
							Computed:    true,
						},
						"value": schema.StringAttribute{
							Description: "Value assigned to the tag.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *IndividualResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *IndividualResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
