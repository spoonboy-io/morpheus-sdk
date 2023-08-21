# ClustersServers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**TypeSet** | [**InlineResponse20079LoadBalancerMonitorLoadBalancerType**](InlineResponse20079LoadBalancerMonitorLoadBalancerType.md) |  | [optional] 
**ComputeServerType** | [**ClustersComputeServerType**](ClustersComputeServerType.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClustersServers = Initialize-PSOpenAPIToolsClustersServers  -Id null `
 -Name null `
 -TypeSet null `
 -ComputeServerType null
```

- Convert the resource to JSON
```powershell
$ClustersServers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

