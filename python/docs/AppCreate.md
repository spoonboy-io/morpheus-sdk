# AppCreate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**blueprint_id** | [**AppCreateBlueprintId**](AppCreateBlueprintId.md) |  | 
**name** | **str** | A unique name for the app | 
**template_id** | **int** |  | [optional] 
**description** | **str** | Description | [optional] 
**labels** | **[str], none_type** | Array of label strings, can be used for filtering. | [optional] 
**group** | [**AppCreateGroup**](AppCreateGroup.md) |  | [optional] 
**default_cloud** | [**AppCreateDefaultCloud**](AppCreateDefaultCloud.md) |  | [optional] 
**environment** | **str** | Environment code (appContext) | [optional] 
**tiers** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Configuration of app elements | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


