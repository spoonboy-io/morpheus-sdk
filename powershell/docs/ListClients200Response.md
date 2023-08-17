# ListClients200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**MetaMeta**](MetaMeta.md) |  | [optional] 
**Clients** | [**Client[]**](Client.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ListClients200Response = Initialize-PSOpenAPIToolsListClients200Response  -Meta null `
 -Clients null
```

- Convert the resource to JSON
```powershell
$ListClients200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

