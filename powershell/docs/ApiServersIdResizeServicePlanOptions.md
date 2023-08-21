# ApiServersIdResizeServicePlanOptions
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCores** | **Int64** | Core Count | [optional] 
**CoresPerSocket** | **Int64** | Cores Per Socket | [optional] 
**MaxMemory** | **Int64** | Memory in bytes For backwards compatability, values less than 1048576 are treated as being in MB and will be converted to bytes | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServersIdResizeServicePlanOptions = Initialize-PSOpenAPIToolsApiServersIdResizeServicePlanOptions  -MaxCores null `
 -CoresPerSocket null `
 -MaxMemory null
```

- Convert the resource to JSON
```powershell
$ApiServersIdResizeServicePlanOptions | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

