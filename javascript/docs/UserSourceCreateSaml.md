# MorpheusApi.UserSourceCreateSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**url** | **String** | Login Redirect URL | [optional] 
**doNotIncludeSAMLRequest** | **Boolean** | Exclude SAML Request Parameter | [optional] [default to false]
**logoutUrl** | **String** | Lougout Post URL | [optional] 
**sAMLSignatureMode** | **String** | SAML Request signing mode | [optional] [default to &#39;NoSignature&#39;]
**request509Certificate** | **String** | Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional] 
**requestPrivateKey** | **String** | RSA Private Key. Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional] 
**doNotValidateSignature** | **Boolean** | SAML Response Signing Flag | [optional] [default to true]
**publicKey** | **String** | Signing Public Key. Only applies when &#x60;doNotValidateSignature&#x3D;true&#x60; | [optional] 
**privateKey** | **String** | Encryption RSA Private Key | [optional] 
**givenNameAttribute** | **String** | Given Name Attribute Name | [optional] 
**surnameAttribute** | **String** | Surname Attribute Name | [optional] 
**roleAttributeName** | **String** | Role Attribute Name | [optional] 
**requiredAttributeValue** | **String** | Role Attibute Required Value | [optional] 



## Enum: SAMLSignatureModeEnum


* `NoSignature` (value: `"NoSignature"`)

* `SelfSigned` (value: `"SelfSigned"`)

* `CustomSignature` (value: `"CustomSignature"`)




