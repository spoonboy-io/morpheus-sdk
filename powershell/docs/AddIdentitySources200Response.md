# AddIdentitySources200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | [**AddIdentitySources200ResponseAllOfTask**](AddIdentitySources200ResponseAllOfTask.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddIdentitySources200Response = Initialize-PSOpenAPIToolsAddIdentitySources200Response  -Task null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddIdentitySources200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

