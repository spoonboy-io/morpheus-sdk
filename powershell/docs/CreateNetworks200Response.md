# CreateNetworks200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | [**Network**](Network.md) |  | [optional] 
**Errors** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**Msg** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateNetworks200Response = Initialize-PSOpenAPIToolsCreateNetworks200Response  -Network null `
 -Errors null `
 -Success null `
 -Msg null
```

- Convert the resource to JSON
```powershell
$CreateNetworks200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

