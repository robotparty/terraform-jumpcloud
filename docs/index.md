---
layout: ""
page_title: "Provider: JumpCloud"
description: |-
  The JumpCloud provider provides resources to interact with the JumpCloud API v1 and v2.
---

# JumpCloud Provider

The JumpCloud provider provides resources to interact with the JumpCloud API v1 and v2.

This provider is still under development. Feel free to open an issue or contribute by visiting the [GitHub repository](https://github.com/robotparty/terraform-jumpcloud).

**Note** that due to simplicity this provider does not heavily validate input data except for the most crucial things such as basic type checking. For example, e-mail addresses will not be validate by the provider. However, creating invalid resources is not possible, since the JumpCloud API will reject invalid requests.

## Example Usage

```terraform
terraform {
  required_providers {
    jumpcloud = {
      source = "robotparty/jumpcloud"
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

- **api_key** (String) The admin API key to access JumpCloud resources. Can be passed via `JUMPCLOUD_API_KEY` environment variable.
- **org_id** (String) The JumpCloud organization ID. Can be passed via `JUMPCLOUD_ORG_ID` environment variable.