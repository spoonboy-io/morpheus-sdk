# CreateSubnetRequestSubnet


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | [**CreateSubnetRequestSubnetType**](CreateSubnetRequestSubnetType.md) |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Configuration object. Settings vary by type. | [optional] 
**tenants** | [**[UpdateBlueprintPermissionsRequestResourcePermissionSitesInner]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of tenant account ID objects that are allowed access | [optional] 
**visibility** | **str** | private or public | [optional]  if omitted the server will use the default value of "private"
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


