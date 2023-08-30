# UpdateCloudResourcePoolRequestResourcePool


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active** | **bool** | Activate &#x60;true&#x60; or disable &#x60;false&#x60; the datastore | [optional] 
**visibility** | **str** | Setting &#x60;private&#x60; or &#x60;public&#x60; | [optional]  if omitted the server will use the default value of "private"
**display_name** | **str** | Optional Display Name (VMware only) | [optional] 
**inventory** | **bool** | Enable &#x60;True&#x60; or disable &#x60;False&#x60; inventory sync for resource pool during cloud refresh | [optional] 
**tenant_permissions** | [**[UpdateCloudFoldersRequestFolderTenantPermissionsInner]**](UpdateCloudFoldersRequestFolderTenantPermissionsInner.md) |  | [optional] 
**resource_permissions** | [**UpdateCloudDatastoresRequestDatastoreResourcePermissions**](UpdateCloudDatastoresRequestDatastoreResourcePermissions.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


