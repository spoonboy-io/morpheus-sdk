# AppPrepareApplyData
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**AutoValidate** | **Boolean** |  | [optional] 
**Terraform** | [**AppPrepareApplyDataTerraform**](AppPrepareApplyDataTerraform.md) |  | [optional] 
**Type** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**BlueprintName** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**TemplateId** | **Int64** |  | [optional] 
**BlueprintId** | **Int64** |  | [optional] 
**Group** | [**AppPrepareApplyDataGroup**](AppPrepareApplyDataGroup.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppPrepareApplyData = Initialize-PSOpenAPIToolsAppPrepareApplyData  -Image null `
 -Name null `
 -AutoValidate null `
 -Terraform null `
 -Type null `
 -Config null `
 -BlueprintName null `
 -Description null `
 -TemplateId null `
 -BlueprintId null `
 -Group null
```

- Convert the resource to JSON
```powershell
$AppPrepareApplyData | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

