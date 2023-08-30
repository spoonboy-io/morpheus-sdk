# GetNetworkProxies200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkProxies** | [**GetNetworkProxies200ResponseAllOfNetworkProxiesInner[]**](GetNetworkProxies200ResponseAllOfNetworkProxiesInner.md) |  | [optional] 
**NetworkProxyCount** | **Int32** |  | [optional] 
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkProxies200Response = Initialize-PSOpenAPIToolsGetNetworkProxies200Response  -NetworkProxies null `
 -NetworkProxyCount null `
 -Meta null
```

- Convert the resource to JSON
```powershell
$GetNetworkProxies200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

