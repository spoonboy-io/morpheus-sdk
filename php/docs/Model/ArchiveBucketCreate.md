# # ArchiveBucketCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **string** | A name for the archive bucket. Must be globally unique. | [optional]
**description** | **string** | A description for the archive bucket | [optional]
**storage_provider** | [**\OpenAPI\Client\Model\ArchiveBucketCreateStorageProvider**](ArchiveBucketCreateStorageProvider.md) |  | [optional]
**visibility** | **string** | Visibility - Set to public to allow all tenants | [optional] [default to 'private']
**is_public** | **bool** | Public URL - Set to true to allow anonymous access | [optional] [default to false]
**accounts** | [**\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
