# MorpheusApi.ApiStorageServersIdStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | [optional] 
**type** | **String** | The &#x60;Storage Type&#x60; Code or ID | [optional] 
**description** | **String** | description | [optional] 
**enabled** | **Boolean** | The enabled flag | [optional] [default to true]
**config** | **Object** | Configuration object with parameters that vary by &#x60;type&#x60; | [optional] 
**visibility** | **String** | private or public | [optional] [default to &#39;private&#39;]
**tenants** | [**[ApiBlueprintsIdUpdatePermissionsResourcePermissionSites]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access | [optional] 



## Enum: VisibilityEnum


* `private` (value: `"private"`)

* `public` (value: `"public"`)




