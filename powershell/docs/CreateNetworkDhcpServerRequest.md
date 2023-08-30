# CreateNetworkDhcpServerRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDhcpServer** | [**NetworkDhcpServerCreate**](NetworkDhcpServerCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworkDhcpServerRequest = Initialize-PSOpenAPIToolsCreateNetworkDhcpServerRequest  -NetworkDhcpServer null
```

- Convert the resource to JSON
```powershell
$CreateNetworkDhcpServerRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

