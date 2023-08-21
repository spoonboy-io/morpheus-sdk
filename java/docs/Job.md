

# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Long** |  |  [optional]
**name** | **String** |  |  [optional]
**labels** | **List&lt;String&gt;** |  |  [optional]
**type** | [**InlineResponse20094Network**](InlineResponse20094Network.md) |  |  [optional]
**workflow** | [**JobWorkflow**](JobWorkflow.md) |  |  [optional]
**task** | [**JobTask**](JobTask.md) |  |  [optional]
**securityPackage** | [**JobSecurityPackage**](JobSecurityPackage.md) |  |  [optional]
**jobSummary** | **String** |  |  [optional]
**scheduleMode** | [**OneOfstringlong**](OneOfstringlong.md) |  |  [optional]
**dateTime** | **String** |  |  [optional]
**status** | **String** |  |  [optional]
**namespace** | **String** |  |  [optional]
**category** | **String** |  |  [optional]
**description** | **String** |  |  [optional]
**enabled** | **Boolean** |  |  [optional]
**dateCreated** | **OffsetDateTime** |  |  [optional]
**lastUpdated** | **OffsetDateTime** |  |  [optional]
**lastRun** | **OffsetDateTime** |  |  [optional]
**lastResult** | **String** |  |  [optional]
**createdBy** | [**JobCreatedBy**](JobCreatedBy.md) |  |  [optional]
**targetType** | **String** |  |  [optional]
**targets** | [**List&lt;JobTargets&gt;**](JobTargets.md) |  |  [optional]
**scanPath** | **String** | Scan Checklist. Only applies to type scap-package. |  [optional]
**securityProfile** | **String** | Security Profile. Only applies to type scap-package. |  [optional]
**customConfig** | **String** |  |  [optional]
**customOptions** | [**JobCustomOptions**](JobCustomOptions.md) |  |  [optional]



