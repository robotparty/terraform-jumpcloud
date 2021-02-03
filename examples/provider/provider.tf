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