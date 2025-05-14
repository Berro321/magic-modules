---
subcategory: "BeyondCorp"
description: |-
  Get information about a Google BeyondCorp Security Gateway Application.
---

# google_beyondcorp_security_gateway_application

Get information about a Google BeyondCorp Security Gateway Application.

## Example Usage

```hcl
data "google_beyondcorp_security_gateway_application" "my-beyondcorp-security-gateway-application" {
  application_id = "my-beyondcorp-security-gateway-application"
  security_gateways_id = "my-beyondcorp-security-gateway"
}
```

## Argument Reference

The following arguments are supported:

* `application_id` - (Required) The name of the Security Gateway Application resource.

* `security_gateways_id` - (Required) The name of the Security Gateway resource.

- - -

* `project` - (Optional) The project in which the resource belongs. If it
    is not provided, the provider project is used.

## Attributes Reference

See [google_beyondcorp_security_gateway_application](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/beyondcorp_security_gateway_application) resource for details of the available attributes.