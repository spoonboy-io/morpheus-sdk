

# CheckWeb

Web check type allows you to perform a standard web request and validate the response came back successfully.  Additionally, you can check for matching text within the result. 
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Unique name scoped to your account for the check |  [optional]
**description** | **String** | Optional description field |  [optional]
**checkType** | [**CheckWebCheckType**](CheckWebCheckType.md) |  |  [optional]
**checkInterval** | **Integer** | Number of seconds you want between check executions (minimum value is 60, depending on your subscription plan) |  [optional]
**inUptime** | **Boolean** | Used to determine if check should affect account wide availability calculations |  [optional]
**active** | **Boolean** | Used to determine if check should be scheduled to execute |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Severity level threshold for sending notifications. |  [optional]
**config** | [**CheckWebConfig**](CheckWebConfig.md) |  |  [optional]



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



