

# CredentialOauth2Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**TypeEnum**](#TypeEnum) | Credential Type Code | 
**name** | **String** | A unique name scoped to your account for the credential | 
**description** | **String** | Optional Description |  [optional]
**enabled** | **Boolean** | Credential enabled |  [optional]
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  |  [optional]
**username** | **String** | Username |  [optional]
**password** | **String** | Password |  [optional]
**config** | [**CredentialOauth2ConfigConfig**](CredentialOauth2ConfigConfig.md) |  | 



## Enum: TypeEnum

Name | Value
---- | -----
OAUTH2 | &quot;oauth2&quot;



