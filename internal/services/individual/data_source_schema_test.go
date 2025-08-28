// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package individual_test

import (
	"context"
	"testing"

	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/services/individual"
	"github.com/dataleonlabs/terraform-provider-dataleonlabs/internal/test_helpers"
)

func TestIndividualDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*individual.IndividualDataSourceModel)(nil)
	schema := individual.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
