# ApiEnvironmentsIdEnvironment
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name for the environment | [optional] 
**Description** | **String** | A description of the environment | [optional] 
**Visibility** | **String** | private or public | [optional] [default to "private"]
**SortOrder** | **Int64** | Sort order | [optional] [default to 0]
**Active** | **Boolean** | Set to false to deactivate the environment | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiEnvironmentsIdEnvironment = Initialize-PSOpenAPIToolsApiEnvironmentsIdEnvironment  -Name Production `
 -Description Tag used for Production environments `
 -Visibility null `
 -SortOrder null `
 -Active true
```

- Convert the resource to JSON
```powershell
$ApiEnvironmentsIdEnvironment | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

