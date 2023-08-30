# SpecTemplateUpdateFile

File, object specifying file type and content

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_type** | **str** | File Source i.e. local, repository, url. | [optional]  if omitted the server will use the default value of "local"
**content** | **str** | File content, the template text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**content_path** | **str** | Content Path, the repo file location or url. Required when sourceType is repository or url. | [optional] 
**content_ref** | **str** | Content Ref, the branch/tag. Only used when sourceType is repo. | [optional] 
**repository** | [**SpecTemplateUpdateFileRepository**](SpecTemplateUpdateFileRepository.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


