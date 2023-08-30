# GetNetworkRouters200ResponseNetworkRoutersInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**RouterType** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**ExternalIp** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**ProviderId** | **String** |  | [optional] 
**Type** | [**GetNetworkRouters200ResponseNetworkRoutersInnerType**](GetNetworkRouters200ResponseNetworkRoutersInnerType.md) |  | [optional] 
**NetworkServer** | **String** |  | [optional] 
**Zone** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Instance** | **String** |  | [optional] 
**ExternalNetwork** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Site** | **String** |  | [optional] 
**Interfaces** | [**GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner[]**](GetNetworkRouters200ResponseNetworkRoutersInnerInterfacesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouters200ResponseNetworkRoutersInner = Initialize-PSOpenAPIToolsGetNetworkRouters200ResponseNetworkRoutersInner  -Id null `
 -Code null `
 -Name null `
 -Description null `
 -Category null `
 -DateCreated null `
 -LastUpdated null `
 -RouterType null `
 -Status null `
 -Enabled null `
 -ExternalIp null `
 -ExternalId null `
 -ProviderId null `
 -Type null `
 -NetworkServer null `
 -Zone null `
 -Instance null `
 -ExternalNetwork null `
 -Site null `
 -Interfaces null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouters200ResponseNetworkRoutersInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

