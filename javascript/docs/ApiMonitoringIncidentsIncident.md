# MorpheusApi.ApiMonitoringIncidentsIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resolution** | **String** | Description of the resolution to this incident | [optional] 
**comment** | **String** | Comment on this incident, updates summary field | [optional] 
**status** | **String** | Set status | [optional] 
**severity** | **String** | Set severity | [optional] 
**name** | **String** | Set display name | 
**startDate** | **Date** | Set start time | [optional] 
**endDate** | **Date** | Set start time | [optional] 
**inUptime** | **Boolean** | Set &#39;In Availability&#39; | [optional] 



## Enum: StatusEnum


* `open` (value: `"open"`)

* `closed` (value: `"closed"`)





## Enum: SeverityEnum


* `info` (value: `"info"`)

* `warning` (value: `"warning"`)

* `critical` (value: `"critical"`)




