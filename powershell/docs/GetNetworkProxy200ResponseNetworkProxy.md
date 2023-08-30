# GetNetworkProxy200ResponseNetworkProxy
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**ProxyHost** | **String** |  | [optional] 
**ProxyPort** | **Int64** |  | [optional] 
**ProxyUser** | **String** |  | [optional] 
**ProxyPassword** | **String** |  | [optional] 
**ProxyDomain** | **String** |  | [optional] 
**ProxyWorkstation** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Account** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Owner** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkProxy200ResponseNetworkProxy = Initialize-PSOpenAPIToolsGetNetworkProxy200ResponseNetworkProxy  -Id null `
 -Name null `
 -ProxyHost null `
 -ProxyPort null `
 -ProxyUser null `
 -ProxyPassword null `
 -ProxyDomain null `
 -ProxyWorkstation null `
 -Visibility null `
 -Account null `
 -Owner null
```

- Convert the resource to JSON
```powershell
$GetNetworkProxy200ResponseNetworkProxy | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

