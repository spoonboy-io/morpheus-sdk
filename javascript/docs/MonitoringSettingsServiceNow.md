# MorpheusApi.MonitoringSettingsServiceNow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **Boolean** | Enabled | [optional] 
**integration** | [**MonitoringSettingsServiceNowIntegration**](MonitoringSettingsServiceNowIntegration.md) |  | [optional] 
**newIncidentAction** | **String** | New Incident Action | [optional] 
**closeIncidentAction** | **String** | Close Incident Action | [optional] 
**infoMapping** | **String** | Info Mapping | [optional] 
**warningMapping** | **String** | Warning Mapping | [optional] 
**criticalMapping** | **String** | Critical Mapping | [optional] 



## Enum: NewIncidentActionEnum


* `create` (value: `"create"`)

* `none` (value: `"none"`)





## Enum: CloseIncidentActionEnum


* `close` (value: `"close"`)

* `activity` (value: `"activity"`)

* `none` (value: `"none"`)





## Enum: InfoMappingEnum


* `low` (value: `"low"`)

* `medium` (value: `"medium"`)

* `high` (value: `"high"`)





## Enum: WarningMappingEnum


* `low` (value: `"low"`)

* `medium` (value: `"medium"`)

* `high` (value: `"high"`)





## Enum: CriticalMappingEnum


* `low` (value: `"low"`)

* `medium` (value: `"medium"`)

* `high` (value: `"high"`)




