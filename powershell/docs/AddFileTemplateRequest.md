# AddFileTemplateRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerTemplate** | [**FileTemplateCreate**](FileTemplateCreate.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddFileTemplateRequest = Initialize-PSOpenAPIToolsAddFileTemplateRequest  -ContainerTemplate null
```

- Convert the resource to JSON
```powershell
$AddFileTemplateRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

