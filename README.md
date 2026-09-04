# tf-aws-module_primitive-dns_record

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY-NC-ND 4.0](https://img.shields.io/badge/License-CC_BY--NC--ND_4.0-lightgrey.svg)](https://creativecommons.org/licenses/by-nc-nd/4.0/)

## Overview

This Terraform module creates Route53 DNS records in the hosted zone identified by `zone_id`. It can create the record types supported by Route 53.

## Usage

See [examples/zone-records](examples/zone-records) and [examples/duplicate-records](examples/duplicate-records) for deployable examples.

<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.0 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | ~> 5.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_route53_record.route53_record](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/route53_record) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_records"></a> [records](#input\_records) | Records and their properties | <pre>map(object({<br/>    type            = string<br/>    ttl             = optional(number)<br/>    name            = string<br/>    records         = optional(list(string))<br/>    set_identifier  = optional(string)<br/>    health_check_id = optional(string)<br/>    alias = optional(object({<br/>      name                   = string<br/>      zone_id                = string<br/>      evaluate_target_health = bool<br/>    }))<br/>    cidr_routing_policy = optional(object({<br/>      collection_id = string<br/>      location_name = string<br/>    }))<br/>    failover_routing_policy = optional(object({<br/>      type = string<br/>    }))<br/>    geolocation_routing_policy = optional(object({<br/>      continent   = string<br/>      country     = string<br/>      subdivision = optional(string)<br/>    }))<br/>    latency_routing_policy = optional(object({<br/>      region = string<br/>    }))<br/>    multivalue_answer_routing_policy = optional(bool)<br/>    weighted_routing_policy = optional(object({<br/>      weight = number<br/>    }))<br/>    allow_overwrite = optional(bool)<br/>  }))</pre> | n/a | yes |
| <a name="input_zone_id"></a> [zone\_id](#input\_zone\_id) | (Required) The ID of the hosted zone to contain this record. | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_record_fqdns"></a> [record\_fqdns](#output\_record\_fqdns) | FQDNs built using the zone domain and name. |
<!-- END_TF_DOCS -->

## Known Issues

Encryption in transit is not currently supported by the AWS Terraform provider for this resource. See https://github.com/hashicorp/terraform-provider-aws/issues/26367 and https://github.com/hashicorp/terraform-provider-aws/pull/26987

## Module Development

### Pre-Requisites

The following commands should be available on your system:

- `asdf` or `mise`
- `make`
- `python3` (for pre-commit)

Additionally, your `git` user and email must be configured. Run the `make configure` command from the root of the repository to ensure that you meet these requirements.

### Pre-Commit hooks

The [.pre-commit-config.yaml](.pre-commit-config.yaml) file defines certain `pre-commit` hooks that are relevant to Terraform and Golang, as well as some common linting tasks. These will be configured for you when you run `make configure`.

### Local Validation

You should validate the changes you make to any module locally, prior to pushing your changes in a branch to GitHub.

1. Ensure that you have run `make configure` successfully.

2. Ensure you are signed into the appropriate cloud provider (e.g. AWS or Azure) for the module under test in your current console session.

3. Run the Terraform and Golang linters with the following command:

```
make lint
```

4. Once you have satisfied the linters, the following command will build example infrastructure in your configured cloud, run the tests, and then tear down the infrastructure it created:

```
make test
```

The pre-commit validations, as well as the `make lint` and `make test` targets, will all be performed in CI. Running these validations locally prior to opening a PR helps ensure a smooth review and merge process.

### Review & Merge Process

Once your change has been tested locally and your branch pushed up, open a new Pull Request for your branch to the default (main) branch of this repository.

The title of your Pull Request will determine the version bump for this change, and the title must be in [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/#specification) format in order to merge. A breaking change will trigger a major version bump, a feature will trigger a minor version bump, and all other types will trigger a patch version bump.

Ensure your CI workflows are passing; seek approval from teammates and address any feedback; seek any explicit approvals required by the CODEOWNERS file. You may merge the PR as soon as all requirements are met, and a new release and tag will be automatically created for you.

### Automatic Updates

The shared configuration and workflow files in this repository are largely managed through the [launch-terraform-skeleton](https://github.com/launchbynttdata/launch-terraform-skeleton) repository. Outside of perhaps the `.gitignore` to account for specific files being generated by certain Terraform modules (e.g. Lambda functions), there should not be much cause to update these files on a per-repo basis, and making changes to them individually is discouraged.

If desired, you can check for and run these updates locally in a branch if you have the `copier` tool installed. Some example commands are included below:

```
# Check for updates, optionally checking prerelease versions
copier check-update [--prereleases]

# Run an update, using default answers if there are any. We use tasks, which requires --trust to be set.
copier update --defaults --trust [--prereleases]

# Recopy from the source, and --overwrite all templated files in the process
copier recopy --defaults --trust --overwrite [--prereleases]
```

Automatic updates will run through a scheduled workflow, and if the post-update tests are successful, the Pull Request created will automatically merge. Conflicts in the update or failures to test may leave a Pull Request outstanding, which needs to be addressed by a Launch Engineer.
