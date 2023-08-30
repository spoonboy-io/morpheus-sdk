# GetNetworkEdgeCluster200ResponseNetworkEdgeCluster
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int32** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**ProviderId** | **String** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**DisplayName** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Config** | [**GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerConfig**](GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerConfig.md) |  | [optional] 
**Owner** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) |  | [optional] 
**NetworkServer** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) |  | [optional] 
**Zone** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInnerOwner.md) |  | [optional] 
**Tenants** | [**GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner[]**](GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInnerTenantsInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkEdgeCluster200ResponseNetworkEdgeCluster = Initialize-PSOpenAPIToolsGetNetworkEdgeCluster200ResponseNetworkEdgeCluster  -Id null `
 -InternalId null `
 -Visibility null `
 -DateCreated null `
 -ProviderId null `
 -LastUpdated null `
 -Active null `
 -DisplayName null `
 -Name null `
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
$GetNetworkEdgeCluster200ResponseNetworkEdgeCluster | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

