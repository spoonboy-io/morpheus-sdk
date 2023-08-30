# ClustersServersInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**TypeSet** | [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancerType.md) |  | [optional] 
**ComputeServerType** | [**ClustersServersInnerComputeServerType**](ClustersServersInnerComputeServerType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClustersServersInner = Initialize-PSOpenAPIToolsClustersServersInner  -Id null `
 -Name null `
 -TypeSet null `
 -ComputeServerType null
```

- Convert the resource to JSON
```powershell
$ClustersServersInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

