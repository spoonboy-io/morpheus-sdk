# MorpheusApi.CredentialUsernamePasswordKeypairConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **String** | Credential Type Code | 
**name** | **String** | A unique name scoped to your account for the credential | 
**description** | **String** | Optional Description | [optional] 
**enabled** | **Boolean** | Credential enabled | [optional] [default to true]
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**username** | **String** | Username | 
**password** | **String** | User password, API Key, or applicable secret | 
**authKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](CredentialEmailPrivateKeyConfigAuthKey.md) |  | 



## Enum: TypeEnum


* `username-password-keypair` (value: `"username-password-keypair"`)




