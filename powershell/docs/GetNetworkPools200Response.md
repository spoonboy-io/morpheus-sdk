# GetNetworkPools200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPools** | [**GetNetworkPools200ResponseAllOfNetworkPoolsInner[]**](GetNetworkPools200ResponseAllOfNetworkPoolsInner.md) |  | [optional] 
**NetworkPoolCount** | **Int64** |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkPools200Response = Initialize-PSOpenAPIToolsGetNetworkPools200Response  -NetworkPools null `
 -NetworkPoolCount null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkPools200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

