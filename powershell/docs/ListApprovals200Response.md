# ListApprovals200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Approvals** | [**Approvals[]**](Approvals.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListApprovals200Response = Initialize-PSOpenAPIToolsListApprovals200Response  -Meta null `
 -Approvals null
```

- Convert the resource to JSON
```powershell
$ListApprovals200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

