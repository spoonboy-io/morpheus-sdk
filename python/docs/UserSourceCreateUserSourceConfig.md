# UserSourceCreateUserSourceConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**url** | **str** | Login Redirect URL | [optional] 
**binding_username** | **str** | Binding Username | [optional] 
**binding_password** | **str** | Binding Password | [optional] 
**required_group** | **str** | Required Group | [optional] 
**organization_id** | **bool** | Organization ID | [optional]  if omitted the server will use the default value of False
**required_role** | **str** | Required Role | [optional] 
**domain** | **str** | Domain | [optional] 
**use_ssl** | **bool** | Use SSL | [optional]  if omitted the server will use the default value of False
**search_member_groups** | **bool** | Include Member Groups | [optional]  if omitted the server will use the default value of False
**administrator_api_token** | **str** | Administrator API Token | [optional] 
**subdomain** | **str** | OneLogin Subdomain | [optional] 
**region** | **str** | OneLogin Region | [optional] 
**client_secret** | **str** | API Client Secret | [optional] 
**client_id** | **str** | API Client ID | [optional] 
**do_not_include_saml_request** | **bool** | Do not include SAMLRequest | [optional]  if omitted the server will use the default value of False
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
**login_url** | **str** | External Login URL | [optional] 
**logout** | **str** | External Logout URL | [optional] 
**encryption_algo** | **str** | Encryption Algorithm | [optional] 
**encryption_key** | **str** | Encryption Key | [optional] 
**endpoint** | **str** | API Endpoint | [optional] 
**api_style** | **str** | API Style | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


