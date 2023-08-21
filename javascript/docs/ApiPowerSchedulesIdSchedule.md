# MorpheusApi.ApiPowerSchedulesIdSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the power schedule | [optional] 
**description** | **String** | A description for the power schedule | [optional] 
**scheduleType** | **String** | Type of schedule &#x60;power&#x60; on or &#x60;power off&#x60; | [optional] 
**scheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to &#39;UTC&#39;]
**enabled** | **Boolean** | Is the power schedule enabled | [optional] [default to true]
**mondayOnTime** | **String** | Monday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**mondayOffTime** | **String** | Monday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]
**tuesdayOnTime** | **String** | Tuesday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**tuesdayOffTime** | **String** | Tuesday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]
**wednesdayOnTime** | **String** | Wednesday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**wednesdayOffTime** | **String** | Wednesday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]
**thursdayOnTime** | **String** | Thursday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**thursdayOffTime** | **String** | Thursday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]
**fridayOnTime** | **String** | Friday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**fridayOffTime** | **String** | Friday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]
**saturdayOnTime** | **String** | Saturday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**saturdayOffTime** | **String** | Saturday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]
**sundayOnTime** | **String** | Sunday Start time of the day in 24-hour format | [optional] [default to &#39;00:00&#39;]
**sundayOffTime** | **String** | Sunday Off time of the day in 24-hour format | [optional] [default to &#39;24:00&#39;]



## Enum: ScheduleTypeEnum


* `power` (value: `"power"`)

* `power off` (value: `"power off"`)




