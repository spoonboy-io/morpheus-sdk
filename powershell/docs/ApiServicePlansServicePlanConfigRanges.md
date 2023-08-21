# ApiServicePlansServicePlanConfigRanges
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | **String** | Custom min storage in GB | [optional] 
**MaxStorage** | **String** | Custom max storage in GB | [optional] 
**MinPerDiskSize** | **String** | Custom min per disk size in GB | [optional] 
**MaxPerDiskSize** | **String** | Custom max per disk size in GB | [optional] 
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
$ApiServicePlansServicePlanConfigRanges = Initialize-PSOpenAPIToolsApiServicePlansServicePlanConfigRanges  -MinStorage null `
 -MaxStorage null `
 -MinPerDiskSize null `
 -MaxPerDiskSize null `
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
$ApiServicePlansServicePlanConfigRanges | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

