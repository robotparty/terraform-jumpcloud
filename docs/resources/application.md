---
page_title: "jumpcloud_application Resource - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Provides a resource for adding an Amazon Web Services (AWS) account application. Note: This resource is due to change in future versions to be more generic and allow for adding various applications supported by JumpCloud.
---

# Resource `jumpcloud_application`

Provides a resource for adding an Amazon Web Services (AWS) account application. **Note:** This resource is due to change in future versions to be more generic and allow for adding various applications supported by JumpCloud.



## Schema

### Required

- **aws_session_duration** (String) Value of the `https://aws.amazon.com/SAML/Attributes/SessionDuration` attribute.
- **display_label** (String) Name of the application to display
- **saml_role_attribute** (String) Value of the `https://aws.amazon.com/SAML/Attributes/Role` attribute.
- **sso_url** (String) The SSO URL suffix to use

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **metadata_xml** (String) The JumpCloud metadata XML file.


