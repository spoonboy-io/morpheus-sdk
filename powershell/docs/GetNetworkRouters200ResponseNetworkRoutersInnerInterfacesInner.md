# GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**InterfaceType** | **String** |  | [optional] 
**NetworkPosition** | **String** |  | [optional] 
**IpAddress** | **String** |  | [optional] 
**Cidr** | **String** |  | [optional] 
**ExternalLink** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Network** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInnerNetwork.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner = Initialize-PSOpenAPIToolsGetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner  -Id null `
 -Name null `
 -Code null `
 -InterfaceType null `
 -NetworkPosition null `
 -IpAddress null `
 -Cidr null `
 -ExternalLink null `
 -Enabled null `
 -Network null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

