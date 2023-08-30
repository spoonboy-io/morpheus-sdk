# AddStorageBuckets200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageBucket** | [**StorageBucket**](StorageBucket.md) |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddStorageBuckets200Response = Initialize-PSOpenAPIToolsAddStorageBuckets200Response  -StorageBucket null `
 -Success null
```

- Convert the resource to JSON
```powershell
$AddStorageBuckets200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

