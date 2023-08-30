# AddServicePlans200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlan** | [**ServicePlan**](ServicePlan.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddServicePlans200Response = Initialize-PSOpenAPIToolsAddServicePlans200Response  -ServicePlan null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddServicePlans200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

