

# ApiPowerSchedulesIdSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the power schedule |  [optional]
**description** | **String** | A description for the power schedule |  [optional]
**scheduleType** | [**ScheduleTypeEnum**](#ScheduleTypeEnum) | Type of schedule &#x60;power&#x60; on or &#x60;power off&#x60; |  [optional]
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. |  [optional]
**enabled** | **Boolean** | Is the power schedule enabled |  [optional]
**mondayOnTime** | **String** | Monday Start time of the day in 24-hour format |  [optional]
**mondayOffTime** | **String** | Monday Off time of the day in 24-hour format |  [optional]
**tuesdayOnTime** | **String** | Tuesday Start time of the day in 24-hour format |  [optional]
**tuesdayOffTime** | **String** | Tuesday Off time of the day in 24-hour format |  [optional]
**wednesdayOnTime** | **String** | Wednesday Start time of the day in 24-hour format |  [optional]
**wednesdayOffTime** | **String** | Wednesday Off time of the day in 24-hour format |  [optional]
**thursdayOnTime** | **String** | Thursday Start time of the day in 24-hour format |  [optional]
**thursdayOffTime** | **String** | Thursday Off time of the day in 24-hour format |  [optional]
**fridayOnTime** | **String** | Friday Start time of the day in 24-hour format |  [optional]
**fridayOffTime** | **String** | Friday Off time of the day in 24-hour format |  [optional]
**saturdayOnTime** | **String** | Saturday Start time of the day in 24-hour format |  [optional]
**saturdayOffTime** | **String** | Saturday Off time of the day in 24-hour format |  [optional]
**sundayOnTime** | **String** | Sunday Start time of the day in 24-hour format |  [optional]
**sundayOffTime** | **String** | Sunday Off time of the day in 24-hour format |  [optional]



## Enum: ScheduleTypeEnum

Name | Value
---- | -----
POWER | &quot;power&quot;
POWER_OFF | &quot;power off&quot;



