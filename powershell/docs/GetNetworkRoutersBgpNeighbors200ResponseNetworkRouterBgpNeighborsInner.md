# GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**IpAddress** | **String** |  | [optional] 
**ForwardingAddress** | **String** |  | [optional] 
**ProtocolAddress** | **String** |  | [optional] 
**RemoteAs** | **String** |  | [optional] 
**Weight** | **Int64** |  | [optional] 
**KeepAlive** | **Int64** |  | [optional] 
**HoldDown** | **Int64** |  | [optional] 
**Password** | **String** |  | [optional] 
**RouteFilteringType** | **String** |  | [optional] 
**RouteFilteringIn** | **String** |  | [optional] 
**RouteFilteringOut** | **String** |  | [optional] 
**BfdEnabled** | **Boolean** |  | [optional] 
**BfdInterval** | **Int64** |  | [optional] 
**BfdMultiple** | **Int64** |  | [optional] 
**AllowAsIn** | **Boolean** |  | [optional] 
**HopLimit** | **Int64** |  | [optional] 
**RestartMode** | **String** |  | [optional] 
**ProviderId** | **String** |  | [optional] 
**SyncSource** | **String** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **String** |  | [optional] 
**Config** | [**GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig**](GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInnerConfig.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner = Initialize-PSOpenAPIToolsGetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner  -Id null `
 -IpAddress null `
 -ForwardingAddress null `
 -ProtocolAddress null `
 -RemoteAs null `
 -Weight null `
 -KeepAlive null `
 -HoldDown null `
 -Password null `
 -RouteFilteringType null `
 -RouteFilteringIn null `
 -RouteFilteringOut null `
 -BfdEnabled null `
 -BfdInterval null `
 -BfdMultiple null `
 -AllowAsIn null `
 -HopLimit null `
 -RestartMode null `
 -ProviderId null `
 -SyncSource null `
 -InternalId null `
 -ExternalId null `
 -RefType null `
 -RefId null `
 -Config null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$GetNetworkRoutersBgpNeighbors200ResponseNetworkRouterBgpNeighborsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

