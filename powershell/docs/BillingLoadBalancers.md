# BillingLoadBalancers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**LoadBalancers** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Count** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingLoadBalancers = Initialize-PSOpenAPIToolsBillingLoadBalancers  -Price null `
 -Cost null `
 -LoadBalancers null `
 -Count null
```

- Convert the resource to JSON
```powershell
$BillingLoadBalancers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

