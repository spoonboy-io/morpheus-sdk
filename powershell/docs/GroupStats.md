# GroupStats
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCounts** | [**GroupStatsInstanceCounts**](GroupStatsInstanceCounts.md) |  | [optional] 
**ServerCounts** | [**ZoneStatsServerCounts**](ZoneStatsServerCounts.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GroupStats = Initialize-PSOpenAPIToolsGroupStats  -InstanceCounts null `
 -ServerCounts null
```

- Convert the resource to JSON
```powershell
$GroupStats | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

