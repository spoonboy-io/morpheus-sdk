# AddJobsRequestJob
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the Job | 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | **Boolean** | Use this to set enabled state | [optional] [default to $true]
**Task** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**Workflow** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**TargetType** | **String** | Target type where job will execute | 
**Targets** | [**WorkflowJobPayloadTargetsInner[]**](WorkflowJobPayloadTargetsInner.md) |  | 
**InstanceLabel** | **String** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | **String** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**ScheduleMode** | [**WorkflowJobPayloadScheduleMode**](WorkflowJobPayloadScheduleMode.md) |  | 
**CustomOptions** | [**SystemCollectionsHashtable**](.md) | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | **String** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | **System.DateTime** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | **Boolean** | If true, executes job | [optional] 
**SecurityPackage** | [**WorkflowJobPayloadTask**](WorkflowJobPayloadTask.md) |  | 
**ScanPath** | **String** | Scan Checklist | [optional] 
**SecurityProfile** | **String** | Security Profile | [optional] 

## Examples

- Prepare the resource
```powershell
$AddJobsRequestJob = Initialize-PSOpenAPIToolsAddJobsRequestJob  -Name Sample Job Name `
 -Labels null `
 -Enabled null `
 -Task null `
 -Workflow null `
 -TargetType null `
 -Targets null `
 -InstanceLabel null `
 -ServerLabel null `
 -ScheduleMode null `
 -CustomOptions null `
 -CustomConfig null `
 -DateTime 2020-02-15T05:00Z `
 -Run null `
 -SecurityPackage null `
 -ScanPath /test_CentOS_Linux_7_Benchmark_v3/test_CentOS_Linux_7_Benchmark_v3.1.1-xccdf.xml `
 -SecurityProfile xccdf_org.cisecurity.benchmarks_profile_Level_2_-_Server
```

- Convert the resource to JSON
```powershell
$AddJobsRequestJob | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

