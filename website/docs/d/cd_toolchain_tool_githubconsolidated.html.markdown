---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_githubconsolidated"
description: |-
  Get information about cd_toolchain_tool_githubconsolidated
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_githubconsolidated

~> **Beta:** This data source is in Beta, and is subject to change.

Provides a read-only data source for cd_toolchain_tool_githubconsolidated. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_githubconsolidated" "cd_toolchain_tool_githubconsolidated" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_githubconsolidated.cd_toolchain_tool_githubconsolidated.toolchain_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_githubconsolidated.
* `crn` - (Required, String) Tool CRN.

* `href` - (Required, String) URI representing the tool.

* `name` - (Optional, String) Tool name.

* `parameters` - (Required, List) Parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `api_root_url` - (Optional, String) e.g. https://api.github.example.com.
	* `auto_init` - (Optional, Boolean) Select this checkbox to initialize this repository with a README.
	  * Constraints: The default value is `false`.
	* `enable_traceability` - (Optional, Boolean) Select this check box to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.
	  * Constraints: The default value is `false`.
	* `git_id` - (Optional, String)
	* `has_issues` - (Optional, Boolean) Select this check box to enable GitHub Issues for lightweight issue tracking.
	  * Constraints: The default value is `true`.
	* `integration_owner` - (Optional, String) Select the user which git operations will be performed as.
	* `owner_id` - (Optional, String)
	* `private_repo` - (Optional, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, String)
	* `repo_url` - (Optional, String) Type the URL of the repository that you are linking to.
	* `source_repo_url` - (Optional, String) Type the URL of the repository that you are forking or cloning.
	* `token_url` - (Optional, String) Integration token URL.
	* `type` - (Optional, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.

* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.

* `resource_group_id` - (Required, String) Resource group where tool can be found.

* `state` - (Required, String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (Required, String) CRN of toolchain which the tool is bound to.

* `updated_at` - (Required, String) Latest tool update timestamp.

