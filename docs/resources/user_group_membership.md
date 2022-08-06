---
page_title: "jumpcloud_user_group_membership Resource - terraform-jumpcloud"
subcategory: ""
description: |-
  Provides a resource for managing user group memberships.
---

# Resource `jumpcloud_user_group_membership`

Provides a resource for managing user group memberships.

## Example Usage

```terraform
resource "jumpcloud_user_group" "example" {
  name = "My User Group"
}

resource "jumpcloud_user" "john_doe" {
  username   = "john.doe"
  email      = "john.doe@acme.org"
  firstname  = "John Smith"
  lastname   = "Doe"
  enable_mfa = true
}

resource "jumpcloud_user_group_membership" "example" {
  user_id  = jumpcloud_user.john_doe.id
  group_id = jumpcloud_user_group.example.id
}
```

## Schema

### Required

- **group_id** (String) The ID of the `resource_user_group` object.
- **user_id** (String) The ID of the `resource_user` object.

### Optional

- **id** (String) The ID of this resource.


