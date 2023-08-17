# PrepareAppApply200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VarData** | [**AppPrepareApplyData**](AppPrepareApplyData.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$PrepareAppApply200Response = Initialize-PSOpenAPIToolsPrepareAppApply200Response  -VarData null `
 -Success null
```

- Convert the resource to JSON
```powershell
$PrepareAppApply200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

