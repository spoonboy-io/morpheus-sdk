# AddCheckApps200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**CheckApp** | [**CheckApp**](CheckApp.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCheckApps200Response = Initialize-PSOpenAPIToolsAddCheckApps200Response  -Success null `
 -CheckApp null
```

- Convert the resource to JSON
```powershell
$AddCheckApps200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

