# GetNetworkEdgeClusters200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkEdgeClusters** | [**GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInner[]**](GetNetworkEdgeClusters200ResponseAllOfNetworkEdgeClustersInner.md) |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkEdgeClusters200Response = Initialize-PSOpenAPIToolsGetNetworkEdgeClusters200Response  -NetworkEdgeClusters null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkEdgeClusters200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

