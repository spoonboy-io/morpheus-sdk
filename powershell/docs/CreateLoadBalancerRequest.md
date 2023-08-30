# CreateLoadBalancerRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | [**LoadBalancerCreate**](LoadBalancerCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerRequest = Initialize-PSOpenAPIToolsCreateLoadBalancerRequest  -LoadBalancer null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

