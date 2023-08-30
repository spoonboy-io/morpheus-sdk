# AddStorageServersRequestStorageServer
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | 
**Type** | **String** | The &#x60;Storage Type&#x60; Code or ID | 
**Description** | **String** | description | [optional] 
**Enabled** | **Boolean** | The enabled flag | [optional] [default to $true]
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by &#x60;type&#x60; | 
**Visibility** | **String** | private or public | [optional] [default to "private"]
**Tenants** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner[]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of tenant account ids that are allowed access | [optional] 

## Examples

- Prepare the resource
```powershell
$AddStorageServersRequestStorageServer = Initialize-PSOpenAPIToolsAddStorageServersRequestStorageServer  -Name null `
 -Type null `
 -Description null `
 -Enabled null `
 -Config null `
 -Visibility null `
 -Tenants null
```

- Convert the resource to JSON
```powershell
$AddStorageServersRequestStorageServer | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

