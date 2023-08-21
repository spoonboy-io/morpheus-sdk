# BillingServers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Servers** | [**BillingServersServers[]**](BillingServersServers.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingServers = Initialize-PSOpenAPIToolsBillingServers  -Price null `
 -Cost null `
 -StartDate null `
 -EndDate null `
 -Servers null
```

- Convert the resource to JSON
```powershell
$BillingServers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

