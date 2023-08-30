# GetNetworkTransportZones200ResponseAllOfNetworkScopesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**ProviderId** | **String** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**StreamType** | **String** |  | [optional] 
**DisplayName** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Config** | [**GetNetworkTransportZones200ResponseAllOfNetworkScopesInnerConfig**](GetNetworkTransportZones200ResponseAllOfNetworkScopesInnerConfig.md) |  | [optional] 
**Owner** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**NetworkServer** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Zone** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Tenants** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkTransportZones200ResponseAllOfNetworkScopesInner = Initialize-PSOpenAPIToolsGetNetworkTransportZones200ResponseAllOfNetworkScopesInner  -Id null `
 -InternalId null `
 -Visibility null `
 -DateCreated null `
 -ProviderId null `
 -LastUpdated null `
 -Active null `
 -StreamType null `
 -DisplayName null `
 -Name null `
 -Status null `
 -Enabled null `
 -ExternalId null `
 -Config null `
 -Owner null `
 -NetworkServer null `
 -Zone null `
 -Tenants null
```

- Convert the resource to JSON
```powershell
$GetNetworkTransportZones200ResponseAllOfNetworkScopesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

