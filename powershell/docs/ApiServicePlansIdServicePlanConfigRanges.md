# ApiServicePlansIdServicePlanConfigRanges
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | **String** | Custom min storage in GB (unless &#x60;storageSizeType&#x60; set to mb) | [optional] 
**MaxStorage** | **String** | Custom max storage in GB (unless &#x60;storageSizeType&#x60; set to mb) | [optional] 
**MinMemory** | **Int64** | Custom min memory in bytes | [optional] 
**MaxMemory** | **Int64** | Custom max memory in bytes | [optional] 
**MinCores** | **String** | Custom min cores | [optional] 
**MaxCores** | **String** | Custom max cores | [optional] 
**MinSockets** | **String** | Custom min sockets | [optional] 
**MaxSockets** | **String** | Custom max sockets | [optional] 
**MinCoresPerSocket** | **String** | Custom min cores allowed per socket | [optional] 
**MaxCoresPerSocket** | **String** | Custom max cores allowed per socket | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServicePlansIdServicePlanConfigRanges = Initialize-PSOpenAPIToolsApiServicePlansIdServicePlanConfigRanges  -MinStorage null `
 -MaxStorage null `
 -MinMemory null `
 -MaxMemory null `
 -MinCores null `
 -MaxCores null `
 -MinSockets null `
 -MaxSockets null `
 -MinCoresPerSocket null `
 -MaxCoresPerSocket null
```

- Convert the resource to JSON
```powershell
$ApiServicePlansIdServicePlanConfigRanges | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

