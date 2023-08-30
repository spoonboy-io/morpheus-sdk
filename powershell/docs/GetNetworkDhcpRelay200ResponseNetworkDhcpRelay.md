# GetNetworkDhcpRelay200ResponseNetworkDhcpRelay
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**ProviderId** | **String** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Name** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ServerIpAddresses** | **String[]** |  | [optional] 
**Owner** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**NetworkServer** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkDhcpRelay200ResponseNetworkDhcpRelay = Initialize-PSOpenAPIToolsGetNetworkDhcpRelay200ResponseNetworkDhcpRelay  -Id null `
 -DateCreated null `
 -ProviderId null `
 -LastUpdated null `
 -Name null `
 -ExternalId null `
 -ServerIpAddresses null `
 -Owner null `
 -NetworkServer null
```

- Convert the resource to JSON
```powershell
$GetNetworkDhcpRelay200ResponseNetworkDhcpRelay | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

