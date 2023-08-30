# CreateLoadBalancerProfile200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancerProfile** | [**ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner**](ListLoadBalancerProfiles200ResponseAllOfLoadBalancerProfilesInner.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerProfile200Response = Initialize-PSOpenAPIToolsCreateLoadBalancerProfile200Response  -LoadBalancerProfile null `
 -Success null `
 -Msg null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerProfile200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

