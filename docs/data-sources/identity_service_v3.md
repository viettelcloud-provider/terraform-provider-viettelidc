---
subcategory: "Identity / Keystone"
layout: "openstack"
page_title: "ViettelIdc: viettelidc_identity_service_v3"
sidebar_current: "docs-openstack-datasource-identity-service-v3"
description: |-
  Get information on an ViettelIdc Service.
---

# viettelidc\_identity\_service\_v3

Use this data source to get the ID of an ViettelIdc service.

~> **Note:** This usually requires admin privileges.

## Example Usage

```hcl
data "viettelidc_identity_service_v3" "service_1" {
  name = "keystone"
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional) The region in which to obtain the V3 Keystone client.
  If omitted, the `region` argument of the provider is used.

* `name` - (Optional) The service name.

* `type` - (Optional) The service type.

* `enabled` - (Optional) The service status.

## Attributes Reference

`id` is set to the ID of the found service. In addition, the following attributes
are exported:

* `region` - See Argument Reference above.
* `name` - See Argument Reference above.
* `type` - See Argument Reference above.
* `enabled` - See Argument Reference above.
* `description` - The service description.
