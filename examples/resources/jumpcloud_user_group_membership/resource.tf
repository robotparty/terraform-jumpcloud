resource "jumpcloud_user_group" "example" {
  name = "My User Group"
}

resource "jumpcloud_user" "john_doe" {
  username   = "john.doe"
  email      = "john.doe@acme.org"
  firstname  = "John Smith"
  lastname   = "Doe"
  enable_mfa = true
}

resource "jumpcloud_user_group_membership" "example" {
  user_id  = jumpcloud_user.john_doe.id
  group_id = jumpcloud_user_group.example.id
}