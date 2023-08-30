# AddCloudResourcePool200ResponseResourcePool
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Description** | **String** |  | [optional] 
**Zone** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Parent** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**Type** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**RegionCode** | **String** |  | [optional] 
**IacId** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**ReadOnly** | **Boolean** |  | [optional] 
**DefaultPool** | **Boolean** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**Status** | **String** |  | [optional] 
**Inventory** | **Boolean** |  | [optional] 
**Config** | [**ZoneResourcePoolConfig**](ZoneResourcePoolConfig.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**DisplayName** | **String** |  | [optional] 
**Tenants** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**ResourcePermission** | [**ResourcePermissions**](ResourcePermissions.md) |  | [optional] 
**Depth** | **Int64** |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCloudResourcePool200ResponseResourcePool = Initialize-PSOpenAPIToolsAddCloudResourcePool200ResponseResourcePool  -Id null `
 -Description null `
 -Zone null `
 -Parent null `
 -Type null `
 -ExternalId null `
 -RegionCode null `
 -IacId null `
 -Visibility null `
 -ReadOnly null `
 -DefaultPool null `
 -Active null `
 -Status null `
 -Inventory null `
 -Config null `
 -Name null `
 -DisplayName null `
 -Tenants null `
 -ResourcePermission null `
 -Depth null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddCloudResourcePool200ResponseResourcePool | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

