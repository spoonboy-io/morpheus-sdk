# ApiJobsIdJob
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | A name for the Job | [optional] 
**Labels** | **String[]** | Array of label strings, can be used for filtering. | [optional] 
**Enabled** | **Boolean** | Use this to set enabled state | [optional] [default to $true]
**Task** | [**ApiJobsIdJobTask**](ApiJobsIdJobTask.md) |  | [optional] 
**Workflow** | [**ApiJobsIdJobTask**](ApiJobsIdJobTask.md) |  | [optional] 
**ScanPath** | **String** | Scan Checklist. Only applies to type scap-package. | [optional] 
**SecurityProfile** | **String** | Security Profile. Only applies to type scap-package. | [optional] 
**TargetType** | **String** | Target type where job will execute | [optional] 
**Targets** | [**ApiJobsIdJobTargets[]**](ApiJobsIdJobTargets.md) |  | [optional] 
**InstanceLabel** | **String** | Instance Label. Only applicable if &#x60;targetType&#x60; is &#x60;instance-label&#x60;. | [optional] 
**ServerLabel** | **String** | Server Label. Only applicable if &#x60;targetType&#x60; is &#x60;server-label&#x60;. | [optional] 
**ScheduleMode** | [**OneOfstringlong**](OneOfstringlong.md) |  | [optional] 
**CustomOptions** | [**SystemCollectionsHashtable**](.md) | Map of options to be used as values in the workflow tasks. These correspond to option types. | [optional] 
**CustomConfig** | **String** | Job custom configuration (String in JSON format) | [optional] 
**DateTime** | **System.DateTime** | Date and Time to execute the job. Use UTC time in the format 2020-02-15T05:00:00Z. Required when scheduleMode is &#39;dateTime&#39;. | [optional] 
**Run** | **Boolean** | If true, executes job | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiJobsIdJob = Initialize-PSOpenAPIToolsApiJobsIdJob  -Name Sample Job Name `
 -Labels null `
 -Enabled null `
 -Task null `
 -Workflow null `
 -ScanPath /test_CentOS_Linux_7_Benchmark_v3/test_CentOS_Linux_7_Benchmark_v3.1.1-xccdf.xml `
 -SecurityProfile xccdf_org.cisecurity.benchmarks_profile_Level_2_-_Server `
 -TargetType null `
 -Targets null `
 -InstanceLabel null `
 -ServerLabel null `
 -ScheduleMode null `
 -CustomOptions null `
 -CustomConfig null `
 -DateTime 2020-02-15T05:00Z `
 -Run null
```

- Convert the resource to JSON
```powershell
$ApiJobsIdJob | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

