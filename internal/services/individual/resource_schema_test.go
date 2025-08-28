// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package individual_test

import (
	"context"
	"testing"

	"github.com/stainless-sdks/dataleonlabs-terraform/internal/services/individual"
	"github.com/stainless-sdks/dataleonlabs-terraform/internal/test_helpers"
)

func TestIndividualModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*individual.IndividualModel)(nil)
	schema := individual.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
