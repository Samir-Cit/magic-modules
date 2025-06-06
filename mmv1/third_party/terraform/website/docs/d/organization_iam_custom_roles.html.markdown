---
subcategory: "Cloud Platform"
description: |-
Get information about a Google Cloud Organization IAM Custom Roles.
---

# google_organization_iam_custom_roles

Get information about a Google Cloud Organization IAM Custom Roles.
Note that you must have the `roles/iam.organizationRoleViewer`.
See [the official documentation](https://cloud.google.com/iam/docs/creating-custom-roles)
and [API](https://cloud.google.com/iam/docs/reference/rest/v1/organizations.roles/list).

```hcl
data "google_organization_iam_custom_roles" "example" {
  org_id       = "1234567890"
  show_deleted = true
  view         = "FULL"
}
```

## Argument Reference

The following arguments are supported:

* `org_id` - (Required) The numeric ID of the organization.

* `show_deleted` - (Optional) Include Roles that have been deleted. Defaults to `false`.

* `view` - (Optional) When `"FULL"` is specified, the `permissions` field is returned, which includes a list of all permissions in the role. The default value is `"BASIC"`, which does not return the `permissions`.

## Attributes Reference

The following attributes are exported:

* `roles` - A list of all retrieved custom roles roles. Structure is [defined below](#nested_roles).

<a name="nested_roles"></a>The `roles` block supports:

* `deleted` - The current deleted state of the role.

* `description` - A human-readable description for the role.

* `id` - an identifier for the resource with the format `organizations/{{org_id}}/roles/{{role_id}}`.

* `name` - The name of the role in the format `organizations/{{org_id}}/roles/{{role_id}}`. Like `id`, this field can be used as a reference in other resources such as IAM role bindings.

* `permissions` -  The names of the permissions this role grants when bound in an IAM policy.

* `role_id` - The camel case role id used for this role.

* `stage` - The current launch stage of the role. List of possible stages is [here](https://cloud.google.com/iam/reference/rest/v1/organizations.roles#Role.RoleLaunchStage).

* `title` - A human-readable title for the role.
