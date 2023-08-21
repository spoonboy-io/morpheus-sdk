# BillingComputeServers
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | **Decimal** |  | [optional] 
**Cost** | **Decimal** |  | [optional] 
**Servers** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Count** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BillingComputeServers = Initialize-PSOpenAPIToolsBillingComputeServers  -Price null `
 -Cost null `
 -Servers null `
 -Count null
```

- Convert the resource to JSON
```powershell
$BillingComputeServers | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

