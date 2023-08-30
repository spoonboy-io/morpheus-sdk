# AddTenantRequestAccount
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Name | 
**Description** | **String** | Description | [optional] 
**Role** | [**AddTenantRequestAccountRole**](AddTenantRequestAccountRole.md) |  | [optional] 
**Subdomain** | **String** | The subdomain. This will be part of the login URL and username for sub tenant users. | [optional] 
**Currency** | [**CurrencyCode**](CurrencyCode.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddTenantRequestAccount = Initialize-PSOpenAPIToolsAddTenantRequestAccount  -Name Test Tenant `
 -Description Testing Tenant Creation `
 -Role null `
 -Subdomain tt `
 -Currency null
```

- Convert the resource to JSON
```powershell
$AddTenantRequestAccount | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

