// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package individual

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = (*IndividualResource)(nil)

func (r *IndividualResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}
