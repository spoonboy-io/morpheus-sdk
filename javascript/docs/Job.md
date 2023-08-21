# MorpheusApi.Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** |  | [optional] 
**name** | **String** |  | [optional] 
**labels** | **[String]** |  | [optional] 
**type** | [**InlineResponse20094Network**](InlineResponse20094Network.md) |  | [optional] 
**workflow** | [**JobWorkflow**](JobWorkflow.md) |  | [optional] 
**task** | [**JobTask**](JobTask.md) |  | [optional] 
**securityPackage** | [**JobSecurityPackage**](JobSecurityPackage.md) |  | [optional] 
**jobSummary** | **String** |  | [optional] 
**scheduleMode** | [**OneOfstringlong**](OneOfstringlong.md) |  | [optional] 
**dateTime** | **String** |  | [optional] 
**status** | **String** |  | [optional] 
**namespace** | **String** |  | [optional] 
**category** | **String** |  | [optional] 
**description** | **String** |  | [optional] 
**enabled** | **Boolean** |  | [optional] 
**dateCreated** | **Date** |  | [optional] 
**lastUpdated** | **Date** |  | [optional] 
**lastRun** | **Date** |  | [optional] 
**lastResult** | **String** |  | [optional] 
**createdBy** | [**JobCreatedBy**](JobCreatedBy.md) |  | [optional] 
**targetType** | **String** |  | [optional] 
**targets** | [**[JobTargets]**](JobTargets.md) |  | [optional] 
**scanPath** | **String** | Scan Checklist. Only applies to type scap-package. | [optional] 
**securityProfile** | **String** | Security Profile. Only applies to type scap-package. | [optional] 
**customConfig** | **String** |  | [optional] 
**customOptions** | [**JobCustomOptions**](JobCustomOptions.md) |  | [optional] 


