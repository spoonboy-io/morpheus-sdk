# BlueprintCreateSuccess
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Blueprint ID | [optional] 
**Name** | **String** | A name for the blueprint | [optional] 
**Description** | **String** | A description for the blueprint | [optional] 
**Labels** | **String[]** |  | [optional] 
**Category** | **String** | Category | [optional] 
**Config** | [**BlueprintCreateSuccessConfig**](BlueprintCreateSuccessConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$BlueprintCreateSuccess = Initialize-PSOpenAPIToolsBlueprintCreateSuccess  -Id null `
 -Name null `
 -Description Basic Template Description `
 -Labels null `
 -Category web, utility, app `
 -Config null
```

- Convert the resource to JSON
```powershell
$BlueprintCreateSuccess | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

