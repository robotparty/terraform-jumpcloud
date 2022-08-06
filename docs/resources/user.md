---
page_title: "jumpcloud_user Resource - terraform-jumpcloud"
subcategory: ""
description: |-
  Provides a JumpCloud system user resource. For additional information refer also to the JumpCloud API user model https://docs.jumpcloud.com/1.0/models/systemuserpost.
---

# Resource `jumpcloud_user`

Provides a JumpCloud system user resource. For additional information refer also to the [JumpCloud API user model](https://docs.jumpcloud.com/1.0/models/systemuserpost).

## Example Usage

```terraform
resource "jumpcloud_user" "john_doe" {
  username   = "john.doe"
  email      = "john.doe@acme.org"
  firstname  = "John Smith"
  lastname   = "Doe"
  enable_mfa = true
}
```

## Schema

### Required

- **email** (String) The users e-mail address, which is also used for log ins. E-mail addresses have to be unique across all JumpCloud accounts, there cannot be two users with the same e-mail address. Example: `john.doe@acme.org`.
- **username** (String) The technical user name. See JumpCloud's [user naming conventions](https://support.jumpcloud.com/support/s/article/naming-convention-for-users1) for naming restrictions. Example: `john.doe`.

### Optional

- **enable_mfa** (Boolean) Require Multi-factor Authentication on the User Portal.
- **firstname** (String) The user's first name. Example: `john`.
- **id** (String) The ID of this resource.
- **lastname** (String) The user's last name. Example: `doe`.


