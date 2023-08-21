# MorpheusApi.MonitoringSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**autoManageChecks** | **Boolean** | Auto Create Checks | [optional] 
**availabilityTimeFrame** | **Number** | Availability Time Frame. The number of days availability should be calculated for. Changes will not take effect until your checks have passed their check interval. | [optional] 
**availabilityPrecision** | **Number** | Availability Precision. The number of decimal places availability should be displayed in. Can be anywhere between 0 and 5. | [optional] 
**defaultCheckInterval** | **Number** | Default Check Interval. The number of minutes to use as the default interval to use when creating new checks. | [optional] 
**serviceNow** | [**MonitoringSettingsServiceNow**](MonitoringSettingsServiceNow.md) |  | [optional] 


