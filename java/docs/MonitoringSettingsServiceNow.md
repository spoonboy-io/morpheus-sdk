

# MonitoringSettingsServiceNow

Service Now Settings
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **Boolean** | Enabled |  [optional]
**integration** | [**MonitoringSettingsServiceNowIntegration**](MonitoringSettingsServiceNowIntegration.md) |  |  [optional]
**newIncidentAction** | [**NewIncidentActionEnum**](#NewIncidentActionEnum) | New Incident Action |  [optional]
**closeIncidentAction** | [**CloseIncidentActionEnum**](#CloseIncidentActionEnum) | Close Incident Action |  [optional]
**infoMapping** | [**InfoMappingEnum**](#InfoMappingEnum) | Info Mapping |  [optional]
**warningMapping** | [**WarningMappingEnum**](#WarningMappingEnum) | Warning Mapping |  [optional]
**criticalMapping** | [**CriticalMappingEnum**](#CriticalMappingEnum) | Critical Mapping |  [optional]



## Enum: NewIncidentActionEnum

Name | Value
---- | -----
CREATE | &quot;create&quot;
NONE | &quot;none&quot;



## Enum: CloseIncidentActionEnum

Name | Value
---- | -----
CLOSE | &quot;close&quot;
ACTIVITY | &quot;activity&quot;
NONE | &quot;none&quot;



## Enum: InfoMappingEnum

Name | Value
---- | -----
LOW | &quot;low&quot;
MEDIUM | &quot;medium&quot;
HIGH | &quot;high&quot;



## Enum: WarningMappingEnum

Name | Value
---- | -----
LOW | &quot;low&quot;
MEDIUM | &quot;medium&quot;
HIGH | &quot;high&quot;



## Enum: CriticalMappingEnum

Name | Value
---- | -----
LOW | &quot;low&quot;
MEDIUM | &quot;medium&quot;
HIGH | &quot;high&quot;



