# ApiUsersIdUser
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **String** | First Name | [optional] 
**LastName** | **String** | Last Name | [optional] 
**Username** | **String** | Username (unique per tenant). | [optional] 
**Email** | **String** | Email address | [optional] 
**Password** | **String** | Password | [optional] 
**Roles** | [**ReferenceObject[]**](ReferenceObject.md) | List of Roles | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiUsersIdUser = Initialize-PSOpenAPIToolsApiUsersIdUser  -FirstName null `
 -LastName null `
 -Username null `
 -Email null `
 -Password null `
 -Roles null
```

- Convert the resource to JSON
```powershell
$ApiUsersIdUser | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

