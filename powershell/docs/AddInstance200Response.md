# AddInstance200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**InstanceCreateSuccessInstance**](InstanceCreateSuccessInstance.md) |  | 
**ZoneId** | **Int64** | The Cloud ID to provision the instance onto. | 
**Success** | **Boolean** |  | [optional] 
**Errors** | [**SystemCollectionsHashtable**](.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$AddInstance200Response = Initialize-PSOpenAPIToolsAddInstance200Response  -Instance null `
 -ZoneId 6 `
 -Success null `
 -Errors null
```

- Convert the resource to JSON
```powershell
$AddInstance200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

