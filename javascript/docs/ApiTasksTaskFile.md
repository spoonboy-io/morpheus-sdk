# MorpheusApi.ApiTasksTaskFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sourceType** | **String** | File Source i.e. &#x60;local&#x60;, &#x60;repository&#x60;, &#x60;url&#x60;. Default is &#x60;local&#x60;. | [default to &#39;local&#39;]
**content** | **String** | File content, the script text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**contentPath** | **String** | Content Path, the repo file location or url. Required when sourceType is &#x60;repository&#x60; or &#x60;url&#x60;. | [optional] 
**contentRef** | **String** | Content Ref, the branch/tag. Only used when sourceType is &#x60;repository&#x60;. | [optional] 
**repository** | [**ApiTasksTaskFileRepository**](ApiTasksTaskFileRepository.md) |  | [optional] 



## Enum: SourceTypeEnum


* `local` (value: `"local"`)

* `repository` (value: `"repository"`)

* `url` (value: `"url"`)




