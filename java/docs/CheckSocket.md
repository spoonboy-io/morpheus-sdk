

# CheckSocket

Socket check confirms a certain TCP port is up and responding in your environment.  It can be configured do an initial send upon connect and compare and expected response of the service. 
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check |  [optional]
**description** | **String** | Optional description field |  [optional]
**checkType** | [**CheckSocketCheckType**](CheckSocketCheckType.md) |  |  [optional]
**checkInterval** | **Integer** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**active** | **Boolean** | Used to determine if check should be scheduled to execute |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Severity level threshold for sending notifications. |  [optional]
**config** | [**CheckSocketConfig**](CheckSocketConfig.md) |  |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



