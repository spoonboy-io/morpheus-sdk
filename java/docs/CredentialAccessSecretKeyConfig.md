

# CredentialAccessSecretKeyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**TypeEnum**](#TypeEnum) | Credential Type Code | 
**name** | **String** | A unique name scoped to your account for the credential | 
**description** | **String** | Optional Description |  [optional]
**enabled** | **Boolean** | Credential enabled |  [optional]
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  |  [optional]
**username** | **String** | Access Key | 
**password** | **String** | Secret Key | 



## Enum: TypeEnum

Name | Value
---- | -----
ACCESS_KEY_SECRET | &quot;access-key-secret&quot;



