---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_keyprotect"
description: |-
  Manages cd_toolchain_tool_keyprotect.
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_keyprotect

~> **Beta:** This resource is in Beta, and is subject to change.

Provides a resource for cd_toolchain_tool_keyprotect. This allows cd_toolchain_tool_keyprotect to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_keyprotect" "cd_toolchain_tool_keyprotect" {
  parameters {
		name = "name"
		region = "region"
		resource_group = "resource-group"
		instance_name = "instance-name"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Required, List) Parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `instance_name` - (Required, String) The name of your Key Protect instance. You should choose an entry from the list provided based on the selected region and resource group. e.g: Key Protect-01.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `name` - (Required, String) Enter a name for this tool integration. This name is displayed on your toolchain.
	* `region` - (Required, String) Region.
	* `resource_group` - (Required, String) Resource group.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_keyprotect.
* `crn` - (Required, String) Tool CRN.
* `tool_id` - (Required, String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `href` - (Required, String) URI representing the tool.
* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.
* `resource_group_id` - (Required, String) Resource group where tool can be found.
* `state` - (Required, String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (Required, String) CRN of toolchain which the tool is bound to.
* `updated_at` - (Required, String) Latest tool update timestamp.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cd_toolchain_tool_keyprotect` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_keyprotect.cd_toolchain_tool_keyprotect <toolchain_id>/<tool_id>
```
