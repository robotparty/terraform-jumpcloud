---
page_title: "jumpcloud_ldap_directory Data Source - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about the JumpCloud LDAP directory. Each account has a LDAP directory by default and there can only be one LDAP directory.
---

# Data Source `jumpcloud_ldap_directory`

Use this data source to get information about the JumpCloud LDAP directory. Each account has a LDAP directory by default and there can only be one LDAP directory.

## Example Usage

```terraform
data "jumpcloud_ldap_directory" "example" {}
```

## Schema

### Read-only

- **id** (String) The ID of this resource.
- **name** (String) Name of the LDAP directory. Example: `JumpCloud LDAP`.
- **type** (String) The directory type. This will always be `ldap_server`.


