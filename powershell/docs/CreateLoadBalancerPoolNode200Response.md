# CreateLoadBalancerPoolNode200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerNode** | [**ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner**](ListLoadBalancerPoolNodes200ResponseAllOfLoadBalancerNodesInner.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerPoolNode200Response = Initialize-PSOpenAPIToolsCreateLoadBalancerPoolNode200Response  -LoadBalancerNode null `
 -Success null `
 -Msg null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerPoolNode200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

