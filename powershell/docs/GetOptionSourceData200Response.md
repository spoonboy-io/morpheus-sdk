# GetOptionSourceData200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VarData** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetOptionSourceData200Response = Initialize-PSOpenAPIToolsGetOptionSourceData200Response  -VarData null `
 -Success null
```

- Convert the resource to JSON
```powershell
$GetOptionSourceData200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

