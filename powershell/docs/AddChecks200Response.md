# AddChecks200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Check** | [**Check**](Check.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddChecks200Response = Initialize-PSOpenAPIToolsAddChecks200Response  -Check null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddChecks200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

