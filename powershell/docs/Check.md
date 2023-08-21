# Check
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**ApiKey** | **String** |  | [optional] 
**Availability** | **Decimal** |  | [optional] 
**CheckAgent** | **String** |  | [optional] 
**CheckInterval** | **Int64** |  | [optional] 
**CheckSpec** | **String** |  | [optional] 
**CheckType** | [**CheckCheckType**](CheckCheckType.md) |  | [optional] 
**Config** | [**AnyOfcheckWebConfigcheckSqlConfigcheckElasticsearchConfigcheckSocketConfigobjectcheckVmConfig**](AnyOfcheckWebConfigcheckSqlConfigcheckElasticsearchConfigcheckSocketConfigobjectcheckVmConfig.md) |  | [optional] 
**Container** | [**CheckContainer**](CheckContainer.md) |  | [optional] 
**CreateIncident** | **Boolean** |  | [optional] 
**Muted** | **Boolean** |  | [optional] 
**CreatedBy** | [**InlineResponse20083LoadBalancerNodeCreatedBy**](InlineResponse20083LoadBalancerNodeCreatedBy.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**Description** | **String** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Health** | **Int64** |  | [optional] 
**InUptime** | **Boolean** |  | [optional] 
**LastBoxStats** | **String** |  | [optional] 
**LastCheckStatus** | **String** |  | [optional] 
**LastError** | **String** |  | [optional] 
**LastErrorDate** | **System.DateTime** |  | [optional] 
**LastMessage** | **String** |  | [optional] 
**LastMetric** | **String** |  | [optional] 
**LastRunDate** | **System.DateTime** |  | [optional] 
**LastStats** | **String** |  | [optional] 
**LastSuccessDate** | **System.DateTime** |  | [optional] 
**LastTimer** | **Int64** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**LastWarningDate** | **System.DateTime** |  | [optional] 
**Name** | **String** |  | [optional] 
**NextRunDate** | **System.DateTime** |  | [optional] 
**OutageTime** | **Int64** |  | [optional] 
**Severity** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**Deleted** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Check = Initialize-PSOpenAPIToolsCheck  -Id null `
 -Account null `
 -Active null `
 -ApiKey null `
 -Availability null `
 -CheckAgent null `
 -CheckInterval null `
 -CheckSpec null `
 -CheckType null `
 -Config null `
 -Container null `
 -CreateIncident null `
 -Muted null `
 -CreatedBy null `
 -DateCreated null `
 -Description null `
 -EndDate null `
 -Health null `
 -InUptime null `
 -LastBoxStats null `
 -LastCheckStatus null `
 -LastError null `
 -LastErrorDate null `
 -LastMessage null `
 -LastMetric null `
 -LastRunDate null `
 -LastStats null `
 -LastSuccessDate null `
 -LastTimer null `
 -LastUpdated null `
 -LastWarningDate null `
 -Name null `
 -NextRunDate null `
 -OutageTime null `
 -Severity null `
 -StartDate null `
 -Deleted null
```

- Convert the resource to JSON
```powershell
$Check | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

