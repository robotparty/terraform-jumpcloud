resource "jumpcloud_application" "example" {
  display_label        = "My AWS Account"
  sso_url              = "https://sso.jumpcloud.com/saml2/example-application"
  saml_role_attribute  = "arn:aws:iam::AWS_ACCOUNT_ID:role/MY_ROLE,arn:aws:iam::AWS_ACCOUNT_ID:saml-provider/MY_SAML_PROVIDER"
  aws_session_duration = 432000
}

resource "jumpcloud_user_group" "example" {
  name = "My User Group"
}

resource "jumpcloud_user_group_association" "example" {
  type      = "application"
  group_id  = jumpcloud_user_group.example.id
  object_id = jumpcloud_application.example.id
}