# GetNetworkDomains200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDomains** | [**NetworkDomain[]**](NetworkDomain.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkDomains200Response = Initialize-PSOpenAPIToolsGetNetworkDomains200Response  -NetworkDomains null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkDomains200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

