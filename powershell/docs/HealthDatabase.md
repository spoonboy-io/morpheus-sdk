# HealthDatabase
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **Boolean** |  | [optional] 
**ConnectionList** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**BusyConnections** | **String[]** |  | [optional] 
**MaxConnections** | **Int64** |  | [optional] 
**MaxUsedConnections** | **Int64** |  | [optional] 
**UsedConnections** | **Int64** |  | [optional] 
**AbortedConnections** | **Int64** |  | [optional] 
**InnodbStatus** | **String** |  | [optional] 
**Stats** | [**HealthDatabaseStats**](HealthDatabaseStats.md) |  | [optional] 
**Scans** | [**HealthDatabaseScans**](HealthDatabaseScans.md) |  | [optional] 
**SlowQueries** | [**HealthDatabaseSlowQueriesInner[]**](HealthDatabaseSlowQueriesInner.md) |  | [optional] 
**InnodbStats** | [**HealthDatabaseInnodbStats**](HealthDatabaseInnodbStats.md) |  | [optional] 
**ScanPercent** | **Decimal** |  | [optional] 
**Status** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$HealthDatabase = Initialize-PSOpenAPIToolsHealthDatabase  -Success null `
 -ConnectionList null `
 -BusyConnections null `
 -MaxConnections null `
 -MaxUsedConnections null `
 -UsedConnections null `
 -AbortedConnections null `
 -InnodbStatus null `
 -Stats null `
 -Scans null `
 -SlowQueries null `
 -InnodbStats null `
 -ScanPercent null `
 -Status null
```

- Convert the resource to JSON
```powershell
$HealthDatabase | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

