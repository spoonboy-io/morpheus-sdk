# MetaMeta
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | **Int64** | Offset records, the number of records to skip, can be used to paginate over results. | [optional] [default to 0]
**Max** | **Int64** | Max size, the maximum number of records to include in the response. | [optional] [default to 25]
**Size** | **Int64** | Number of records returned in the response | [optional] [default to 0]
**Total** | **Int64** | Total number of records found | [optional] [default to 0]

## Examples

- Prepare the resource
```powershell
$MetaMeta = Initialize-PSOpenAPIToolsMetaMeta  -Offset null `
 -Max null `
 -Size null `
 -Total null
```

- Convert the resource to JSON
```powershell
$MetaMeta | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

