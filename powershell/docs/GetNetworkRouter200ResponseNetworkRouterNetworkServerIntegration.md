# GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Type** | **String** |  | [optional] 
**IntegrationType** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**Url** | **String** |  | [optional] 
**Port** | **String** |  | [optional] 
**Username** | **String** |  | [optional] 
**Password** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **String** |  | [optional] 
**IsPlugin** | **Boolean** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Status** | **String** |  | [optional] 
**StatusDate** | **System.DateTime** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**LastSync** | **System.DateTime** |  | [optional] 
**LastSyncDuration** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration = Initialize-PSOpenAPIToolsGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration  -Id null `
 -Name null `
 -Enabled null `
 -Type null `
 -IntegrationType null `
 -Url null `
 -Port null `
 -Username null `
 -Password null `
 -RefType null `
 -RefId null `
 -IsPlugin null `
 -Config null `
 -Status null `
 -StatusDate null `
 -StatusMessage null `
 -LastSync null `
 -LastSyncDuration null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

