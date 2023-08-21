# # SpecTemplateUpdateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**source_type** | **string** | File Source i.e. local, repository, url. | [optional] [default to 'local']
**content** | **string** | File content, the template text. Only required when sourceType is &#x60;local&#x60;. | [optional]
**content_path** | **string** | Content Path, the repo file location or url. Required when sourceType is repository or url. | [optional]
**content_ref** | **string** | Content Ref, the branch/tag. Only used when sourceType is repo. | [optional]
**repository** | [**\OpenAPI\Client\Model\SpecTemplateUpdateFileRepository**](SpecTemplateUpdateFileRepository.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
