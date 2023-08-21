# ApiServicePlansIdServicePlan
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Service plan name | [optional] 
**Code** | **String** | Service plan code, must be unique | [optional] 
**Description** | **String** | Service plan description | [optional] 
**Editable** | **Boolean** | Can be used to enable / disable the editability of the service plan. | [optional] [default to $true]
**MaxStorage** | **Decimal** | Max storage size in bytes | [optional] 
**MaxMemory** | **Decimal** | Max memory size in bytes | [optional] 
**MaxCores** | **Decimal** | Max cores | [optional] 
**MaxDisks** | **Decimal** | Max disks allowed | [optional] 
**ProvisionType** | [**ApiServicePlansIdServicePlanProvisionType[]**](ApiServicePlansIdServicePlanProvisionType.md) |  | [optional] 
**CustomCores** | **Boolean** | Can be used to enable / disable customizable cores | [optional] [default to $false]
**CustomMaxStorage** | **Boolean** | Can be used to enable / disable customizable storage | [optional] [default to $false]
**CustomMaxDataStorage** | **Boolean** | Can be used to enable / disable customizable extra volumes. | [optional] [default to $false]
**CustomMaxMemory** | **Boolean** | Can be used to enable / disable customizable memory. | [optional] [default to $false]
**AddVolumes** | **Boolean** | Can be used to enable / disable ability to add volumes | [optional] [default to $false]
**SortOrder** | **Decimal** | Sort order | [optional] 
**PriceSets** | **Int64[]** | List of price sets to include in service plan | [optional] 
**Config** | [**ApiServicePlansIdServicePlanConfig**](ApiServicePlansIdServicePlanConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServicePlansIdServicePlan = Initialize-PSOpenAPIToolsApiServicePlansIdServicePlan  -Name null `
 -Code null `
 -Description null `
 -Editable null `
 -MaxStorage null `
 -MaxMemory null `
 -MaxCores null `
 -MaxDisks null `
 -ProvisionType null `
 -CustomCores null `
 -CustomMaxStorage null `
 -CustomMaxDataStorage null `
 -CustomMaxMemory null `
 -AddVolumes null `
 -SortOrder null `
 -PriceSets null `
 -Config null
```

- Convert the resource to JSON
```powershell
$ApiServicePlansIdServicePlan | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

