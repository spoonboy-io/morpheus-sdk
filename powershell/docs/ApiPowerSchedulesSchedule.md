# ApiPowerSchedulesSchedule
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the power schedule | 
**Description** | **String** | A description for the power schedule | [optional] 
**ScheduleType** | **String** | Type of schedule &#x60;power&#x60; on or &#x60;power off&#x60; | [optional] 
**ScheduleTimezone** | **String** | Time Zone eg. America/New_York, Europe/Amsterdam, etc. | [optional] [default to "UTC"]
**Enabled** | **Boolean** | Is the power schedule enabled | [optional] [default to $true]
**MondayOnTime** | **String** | Monday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**MondayOffTime** | **String** | Monday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**TuesdayOnTime** | **String** | Tuesday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**TuesdayOffTime** | **String** | Tuesday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**WednesdayOnTime** | **String** | Wednesday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**WednesdayOffTime** | **String** | Wednesday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**ThursdayOnTime** | **String** | Thursday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**ThursdayOffTime** | **String** | Thursday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**FridayOnTime** | **String** | Friday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**FridayOffTime** | **String** | Friday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**SaturdayOnTime** | **String** | Saturday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**SaturdayOffTime** | **String** | Saturday Off time of the day in 24-hour format | [optional] [default to "24:00"]
**SundayOnTime** | **String** | Sunday Start time of the day in 24-hour format | [optional] [default to "00:00"]
**SundayOffTime** | **String** | Sunday Off time of the day in 24-hour format | [optional] [default to "24:00"]

## Examples

- Prepare the resource
```powershell
$ApiPowerSchedulesSchedule = Initialize-PSOpenAPIToolsApiPowerSchedulesSchedule  -Name Sample Threshold `
 -Description null `
 -ScheduleType null `
 -ScheduleTimezone null `
 -Enabled null `
 -MondayOnTime null `
 -MondayOffTime null `
 -TuesdayOnTime null `
 -TuesdayOffTime null `
 -WednesdayOnTime null `
 -WednesdayOffTime null `
 -ThursdayOnTime null `
 -ThursdayOffTime null `
 -FridayOnTime null `
 -FridayOffTime null `
 -SaturdayOnTime null `
 -SaturdayOffTime null `
 -SundayOnTime null `
 -SundayOffTime null
```

- Convert the resource to JSON
```powershell
$ApiPowerSchedulesSchedule | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

