# ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | [optional] 
**Description** | **String** | Description | [optional] 
**MonitorType** | **String** |  | [optional] 
**MonitorTimeout** | **Int64** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by type. | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor = Initialize-PSOpenAPIToolsApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor  -Name null `
 -Description null `
 -MonitorType null `
 -MonitorTimeout null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ApiLoadBalancersLoadBalancerIdMonitorsLoadBalancerMonitor | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

