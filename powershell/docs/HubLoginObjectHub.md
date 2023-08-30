# HubLoginObjectHub
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **String** | Email for existing Morpheus Hub user | 
**Password** | **String** | Password for existing Morpheus Hub user | 

## Examples

- Prepare the resource
```powershell
$HubLoginObjectHub = Initialize-PSOpenAPIToolsHubLoginObjectHub  -Email null `
 -Password null
```

- Convert the resource to JSON
```powershell
$HubLoginObjectHub | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

