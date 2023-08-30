# AddCredentialsRequestCredential

Payload for creating a new credential

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | Optional Description | [optional] 
**enabled** | **bool** | Credential enabled | [optional]  if omitted the server will use the default value of True
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**type** | **str** | Credential Type Code | [optional]  if omitted the server will use the default value of "oauth2"
**name** | **str** | A unique name scoped to your account for the credential | [optional] 
**username** | **str** | Username | [optional] 
**password** | **str** | Password | [optional] 
**auth_key** | [**CredentialEmailPrivateKeyConfigAuthKey**](CredentialEmailPrivateKeyConfigAuthKey.md) |  | [optional] 
**auth_path** | **str** | Tenant name | [optional] 
**config** | [**CredentialOauth2ConfigConfig**](CredentialOauth2ConfigConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


