# AddTenant200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**Tenant**](Tenant.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddTenant200Response = Initialize-PSOpenAPIToolsAddTenant200Response  -Account null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddTenant200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

