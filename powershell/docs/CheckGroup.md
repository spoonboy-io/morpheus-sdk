# CheckGroup
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Instance** | [**InlineResponse20082LoadBalancerInstanceSslCert**](InlineResponse20082LoadBalancerInstanceSslCert.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**InUptime** | **Boolean** |  | [optional] 
**LastCheckStatus** | **String** |  | [optional] 
**LastWarningDate** | **System.DateTime** |  | [optional] 
**LastErrorDate** | **System.DateTime** |  | [optional] 
**LastSuccessDate** | **System.DateTime** |  | [optional] 
**LastRunDate** | **System.DateTime** |  | [optional] 
**LastError** | **String** |  | [optional] 
**OutageTime** | **Int64** |  | [optional] 
**LastTimer** | **Int64** |  | [optional] 
**Health** | **Int64** |  | [optional] 
**History** | **String** |  | [optional] 
**MinHappy** | **Int64** |  | [optional] 
**LastMetric** | **String** |  | [optional] 
**Severity** | **String** |  | [optional] 
**CreateIncident** | **Boolean** |  | [optional] 
**Muted** | **Boolean** |  | [optional] 
**CreatedBy** | [**InlineResponse200107NetworkPoolCreatedBy**](InlineResponse200107NetworkPoolCreatedBy.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Availability** | **Decimal** |  | [optional] 
**CheckType** | [**CheckCheckType**](CheckCheckType.md) |  | [optional] 
**Checks** | **Int64[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckGroup = Initialize-PSOpenAPIToolsCheckGroup  -Id null `
 -Account null `
 -Instance null `
 -Name null `
 -Description null `
 -InUptime null `
 -LastCheckStatus null `
 -LastWarningDate null `
 -LastErrorDate null `
 -LastSuccessDate null `
 -LastRunDate null `
 -LastError null `
 -OutageTime null `
 -LastTimer null `
 -Health null `
 -History null `
 -MinHappy null `
 -LastMetric null `
 -Severity null `
 -CreateIncident null `
 -Muted null `
 -CreatedBy null `
 -DateCreated null `
 -LastUpdated null `
 -Availability null `
 -CheckType null `
 -Checks null
```

- Convert the resource to JSON
```powershell
$CheckGroup | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

