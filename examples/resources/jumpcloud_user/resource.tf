resource "jumpcloud_user" "john_doe" {
  username   = "john.doe"
  email      = "john.doe@acme.org"
  firstname  = "John Smith"
  lastname   = "Doe"
  enable_mfa = true
}