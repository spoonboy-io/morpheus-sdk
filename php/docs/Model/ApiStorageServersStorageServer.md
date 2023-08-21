# # ApiStorageServersStorageServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | Name |
**type** | **string** | The &#x60;Storage Type&#x60; Code or ID |
**description** | **string** | description | [optional]
**enabled** | **bool** | The enabled flag | [optional] [default to true]
**config** | **object** | Configuration object with parameters that vary by &#x60;type&#x60; |
**visibility** | **string** | private or public | [optional] [default to 'private']
**tenants** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites[]**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) | Array of tenant account ids that are allowed access | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
