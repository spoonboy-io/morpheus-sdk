# CreateLoadBalancerPoolRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerPool** | [**CreateLoadBalancerPoolRequestLoadBalancerPool**](CreateLoadBalancerPoolRequestLoadBalancerPool.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerPoolRequest = Initialize-PSOpenAPIToolsCreateLoadBalancerPoolRequest  -LoadBalancerPool null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerPoolRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

