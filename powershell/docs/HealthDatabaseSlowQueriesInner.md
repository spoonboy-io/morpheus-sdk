# HealthDatabaseSlowQueriesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **Int64** |  | [optional] 
**Time** | **Int64** |  | [optional] 
**Query** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthDatabaseSlowQueriesInner = Initialize-PSOpenAPIToolsHealthDatabaseSlowQueriesInner  -Count null `
 -Time null `
 -Query null
```

- Convert the resource to JSON
```powershell
$HealthDatabaseSlowQueriesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

