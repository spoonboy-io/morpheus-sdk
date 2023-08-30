# UpdateHostManagedRequestServer

Object containing server configuration parameters

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ssh_username** | **str** | SSH username to use when provisioning | [optional] 
**ssh_password** | **str** | SSH password to use, if not specified the account public key can be used | [optional] 
**server_os** | [**UpdateHostInstallAgentRequestServerServerOs**](UpdateHostInstallAgentRequestServerServerOs.md) |  | [optional] 
**plan** | [**UpdateHostManagedRequestServerPlan**](UpdateHostManagedRequestServerPlan.md) |  | [optional] 
**account** | [**UpdateHostManagedRequestServerAccount**](UpdateHostManagedRequestServerAccount.md) |  | [optional] 
**provision_site_id** | **int** | Specific group to assign the server | [optional] 
**tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. | [optional] 
**config** | [**UpdateHostManagedRequestServerConfig**](UpdateHostManagedRequestServerConfig.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


