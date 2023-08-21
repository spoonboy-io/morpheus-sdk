# AppStateWorkloads
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefType** | **String** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**RefName** | **String** |  | [optional] 
**SubRefName** | **String** |  | [optional] 
**StateDate** | **System.DateTime** |  | [optional] 
**Status** | **String** |  | [optional] 
**IacDrift** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AppStateWorkloads = Initialize-PSOpenAPIToolsAppStateWorkloads  -RefType null `
 -RefId null `
 -RefName null `
 -SubRefName null `
 -StateDate null `
 -Status null `
 -IacDrift null
```

- Convert the resource to JSON
```powershell
$AppStateWorkloads | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

