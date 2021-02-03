---
page_title: "jumpcloud Provider"
subcategory: ""
description: |-
  
---

# jumpcloud Provider



## Example Usage

```terraform
terraform {
  required_providers {
    jumpcloud = {
      source  = "sagewave.io/sagewave/jumpcloud"
      version = "0.1.0"
    }
  }
}

provider "jumpcloud" {
  api_key = "test"
  org_id  = "test"
}
```

## Schema

### Required

- **api_key** (String) The admin API key to access JumpCloud resources
- **org_id** (String) The JumpCloud organization ID
