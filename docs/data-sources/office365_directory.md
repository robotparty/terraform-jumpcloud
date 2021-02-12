---
page_title: "jumpcloud_office365_directory Data Source - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about a JumpCloud Office 365 directory.
---

# Data Source `jumpcloud_office365_directory`

Use this data source to get information about a JumpCloud Office 365 directory.

## Example Usage

```terraform
data "jumpcloud_office365_directory" "example" {
  name = "My Office 365 Directory"
}
```

## Schema

### Required

- **name** (String) The user defined name, e.g. `My G Suite directory`.

### Read-only

- **id** (String) The ID of this resource.
- **type** (String) The directory type. This will always be `office_365`.


