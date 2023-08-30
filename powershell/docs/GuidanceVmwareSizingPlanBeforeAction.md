# GuidanceVmwareSizingPlanBeforeAction
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**SortOrder** | **Int64** |  | [optional] 
**Description** | **String** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**MaxMemory** | **Int64** |  | [optional] 
**MaxCpu** | **String** |  | [optional] 
**MaxCores** | **Int64** |  | [optional] 
**MaxDisks** | **String** |  | [optional] 
**CoresPerSocket** | **Int64** |  | [optional] 
**CustomCpu** | **Boolean** |  | [optional] 
**CustomCores** | **Boolean** |  | [optional] 
**CustomMaxStorage** | **Boolean** |  | [optional] 
**CustomMaxDataStorage** | **Boolean** |  | [optional] 
**CustomMaxMemory** | **Boolean** |  | [optional] 
**AddVolumes** | **Boolean** |  | [optional] 
**MemoryOptionSource** | **String** |  | [optional] 
**CpuOptionSource** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**RegionCode** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Editable** | **Boolean** |  | [optional] 
**ProvisionType** | [**GuidanceVmwareSizingPlanBeforeActionProvisionType**](GuidanceVmwareSizingPlanBeforeActionProvisionType.md) |  | [optional] 
**Tenants** | **String** |  | [optional] 
**PriceSets** | [**GuidanceVmwareSizingPlanBeforeActionPriceSetsInner[]**](GuidanceVmwareSizingPlanBeforeActionPriceSetsInner.md) |  | [optional] 
**Config** | [**GuidanceVmwareSizingPlanBeforeActionConfig**](GuidanceVmwareSizingPlanBeforeActionConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceVmwareSizingPlanBeforeAction = Initialize-PSOpenAPIToolsGuidanceVmwareSizingPlanBeforeAction  -Id null `
 -Name null `
 -Code null `
 -Active null `
 -SortOrder null `
 -Description null `
 -MaxStorage null `
 -MaxMemory null `
 -MaxCpu null `
 -MaxCores null `
 -MaxDisks null `
 -CoresPerSocket null `
 -CustomCpu null `
 -CustomCores null `
 -CustomMaxStorage null `
 -CustomMaxDataStorage null `
 -CustomMaxMemory null `
 -AddVolumes null `
 -MemoryOptionSource null `
 -CpuOptionSource null `
 -DateCreated null `
 -LastUpdated null `
 -RegionCode null `
 -Visibility null `
 -Editable null `
 -ProvisionType null `
 -Tenants null `
 -PriceSets null `
 -Config null
```

- Convert the resource to JSON
```powershell
$GuidanceVmwareSizingPlanBeforeAction | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

