# ApiServicePlansServicePlan
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** | Service plan name | 
**Code** | **String** | Service plan code, must be unique | 
**Description** | **String** | Service plan description | [optional] 
**Editable** | **Boolean** | Can be used to enable / disable the editability of the service plan. | [optional] [default to $true]
**MaxStorage** | **Decimal** | Max storage size in bytes | 
**MaxMemory** | **Decimal** | Max memory size in bytes | 
**MaxCores** | **Decimal** | Max cores | [optional] 
**MaxDisks** | **Decimal** | Max disks allowed | [optional] 
**ProvisionType** | [**ApiServicePlansServicePlanProvisionType[]**](ApiServicePlansServicePlanProvisionType.md) |  | 
**CustomCores** | **Boolean** | Can be used to enable / disable customizable cores | [optional] [default to $false]
**CustomMaxStorage** | **Boolean** | Can be used to enable / disable customizable storage | [optional] [default to $false]
**CustomMaxDataStorage** | **Boolean** | Can be used to enable / disable customizable extra volumes. | [optional] [default to $false]
**CustomMaxMemory** | **Boolean** | Can be used to enable / disable customizable memory. | [optional] [default to $false]
**AddVolumes** | **Boolean** | Can be used to enable / disable ability to add volumes | [optional] [default to $false]
**SortOrder** | **Decimal** | Sort order | [optional] 
**PriceSets** | **Int64[]** | List of price sets to include in service plan | [optional] 
**Config** | [**ApiServicePlansServicePlanConfig**](ApiServicePlansServicePlanConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ApiServicePlansServicePlan = Initialize-PSOpenAPIToolsApiServicePlansServicePlan  -Name null `
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
$ApiServicePlansServicePlan | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

