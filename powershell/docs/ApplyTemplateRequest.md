# ApplyTemplateRequest
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceUrl** | **String** | Url of desired template to apply to cluster | [optional] 
**SpecTemplate** | **String** | Name or ID of desired Spec Template to apply to cluster | [optional] 
**SpecYaml** | **String** | Yaml of template to apply to cluster | [optional] 

## Examples

- Prepare the resource
```powershell
$ApplyTemplateRequest = Initialize-PSOpenAPIToolsApplyTemplateRequest  -ServiceUrl null `
 -SpecTemplate null `
 -SpecYaml null
```

- Convert the resource to JSON
```powershell
$ApplyTemplateRequest | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

