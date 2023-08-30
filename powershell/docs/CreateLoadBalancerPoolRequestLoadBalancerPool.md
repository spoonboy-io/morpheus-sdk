# CreateLoadBalancerPoolRequestLoadBalancerPool
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | [optional] 
**Description** | **String** | Description | [optional] 
**VipBalance** | **String** | Balance Algorithm | [optional] 
**MinActive** | **Int64** | Min Active Members | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by type. | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerPoolRequestLoadBalancerPool = Initialize-PSOpenAPIToolsCreateLoadBalancerPoolRequestLoadBalancerPool  -Name null `
 -Description null `
 -VipBalance null `
 -MinActive null `
 -Config null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerPoolRequestLoadBalancerPool | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

