# CheckApp
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Account** | [**UpdateBlueprintPermissionsRequestResourcePermissionSitesInner**](UpdateBlueprintPermissionsRequestResourcePermissionSitesInner.md) |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**App** | [**CheckGroupInstance**](CheckGroupInstance.md) |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**InUptime** | **Boolean** |  | [optional] 
**LastCheckStatus** | **String** |  | [optional] 
**LastWarningDate** | **System.DateTime** |  | [optional] 
**LastErrorDate** | **System.DateTime** |  | [optional] 
**LastSuccessDate** | **System.DateTime** |  | [optional] 
**LastRunDate** | **System.DateTime** |  | [optional] 
**LastError** | **String** |  | [optional] 
**LastTimer** | **Int64** |  | [optional] 
**Health** | **Int64** |  | [optional] 
**History** | **String** |  | [optional] 
**Severity** | **String** |  | [optional] 
**CreateIncident** | **Boolean** |  | [optional] 
**Muted** | **Boolean** |  | [optional] 
**CreatedBy** | [**ActivityActivityInnerUser**](ActivityActivityInnerUser.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Availability** | **String** |  | [optional] 
**Checks** | **Int64[]** |  | [optional] 
**CheckGroups** | **Int64[]** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CheckApp = Initialize-PSOpenAPIToolsCheckApp  -Id null `
 -Account null `
 -Active null `
 -App null `
 -Name null `
 -Description null `
 -InUptime null `
 -LastCheckStatus null `
 -LastWarningDate null `
 -LastErrorDate null `
 -LastSuccessDate null `
 -LastRunDate null `
 -LastError null `
 -LastTimer null `
 -Health null `
 -History null `
 -Severity null `
 -CreateIncident null `
 -Muted null `
 -CreatedBy null `
 -DateCreated null `
 -LastUpdated null `
 -Availability null `
 -Checks null `
 -CheckGroups null
```

- Convert the resource to JSON
```powershell
$CheckApp | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

