# CreateLoadBalancerMonitor200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerMonitor** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInner.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerMonitor200Response = Initialize-PSOpenAPIToolsCreateLoadBalancerMonitor200Response  -LoadBalancerMonitor null `
 -Success null `
 -Msg null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerMonitor200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

