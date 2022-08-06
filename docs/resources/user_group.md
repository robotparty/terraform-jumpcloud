---
page_title: "jumpcloud_user_group Resource - terraform-jumpcloud"
subcategory: ""
description: |-
  Provides a JumpCloud user group resource. Refer to the JumpCloud API model https://docs.jumpcloud.com/2.0/models/usergroup for further details.
---

# Resource `jumpcloud_user_group`

Provides a JumpCloud user group resource. Refer to the [JumpCloud API model](https://docs.jumpcloud.com/2.0/models/usergroup) for further details.

## Example Usage

```terraform
resource "jumpcloud_user_group" "example" {
  name = "My User Group"
}
```

## Schema

### Required

- **name** (String) The name of the group. Example: `My Group`.

### Optional

- **attributes** (Block List) TODO (see [below for nested schema](#nestedblock--attributes))
- **id** (String) The ID of this resource.

<a id="nestedblock--attributes"></a>
### Nested Schema for `attributes`

Optional:

- **posix_groups** (String)


