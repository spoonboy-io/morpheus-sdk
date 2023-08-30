# CreateLoadBalancerMonitorRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerMonitor** | [**CreateLoadBalancerMonitorRequestLoadBalancerMonitor**](CreateLoadBalancerMonitorRequestLoadBalancerMonitor.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerMonitorRequest = Initialize-PSOpenAPIToolsCreateLoadBalancerMonitorRequest  -LoadBalancerMonitor null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerMonitorRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

