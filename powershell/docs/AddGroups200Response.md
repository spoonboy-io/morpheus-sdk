# AddGroups200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**Group**](Group.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddGroups200Response = Initialize-PSOpenAPIToolsAddGroups200Response  -Group null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddGroups200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

