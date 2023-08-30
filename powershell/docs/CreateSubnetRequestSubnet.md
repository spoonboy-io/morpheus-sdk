# CreateSubnetRequestSubnet
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CreateSubnetRequestSubnetType**](CreateSubnetRequestSubnetType.md) |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object. Settings vary by type. | [optional] 
**Tenants** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner[]**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) | Array of tenant account ID objects that are allowed access | [optional] 
**Visibility** | **String** | private or public | [optional] [default to "private"]
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateSubnetRequestSubnet = Initialize-PSOpenAPIToolsCreateSubnetRequestSubnet  -Type null `
 -Config null `
 -Tenants null `
 -Visibility null `
 -Labels null
```

- Convert the resource to JSON
```powershell
$CreateSubnetRequestSubnet | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

