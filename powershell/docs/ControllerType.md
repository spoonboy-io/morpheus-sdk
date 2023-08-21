# ControllerType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**DisplayOrder** | **Int64** |  | [optional] 
**Category** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Creatable** | **Boolean** |  | [optional] 
**MaxDevices** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ControllerType = Initialize-PSOpenAPIToolsControllerType  -Id null `
 -Name null `
 -DisplayOrder null `
 -Category null `
 -Enabled null `
 -Creatable null `
 -MaxDevices null
```

- Convert the resource to JSON
```powershell
$ControllerType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

