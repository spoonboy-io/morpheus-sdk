# MorpheusApi.CredentialTenantUsernameKeypairConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **String** | Credential Type Code | 
**name** | **String** | A unique name scoped to your account for the credential | 
**description** | **String** | Optional Description | [optional] 
**enabled** | **Boolean** | Credential enabled | [optional] [default to true]
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**authPath** | **String** | Tenant name | 
**username** | **String** | Username | 
**authKey** | [**CredentialEmailPrivateKeyConfigAuthKey**](CredentialEmailPrivateKeyConfigAuthKey.md) |  | 



## Enum: TypeEnum


* `tenant-username-keypair` (value: `"tenant-username-keypair"`)




