# HealthRabbit
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**BusyQueues** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**ErrorQueues** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Status** | **String** |  | [optional] 
**Queues** | [**HealthRabbitQueuesInner[]**](HealthRabbitQueuesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthRabbit = Initialize-PSOpenAPIToolsHealthRabbit  -Success null `
 -BusyQueues null `
 -ErrorQueues null `
 -Status null `
 -Queues null
```

- Convert the resource to JSON
```powershell
$HealthRabbit | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

