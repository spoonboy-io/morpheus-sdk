# GetNetworkPoolIps200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPoolIps** | [**GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner[]**](GetNetworkPoolIps200ResponseAllOfNetworkPoolIpsInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkPoolIps200Response = Initialize-PSOpenAPIToolsGetNetworkPoolIps200Response  -NetworkPoolIps null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkPoolIps200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

