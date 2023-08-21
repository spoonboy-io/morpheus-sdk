# HubRegisterObjectHub
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompanyName** | **String** | Company Name for new Morpheus Hub organization | 
**FirstName** | **String** | First Name for new Morpheus Hub user | 
**LastName** | **String** | Last Name for new Morpheus Hub user | 
**Email** | **String** | Email for new Morpheus Hub user | 
**Password** | **String** | Password for new Morpheus Hub user | 
**JobTitle** | **String** | Job title of new Morpheus Hub user | 

## Examples

- Prepare the resource
```powershell
$HubRegisterObjectHub = Initialize-PSOpenAPIToolsHubRegisterObjectHub  -CompanyName null `
 -FirstName null `
 -LastName null `
 -Email null `
 -Password null `
 -JobTitle null
```

- Convert the resource to JSON
```powershell
$HubRegisterObjectHub | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

