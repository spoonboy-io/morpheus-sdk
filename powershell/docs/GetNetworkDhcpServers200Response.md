# GetNetworkDhcpServers200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDhcpServers** | [**GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner[]**](GetNetworkDhcpServers200ResponseAllOfNetworkDhcpServersInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkDhcpServers200Response = Initialize-PSOpenAPIToolsGetNetworkDhcpServers200Response  -NetworkDhcpServers null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkDhcpServers200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

