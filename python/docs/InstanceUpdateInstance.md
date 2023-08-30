# InstanceUpdateInstance


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Unique name scoped to your account for the instance. | [optional] 
**description** | **str** | Optional description field. | [optional] 
**instance_context** | **str** | Environment | [optional] 
**labels** | **[str]** | Array of strings (keywords). | [optional] 
**tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**add_tags** | [**[UpdateHostManagedRequestServerTagsInner]**](UpdateHostManagedRequestServerTagsInner.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**remove_tags** | [**[InstanceUpdateInstanceRemoveTagsInner]**](InstanceUpdateInstanceRemoveTagsInner.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**power_schedule_type** | **int** | Power schedule ID. | [optional] 
**site** | [**InstanceUpdateInstanceSite**](InstanceUpdateInstanceSite.md) |  | [optional] 
**owner_id** | **int** | User ID, can be used to change instance owner. | [optional] 
**display_name** | **str** | Name used in the UI for display | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


