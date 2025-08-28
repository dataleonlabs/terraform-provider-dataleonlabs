// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package company_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/dataleonlabs-terraform/internal/services/company"
	"github.com/stainless-sdks/dataleonlabs-terraform/internal/test_helpers"
)

func TestCompanyDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*company.CompanyDataSourceModel)(nil)
	schema := company.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
