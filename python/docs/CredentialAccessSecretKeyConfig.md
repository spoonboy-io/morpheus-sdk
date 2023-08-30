# CredentialAccessSecretKeyConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A unique name scoped to your account for the credential | 
**username** | **str** | Access Key | 
**password** | **str** | Secret Key | 
**type** | **str** | Credential Type Code | defaults to "access-key-secret"
**description** | **str** | Optional Description | [optional] 
**enabled** | **bool** | Credential enabled | [optional]  if omitted the server will use the default value of True
**integration** | [**CredentialAccessSecretKeyConfigIntegration**](CredentialAccessSecretKeyConfigIntegration.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


