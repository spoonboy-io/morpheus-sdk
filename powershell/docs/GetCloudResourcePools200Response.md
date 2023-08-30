# GetCloudResourcePools200Response
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | [**ZoneResourcePool**](ZoneResourcePool.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetCloudResourcePools200Response = Initialize-PSOpenAPIToolsGetCloudResourcePools200Response  -ResourcePool null
```

- Convert the resource to JSON
```powershell
$GetCloudResourcePools200Response | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

