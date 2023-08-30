# HealthDatabase


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** |  | [optional] 
**connection_list** | **[{str: (bool, date, datetime, dict, float, int, list, str, none_type)}], none_type** |  | [optional] 
**busy_connections** | **[str, none_type]** |  | [optional] 
**max_connections** | **int** |  | [optional] 
**max_used_connections** | **int** |  | [optional] 
**used_connections** | **int** |  | [optional] 
**aborted_connections** | **int** |  | [optional] 
**innodb_status** | **str, none_type** |  | [optional] 
**stats** | [**HealthDatabaseStats**](HealthDatabaseStats.md) |  | [optional] 
**scans** | [**HealthDatabaseScans**](HealthDatabaseScans.md) |  | [optional] 
**slow_queries** | [**[HealthDatabaseSlowQueriesInner]**](HealthDatabaseSlowQueriesInner.md) |  | [optional] 
**innodb_stats** | [**HealthDatabaseInnodbStats**](HealthDatabaseInnodbStats.md) |  | [optional] 
**scan_percent** | **float** |  | [optional] 
**status** | **str** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


