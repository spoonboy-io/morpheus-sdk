# ApplyAppStateRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateParameter** | [**SystemCollectionsHashtable**](.md) | Template Parameter object. A map of key-value pairs that correspond to the template variables i.e. tfvars | [optional] 

## Examples

- Prepare the resource
```powershell
$ApplyAppStateRequest = Initialize-PSOpenAPIToolsApplyAppStateRequest  -TemplateParameter null
```

- Convert the resource to JSON
```powershell
$ApplyAppStateRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

