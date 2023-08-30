# CreateResourcePoolGroup200ResponseResourcePoolGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Mode** | **String** | Pool selection mode. Valid values are &#x60;roundrobin&#x60; or &#x60;availablecapacity&#x60;. | [optional] 
**Pools** | **Int64[]** | Array of Resource Pool IDs | [optional] 
**Tenants** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**ResourcePermission** | [**GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission**](GetResourcePoolGroups200ResponseResourcePoolGroupsInnerResourcePermission.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateResourcePoolGroup200ResponseResourcePoolGroup = Initialize-PSOpenAPIToolsCreateResourcePoolGroup200ResponseResourcePoolGroup  -Id null `
 -Name null `
 -Description null `
 -Visibility null `
 -Mode null `
 -Pools null `
 -Tenants null `
 -ResourcePermission null
```

- Convert the resource to JSON
```powershell
$CreateResourcePoolGroup200ResponseResourcePoolGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

