# GetNetworkDhcpRelays200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDhcpRelays** | [**GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner[]**](GetNetworkDhcpRelays200ResponseAllOfNetworkDhcpRelaysInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkDhcpRelays200Response = Initialize-PSOpenAPIToolsGetNetworkDhcpRelays200Response  -NetworkDhcpRelays null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkDhcpRelays200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

