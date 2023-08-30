# FileTemplateCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | File template name | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**FileName** | **String** | Filename for the file template | 
**FilePath** | **String** | Path for the file template | [optional] 
**Category** | **String** | Category | [optional] 
**TemplatePhase** | **String** | Template Phase, provision, start, etc. | [optional] 
**Template** | **String** | Template content, that is, the file template content itself. | [optional] 
**FileOwner** | **Int64** | File Owner | [optional] 
**SettingName** | **String** | Setting Name | [optional] 
**SettingCategory** | **String** | Setting Category | [optional] 

## Examples

- Prepare the resource
```powershell
$FileTemplateCreate = Initialize-PSOpenAPIToolsFileTemplateCreate  -Name null `
 -Labels null `
 -FileName null `
 -FilePath null `
 -Category null `
 -TemplatePhase null `
 -Template null `
 -FileOwner null `
 -SettingName null `
 -SettingCategory null
```

- Convert the resource to JSON
```powershell
$FileTemplateCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

