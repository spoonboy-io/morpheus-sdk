# CreateLoadBalancerPoolNodeRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerNode** | [**CreateLoadBalancerPoolNodeRequestLoadBalancerNode**](CreateLoadBalancerPoolNodeRequestLoadBalancerNode.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerPoolNodeRequest = Initialize-PSOpenAPIToolsCreateLoadBalancerPoolNodeRequest  -LoadBalancerNode null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerPoolNodeRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

