# MorpheusApi.SpecTemplateUpdateFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sourceType** | **String** | File Source i.e. local, repository, url. | [optional] [default to &#39;local&#39;]
**content** | **String** | File content, the template text. Only required when sourceType is &#x60;local&#x60;. | [optional] 
**contentPath** | **String** | Content Path, the repo file location or url. Required when sourceType is repository or url. | [optional] 
**contentRef** | **String** | Content Ref, the branch/tag. Only used when sourceType is repo. | [optional] 
**repository** | [**SpecTemplateUpdateFileRepository**](SpecTemplateUpdateFileRepository.md) |  | [optional] 


