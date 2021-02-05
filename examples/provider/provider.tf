terraform {
  required_providers {
    jumpcloud = {
      source = "sagewave/jumpcloud"
    }
  }
}

provider "jumpcloud" {
  api_key = "test"
  org_id  = "test"
}