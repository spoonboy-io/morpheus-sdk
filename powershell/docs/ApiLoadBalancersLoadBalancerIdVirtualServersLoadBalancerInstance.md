# ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VipName** | **String** | VIP Name | [optional] 
**Description** | **String** | Description | [optional] 
**VipAddress** | **String** | VIP Address | [optional] 
**VipPort** | **String** | VIP Port | [optional] 
**VipProtocol** | **String** | VIP Protocol | [optional] 
**VipHostname** | **String** | VIP Hostname | [optional] 
**SslCert** | **Int64** | SSL Client Certificate ID | [optional] 
**SslServerCert** | **Int64** | SSL Server Certificate ID | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) | Configuration object with parameters that vary by type. | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance = Initialize-PSOpenAPIToolsApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance  -VipName null `
 -Description null `
 -VipAddress null `
 -VipPort null `
 -VipProtocol null `
 -VipHostname null `
 -SslCert null `
 -SslServerCert null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ApiLoadBalancersLoadBalancerIdVirtualServersLoadBalancerInstance | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

