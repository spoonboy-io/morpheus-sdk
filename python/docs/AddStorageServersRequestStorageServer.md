# AddStorageServersRequestStorageServer


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Name | 
**type** | **str** | The &#x60;Storage Type&#x60; Code or ID | 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Configuration object with parameters that vary by &#x60;type&#x60; | 
**description** | **str** | description | [optional] 
**enabled** | **bool** | The enabled flag | [optional]  if omitted the server will use the default value of True
**visibility** | **str** | private or public | [optional]  if omitted the server will use the default value of "private"
**tenants** | [**[UpdateBlueprintPermissionsRequestResourcePermissionSitesInner]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of tenant account ids that are allowed access | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


