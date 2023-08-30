# AddIncident200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incident** | [**Incident**](Incident.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddIncident200Response = Initialize-PSOpenAPIToolsAddIncident200Response  -Incident null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddIncident200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

