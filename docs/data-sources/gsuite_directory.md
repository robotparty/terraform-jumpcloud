---
page_title: "jumpcloud_gsuite_directory Data Source - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about a JumpCloud G Suite directory.
---

# Data Source `jumpcloud_gsuite_directory`

Use this data source to get information about a JumpCloud G Suite directory.

## Example Usage

```terraform
data "jumpcloud_gsuite_directory" "example" {}
```

## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **name** (String) The user defined name, e.g. `My G Suite directory`.
- **type** (String) The directory type. This will always be `g_suite`.


