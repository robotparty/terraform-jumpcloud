---
page_title: "jumpcloud_ldap_directory Data Source - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about the JumpCloud LDAP directory.
---

# Data Source `jumpcloud_ldap_directory`

Use this data source to get information about the JumpCloud LDAP directory.



## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **name** (String) Name of the LDAP directory. Example: `JumpCloud LDAP`.
- **type** (String) The directory type. This will always be `ldap_server`.


