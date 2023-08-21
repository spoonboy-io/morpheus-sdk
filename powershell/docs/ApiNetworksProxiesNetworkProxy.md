# ApiNetworksProxiesNetworkProxy
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | [optional] 
**ProxyHost** | **String** | Proxy Host | [optional] 
**ProxyPort** | **String** | Proxy Port | [optional] 
**ProxyUser** | **String** | Proxy Username | [optional] 
**ProxyPassword** | **String** | Proxy Password | [optional] 
**ProxyDomain** | **String** | Proxy Domain | [optional] 
**ProxyWorkstation** | **String** | Proxy Workstation | [optional] 
**Visibility** | **String** | Visibility | [optional] [default to "private"]
**Account** | [**ApiNetworksProxiesNetworkProxyAccount**](ApiNetworksProxiesNetworkProxyAccount.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiNetworksProxiesNetworkProxy = Initialize-PSOpenAPIToolsApiNetworksProxiesNetworkProxy  -Name null `
 -ProxyHost null `
 -ProxyPort null `
 -ProxyUser null `
 -ProxyPassword null `
 -ProxyDomain null `
 -ProxyWorkstation null `
 -Visibility null `
 -Account null
```

- Convert the resource to JSON
```powershell
$ApiNetworksProxiesNetworkProxy | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

