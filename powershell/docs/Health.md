# Health
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**ApplianceUrl** | **String** |  | [optional] 
**BuildVersion** | **String** |  | [optional] 
**SetupNeeded** | **Boolean** |  | [optional] 
**Date** | **System.DateTime** |  | [optional] 
**Cpu** | [**HealthCpu**](HealthCpu.md) |  | [optional] 
**Memory** | [**HealthMemory**](HealthMemory.md) |  | [optional] 
**Threads** | [**HealthThreads**](HealthThreads.md) |  | [optional] 
**Database** | [**HealthDatabase**](HealthDatabase.md) |  | [optional] 
**Elastic** | [**HealthElastic**](HealthElastic.md) |  | [optional] 
**Rabbit** | [**HealthRabbit**](HealthRabbit.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Health = Initialize-PSOpenAPIToolsHealth  -Success null `
 -StatusMessage null `
 -ApplianceUrl null `
 -BuildVersion null `
 -SetupNeeded null `
 -Date null `
 -Cpu null `
 -Memory null `
 -Threads null `
 -Database null `
 -Elastic null `
 -Rabbit null
```

- Convert the resource to JSON
```powershell
$Health | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

