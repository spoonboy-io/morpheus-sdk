# GetCheckGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckGroup** | [**CheckGroup**](CheckGroup.md) |  | [optional] 
**Checks** | [**Check[]**](Check.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetCheckGroups200Response = Initialize-PSOpenAPIToolsGetCheckGroups200Response  -CheckGroup null `
 -Checks null
```

- Convert the resource to JSON
```powershell
$GetCheckGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

