# GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**ProviderId** | **String** |  | [optional] 
**ServerIpAddress** | **String** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**LeaseTime** | **Int32** |  | [optional] 
**Name** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Config** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerConfig.md) |  | [optional] 
**Owner** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) |  | [optional] 
**NetworkServer** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner = Initialize-PSOpenAPIToolsGetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner  -Id null `
 -DateCreated null `
 -ProviderId null `
 -ServerIpAddress null `
 -LastUpdated null `
 -LeaseTime null `
 -Name null `
 -ExternalId null `
 -Config null `
 -Owner null `
 -NetworkServer null
```

- Convert the resource to JSON
```powershell
$GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

