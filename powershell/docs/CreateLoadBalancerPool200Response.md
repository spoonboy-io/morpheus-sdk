# CreateLoadBalancerPool200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerPool** | [**ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner**](ListLoadBalancerPools200ResponseAllOfLoadBalancerPoolsInner.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerPool200Response = Initialize-PSOpenAPIToolsCreateLoadBalancerPool200Response  -LoadBalancerPool null `
 -Success null `
 -Msg null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerPool200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

