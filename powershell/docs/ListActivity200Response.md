# ListActivity200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Activity** | [**Activity[]**](Activity.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListActivity200Response = Initialize-PSOpenAPIToolsListActivity200Response  -Meta null `
 -Activity null
```

- Convert the resource to JSON
```powershell
$ListActivity200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

