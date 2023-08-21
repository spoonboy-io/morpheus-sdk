# # ApiTasksTaskFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_type** | **string** | File Source i.e. &#x60;local&#x60;, &#x60;repository&#x60;, &#x60;url&#x60;. Default is &#x60;local&#x60;. | [default to 'local']
**content** | **string** | File content, the script text. Only required when sourceType is &#x60;local&#x60;. | [optional]
**content_path** | **string** | Content Path, the repo file location or url. Required when sourceType is &#x60;repository&#x60; or &#x60;url&#x60;. | [optional]
**content_ref** | **string** | Content Ref, the branch/tag. Only used when sourceType is &#x60;repository&#x60;. | [optional]
**repository** | [**\OpenAPI\Client\Model\ApiTasksTaskFileRepository**](ApiTasksTaskFileRepository.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
