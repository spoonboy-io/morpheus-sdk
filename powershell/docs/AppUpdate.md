# AppUpdate
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A unique name for the app | [optional] 
**Description** | **String** | Description | [optional] 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Environment** | **String** | Environment code (appContext) | [optional] 
**OwnerId** | **Int64** | User ID, can be used to change app owner. This also changes the owner for each instance in the app. | [optional] 

## Examples

- Prepare the resource
```powershell
$AppUpdate = Initialize-PSOpenAPIToolsAppUpdate  -Name null `
 -Description null `
 -Labels null `
 -Environment null `
 -OwnerId null
```

- Convert the resource to JSON
```powershell
$AppUpdate | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

