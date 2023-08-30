# GetNetworkRouter200ResponseNetworkRouter
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
**Type** | [**GetNetworkRouter200ResponseNetworkRouterType**](GetNetworkRouter200ResponseNetworkRouterType.md) |  | [optional] 
**NetworkServer** | [**GetNetworkRouter200ResponseNetworkRouterNetworkServer**](GetNetworkRouter200ResponseNetworkRouterNetworkServer.md) |  | [optional] 
**Zone** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Instance** | **String** |  | [optional] 
**ExternalNetwork** | **String** |  | [optional] 
**Site** | **String** |  | [optional] 
**Interfaces** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Firewall** | [**GetNetworkRouter200ResponseNetworkRouterFirewall**](GetNetworkRouter200ResponseNetworkRouterFirewall.md) |  | [optional] 
**Routes** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Nats** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Permissions** | [**GetNetworkRouter200ResponseNetworkRouterPermissions**](GetNetworkRouter200ResponseNetworkRouterPermissions.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouter200ResponseNetworkRouter = Initialize-PSOpenAPIToolsGetNetworkRouter200ResponseNetworkRouter  -Id null `
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
 -Interfaces null `
 -Firewall null `
 -Routes null `
 -Nats null `
 -Permissions null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouter200ResponseNetworkRouter | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

