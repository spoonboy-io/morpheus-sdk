

# CheckSql

SQL Server check allows to execute a query so that you may validate the value returned in addition to verifying the database is responding.  This can be useful for doing a slow query check or just making sure something isn't growing out of control. 
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check |  [optional]
**description** | **String** | Optional description field |  [optional]
**checkType** | [**CheckSqlCheckType**](CheckSqlCheckType.md) |  |  [optional]
**checkInterval** | **Integer** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**active** | **Boolean** | Used to determine if check should be scheduled to execute |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Severity level threshold for sending notifications. |  [optional]
**config** | [**CheckSqlConfig**](CheckSqlConfig.md) |  |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



