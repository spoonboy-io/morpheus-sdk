# ClusterJobs
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**Type** | [**InlineResponse20094Network**](InlineResponse20094Network.md) |  | [optional] 
**JobSummary** | **String** |  | [optional] 
**ScheduleMode** | **String** |  | [optional] 
**DateTime** | **System.DateTime** |  | [optional] 
**Status** | **String** |  | [optional] 
**Namespace** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**LastRun** | **System.DateTime** |  | [optional] 
**LastResult** | **String** |  | [optional] 
**CreatedBy** | [**InlineResponse20083LoadBalancerNodeCreatedBy**](InlineResponse20083LoadBalancerNodeCreatedBy.md) |  | [optional] 
**TargetType** | **String** |  | [optional] 
**Targets** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**CustomConfig** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**CustomOptions** | [**SystemCollectionsHashtable**](.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterJobs = Initialize-PSOpenAPIToolsClusterJobs  -Id null `
 -Name null `
 -Labels null `
 -Type null `
 -JobSummary null `
 -ScheduleMode null `
 -DateTime null `
 -Status null `
 -Namespace null `
 -Category null `
 -Description null `
 -Enabled null `
 -DateCreated null `
 -LastUpdated null `
 -LastRun null `
 -LastResult null `
 -CreatedBy null `
 -TargetType null `
 -Targets null `
 -CustomConfig null `
 -CustomOptions null
```

- Convert the resource to JSON
```powershell
$ClusterJobs | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

