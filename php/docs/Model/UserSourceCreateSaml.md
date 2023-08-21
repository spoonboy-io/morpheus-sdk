# # UserSourceCreateSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**url** | **string** | Login Redirect URL | [optional]
**do_not_include_saml_request** | **bool** | Exclude SAML Request Parameter | [optional] [default to false]
**logout_url** | **string** | Lougout Post URL | [optional]
**saml_signature_mode** | **string** | SAML Request signing mode | [optional] [default to 'NoSignature']
**request509_certificate** | **string** | Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional]
**request_private_key** | **string** | RSA Private Key. Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional]
**do_not_validate_signature** | **bool** | SAML Response Signing Flag | [optional] [default to true]
**public_key** | **string** | Signing Public Key. Only applies when &#x60;doNotValidateSignature&#x3D;true&#x60; | [optional]
**private_key** | **string** | Encryption RSA Private Key | [optional]
**given_name_attribute** | **string** | Given Name Attribute Name | [optional]
**surname_attribute** | **string** | Surname Attribute Name | [optional]
**role_attribute_name** | **string** | Role Attribute Name | [optional]
**required_attribute_value** | **string** | Role Attibute Required Value | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
