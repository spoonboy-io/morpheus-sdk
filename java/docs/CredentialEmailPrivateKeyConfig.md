

# CredentialEmailPrivateKeyConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**TypeEnum**](#TypeEnum) | Credential Type Code | 
**name** | **String** | A unique name scoped to your account for the credential | 
**description** | **String** | Optional Description |  [optional]
**enabled** | **Boolean** | Credential enabled |  [optional]
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  |  [optional]
**username** | **String** | Email | 
**authKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](CredentialEmailPrivateKeyConfigAuthKey.md) |  | 



## Enum: TypeEnum

Name | Value
---- | -----
EMAIL_PRIVATE_KEY | &quot;email-private-key&quot;



