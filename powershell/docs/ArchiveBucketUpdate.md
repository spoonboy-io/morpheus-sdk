# ArchiveBucketUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the archive bucket. Must be globally unique. | [optional] 
**Description** | **String** | A description for the archive bucket | [optional] 
**Visibility** | **String** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**IsPublic** | **Boolean** | Public URL - Set to true to allow anonymous access | [optional] [default to $false]
**Accounts** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ArchiveBucketUpdate = Initialize-PSOpenAPIToolsArchiveBucketUpdate  -Name null `
 -Description null `
 -Visibility null `
 -IsPublic null `
 -Accounts null
```

- Convert the resource to JSON
```powershell
$ArchiveBucketUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

