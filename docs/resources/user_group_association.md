---
page_title: "jumpcloud_user_group_association Resource - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Provides a resource for associating a JumpCloud user group to objects like SSO applications, G Suite, Office 365, LDAP and more.
---

# Resource `jumpcloud_user_group_association`

Provides a resource for associating a JumpCloud user group to objects like SSO applications, G Suite, Office 365, LDAP and more.

## Example Usage

```terraform
resource "jumpcloud_application" "example" {
  display_label        = "My AWS Account"
  sso_url              = "https://sso.jumpcloud.com/saml2/example-application"
  saml_role_attribute  = "arn:aws:iam::AWS_ACCOUNT_ID:role/MY_ROLE,arn:aws:iam::AWS_ACCOUNT_ID:saml-provider/MY_SAML_PROVIDER"
  aws_session_duration = 432000
}

resource "jumpcloud_user_group" "example" {
  name = "My User Group"
}

resource "jumpcloud_user_group_association" "example" {
  type      = "application"
  group_id  = jumpcloud_user_group.example.id
  object_id = jumpcloud_application.example.id
}
```

## Schema

### Required

- **group_id** (String) The ID of the `resource_user_group` resource.
- **object_id** (String) The ID of the object to associate to the group.
- **type** (String) The type of the object to associate to the given group. Possible values: `active_directory`, `application`, `command`, `g_suite`, `ldap_server`, `office_365`, `policy`, `radius_server`, `system`, `system_group`.

### Optional

- **id** (String) The ID of this resource.


