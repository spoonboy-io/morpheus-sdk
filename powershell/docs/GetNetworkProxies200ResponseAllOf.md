# GetNetworkProxies200ResponseAllOf
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkProxies** | [**GetNetworkProxies200ResponseAllOfNetworkProxiesInner[]**](GetNetworkProxies200ResponseAllOfNetworkProxiesInner.md) |  | [optional] 
**NetworkProxyCount** | **Int32** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkProxies200ResponseAllOf = Initialize-PSOpenAPIToolsGetNetworkProxies200ResponseAllOf  -NetworkProxies null `
 -NetworkProxyCount null
```

- Convert the resource to JSON
```powershell
$GetNetworkProxies200ResponseAllOf | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

