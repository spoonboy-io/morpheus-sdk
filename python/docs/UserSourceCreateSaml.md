# UserSourceCreateSaml

SAML Configuration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**url** | **str** | Login Redirect URL | [optional] 
**do_not_include_saml_request** | **bool** | Exclude SAML Request Parameter | [optional]  if omitted the server will use the default value of False
**logout_url** | **str** | Lougout Post URL | [optional] 
**saml_signature_mode** | **str** | SAML Request signing mode | [optional]  if omitted the server will use the default value of "NoSignature"
**request509_certificate** | **str** | Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional] 
**request_private_key** | **str** | RSA Private Key. Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional] 
**do_not_validate_signature** | **bool** | SAML Response Signing Flag | [optional]  if omitted the server will use the default value of True
**public_key** | **str** | Signing Public Key. Only applies when &#x60;doNotValidateSignature&#x3D;true&#x60; | [optional] 
**private_key** | **str** | Encryption RSA Private Key | [optional] 
**given_name_attribute** | **str** | Given Name Attribute Name | [optional] 
**surname_attribute** | **str** | Surname Attribute Name | [optional] 
**role_attribute_name** | **str** | Role Attribute Name | [optional] 
**required_attribute_value** | **str** | Role Attibute Required Value | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


