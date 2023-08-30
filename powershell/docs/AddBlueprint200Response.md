# AddBlueprint200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | [**BlueprintCreateSuccess**](BlueprintCreateSuccess.md) |  | [optional] 
**Msg** | **String** |  | [optional] 
**Errors** | **String** |  | [optional] 
**ErrorCode** | **String** |  | [optional] 
**InProgress** | **Boolean** |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddBlueprint200Response = Initialize-PSOpenAPIToolsAddBlueprint200Response  -Blueprint null `
 -Msg null `
 -Errors null `
 -ErrorCode null `
 -InProgress null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddBlueprint200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

