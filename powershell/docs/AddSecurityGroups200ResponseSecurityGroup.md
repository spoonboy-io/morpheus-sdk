# AddSecurityGroups200ResponseSecurityGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**GroupSource** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Enabled** | **String** |  | [optional] 
**SyncSource** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**Zone** | [**ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert**](ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerSslCert.md) |  | [optional] 
**Locations** | [**SecurityGroupLocationsInner[]**](SecurityGroupLocationsInner.md) |  | [optional] 
**Rules** | [**SecurityGroupRulesInner[]**](SecurityGroupRulesInner.md) |  | [optional] 
**Tenants** | [**SecurityGroupTenantsInner[]**](SecurityGroupTenantsInner.md) |  | [optional] 
**ResourcePermission** | [**ClusterDatastoresPermissionsResourcePermissions**](ClusterDatastoresPermissionsResourcePermissions.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddSecurityGroups200ResponseSecurityGroup = Initialize-PSOpenAPIToolsAddSecurityGroups200ResponseSecurityGroup  -Id null `
 -Name null `
 -Description null `
 -AccountId null `
 -GroupSource null `
 -ExternalId null `
 -Enabled null `
 -SyncSource null `
 -Visibility null `
 -Active null `
 -Zone null `
 -Locations null `
 -Rules null `
 -Tenants null `
 -ResourcePermission null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddSecurityGroups200ResponseSecurityGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

