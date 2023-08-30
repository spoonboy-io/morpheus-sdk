# CreateLoadBalancerProfileRequestLoadBalancerProfile
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | [optional] 
**Description** | **String** | Description | [optional] 
**ServiceType** | **String** | Service Type | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by type. | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerProfileRequestLoadBalancerProfile = Initialize-PSOpenAPIToolsCreateLoadBalancerProfileRequestLoadBalancerProfile  -Name null `
 -Description null `
 -ServiceType null `
 -Config null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerProfileRequestLoadBalancerProfile | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

