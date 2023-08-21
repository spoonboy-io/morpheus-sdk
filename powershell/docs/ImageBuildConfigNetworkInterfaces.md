# ImageBuildConfigNetworkInterfaces
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryInterface** | **Boolean** |  | [optional] 
**Network** | [**ZoneVcenterConfigNetworkServer**](ZoneVcenterConfigNetworkServer.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ImageBuildConfigNetworkInterfaces = Initialize-PSOpenAPIToolsImageBuildConfigNetworkInterfaces  -PrimaryInterface null `
 -Network null
```

- Convert the resource to JSON
```powershell
$ImageBuildConfigNetworkInterfaces | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

