# CredentialUsernamePasswordConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A unique name scoped to your account for the credential | 
**username** | **str** | Username | 
**password** | **str** | Password | 
**type** | **str** | Credential Type Code | defaults to "username-password"
**description** | **str** | Optional Description | [optional] 
**enabled** | **bool** | Credential enabled | [optional]  if omitted the server will use the default value of True
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


