# AddEnvironmentsRequestEnvironment
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name for the environment | 
**Code** | **String** | A unique code for the environment | 
**Description** | **String** | A description of the environment | [optional] 
**Visibility** | **String** | private or public | [optional] [default to "private"]
**SortOrder** | **Int64** | Sort order | [optional] [default to 0]

## Examples

- Prepare the resource
```powershell
$AddEnvironmentsRequestEnvironment = Initialize-PSOpenAPIToolsAddEnvironmentsRequestEnvironment  -Name Production `
 -Code prod `
 -Description Tag used for Production environments `
 -Visibility null `
 -SortOrder null
```

- Convert the resource to JSON
```powershell
$AddEnvironmentsRequestEnvironment | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

