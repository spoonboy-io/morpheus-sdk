# AddJobs200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | [**Job**](Job.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddJobs200Response = Initialize-PSOpenAPIToolsAddJobs200Response  -Job null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddJobs200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

