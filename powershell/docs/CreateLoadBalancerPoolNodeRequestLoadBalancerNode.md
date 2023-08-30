# CreateLoadBalancerPoolNodeRequestLoadBalancerNode
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | [optional] 
**Description** | **String** | Description | [optional] 
**IpAddress** | **String** | IP Address | [optional] 
**Port** | **Int32** | Port number | [optional] 
**Weight** | **Int32** | Weight applied balance algoritm | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by type. | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateLoadBalancerPoolNodeRequestLoadBalancerNode = Initialize-PSOpenAPIToolsCreateLoadBalancerPoolNodeRequestLoadBalancerNode  -Name null `
 -Description null `
 -IpAddress null `
 -Port null `
 -Weight null `
 -Config null
```

- Convert the resource to JSON
```powershell
$CreateLoadBalancerPoolNodeRequestLoadBalancerNode | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

