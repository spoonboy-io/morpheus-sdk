# AddTasksRequestTaskFile

File, object specifying type and content, see File Object. This is required for task types that expect a script, having scriptable:true and an optionType of `type:\"file-content\"`. 

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_type** | **str** | File Source i.e. &#x60;local&#x60;, &#x60;repository&#x60;, &#x60;url&#x60;. Default is &#x60;local&#x60;. | defaults to "local"
**content** | **str** | File content, the script text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**content_path** | **str** | Content Path, the repo file location or url. Required when sourceType is &#x60;repository&#x60; or &#x60;url&#x60;. | [optional] 
**content_ref** | **str** | Content Ref, the branch/tag. Only used when sourceType is &#x60;repository&#x60;. | [optional] 
**repository** | [**AddTasksRequestTaskFileRepository**](AddTasksRequestTaskFileRepository.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


