

# ApiStorageServersStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name | 
**type** | **String** | The &#x60;Storage Type&#x60; Code or ID | 
**description** | **String** | description |  [optional]
**enabled** | **Boolean** | The enabled flag |  [optional]
**config** | **Object** | Configuration object with parameters that vary by &#x60;type&#x60; | 
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | private or public |  [optional]
**tenants** | [**List&lt;ApiBlueprintsIdUpdatePermissionsResourcePermissionSites&gt;**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access |  [optional]



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



