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