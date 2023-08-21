

# ApiMonitoringIncidentsIdIncident

Payload for update an incident
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**resolution** | **String** | Description of the resolution to this incident |  [optional]
**comment** | **String** | Comment on this incident, updates summary field |  [optional]
**status** | [**StatusEnum**](#StatusEnum) | Set status |  [optional]
**severity** | [**SeverityEnum**](#SeverityEnum) | Set severity |  [optional]
**name** | **String** | Set display name |  [optional]
**startDate** | **OffsetDateTime** | Set start time |  [optional]
**endDate** | **OffsetDateTime** | Set start time |  [optional]
**inUptime** | **Boolean** | Set &#39;In Availability&#39; |  [optional]



## Enum: StatusEnum

Name | Value
---- | -----
OPEN | &quot;open&quot;
CLOSED | &quot;closed&quot;



## Enum: SeverityEnum

Name | Value
---- | -----
INFO | &quot;info&quot;
WARNING | &quot;warning&quot;
CRITICAL | &quot;critical&quot;



