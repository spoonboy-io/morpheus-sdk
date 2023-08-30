# AddCheckGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckGroup** | [**CheckGroup**](CheckGroup.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddCheckGroups200Response = Initialize-PSOpenAPIToolsAddCheckGroups200Response  -CheckGroup null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddCheckGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

