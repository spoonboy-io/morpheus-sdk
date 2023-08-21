

# CredentialUsernameAPIKeyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**TypeEnum**](#TypeEnum) | Credential Type Code | 
**name** | **String** | A unique name scoped to your account for the credential | 
**description** | **String** | Optional Description |  [optional]
**enabled** | **Boolean** | Credential enabled |  [optional]
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  |  [optional]
**username** | **String** | Username | 
**password** | **String** | API Key | 



## Enum: TypeEnum

Name | Value
---- | -----
USERNAME_API_KEY | &quot;username-api-key&quot;



