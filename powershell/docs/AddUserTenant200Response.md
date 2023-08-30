# AddUserTenant200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**User**](User.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddUserTenant200Response = Initialize-PSOpenAPIToolsAddUserTenant200Response  -User null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddUserTenant200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

