# Dataleonlabs Terraform Provider

The [Dataleonlabs Terraform provider](https://registry.terraform.io/providers/dataleonlabs/dataleonlabs/latest/docs) provides convenient access to
the Dataleonlabs REST API from Terraform.

It is generated with [Stainless](https://www.stainless.com/).

## Requirements

This provider requires Terraform CLI 1.0 or later. You can [install it for your system](https://developer.hashicorp.com/terraform/install)
on Hashicorp's website.

## Usage

Add the following to your `main.tf` file:

<!-- x-release-please-start-version -->

```hcl
# Declare the provider and version
terraform {
  required_providers {
    dataleonlabs = {
      source  = "dataleonlabs/dataleonlabs"
      version = "~> 0.3.0"
    }
  }
}

# Initialize the provider
provider "dataleonlabs" {
  # API key needed to authorize requests.  You must provide a valid API key in the `Api-Key` header. Get your API key from the Dataleon dashboard.
  api_key = "My API Key" # or set DATALEONLABS_API_KEY env variable
}

# Configure a resource
resource "dataleonlabs_company" "example_company" {

}
```

<!-- x-release-please-end -->

Initialize your project by running `terraform init` in the directory.

Additional examples can be found in the [./examples](./examples) folder within this repository, and you can
refer to the full documentation on [the Terraform Registry](https://registry.terraform.io/providers/dataleonlabs/dataleonlabs/latest/docs).

### Provider Options

When you initialize the provider, the following options are supported. It is recommended to use environment variables for sensitive values like access tokens.
If an environment variable is provided, then the option does not need to be set in the terraform source.

| Property | Environment variable   | Required | Default value |
| -------- | ---------------------- | -------- | ------------- |
| api_key  | `DATALEONLABS_API_KEY` | true     | —             |

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/dataleonlabs/terraform-provider-dataleonlabs/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
