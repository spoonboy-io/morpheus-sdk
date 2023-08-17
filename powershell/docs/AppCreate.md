# AppCreate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **Int64** |  | [optional] 
**BlueprintId** | [**AppCreateBlueprintId**](AppCreateBlueprintId.md) |  | 
**Name** | **String** | A unique name for the app | 
**Description** | **String** | Description | [optional] 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Group** | [**AppCreateGroup**](AppCreateGroup.md) |  | [optional] 
**DefaultCloud** | [**AppCreateDefaultCloud**](AppCreateDefaultCloud.md) |  | [optional] 
**Environment** | **String** | Environment code (appContext) | [optional] 
**Tiers** | [**SystemCollectionsHashtable**](.md) | Configuration of app elements | [optional] 

## Examples

- Prepare the resource
```powershell
$AppCreate = Initialize-PSOpenAPIToolsAppCreate  -TemplateId null `
 -BlueprintId null `
 -Name null `
 -Description null `
 -Labels null `
 -Group null `
 -DefaultCloud null `
 -Environment null `
 -Tiers null
```

- Convert the resource to JSON
```powershell
$AppCreate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

