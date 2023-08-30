# GuidanceTypesInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Title** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceTypesInner = Initialize-PSOpenAPIToolsGuidanceTypesInner  -Id null `
 -Name null `
 -Code null `
 -Category null `
 -Title null
```

- Convert the resource to JSON
```powershell
$GuidanceTypesInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

