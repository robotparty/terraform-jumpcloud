---
page_title: "jumpcloud_gsuite_directory Data Source - terraform-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about a JumpCloud G Suite directory.
---

# Data Source `jumpcloud_gsuite_directory`

Use this data source to get information about a JumpCloud G Suite directory.

## Example Usage

```terraform
data "jumpcloud_gsuite_directory" "example" {
  name = "My G Suite Directory"
}
```

## Schema

### Required

- **name** (String) The user defined name, e.g. `My G Suite directory`.

### Read-only

- **id** (String) The ID of this resource.
- **type** (String) The directory type. This will always be `g_suite`.


