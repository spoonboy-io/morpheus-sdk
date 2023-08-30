# CreateLoadBalancerVirtualServerRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerInstance** | [**CreateLoadBalancerVirtualServerRequestLoadBalancerInstance**](CreateLoadBalancerVirtualServerRequestLoadBalancerInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerVirtualServerRequest = Initialize-PSOpenAPIToolsCreateLoadBalancerVirtualServerRequest  -LoadBalancerInstance null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerVirtualServerRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

