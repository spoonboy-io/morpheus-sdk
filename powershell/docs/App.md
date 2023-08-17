# App
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**Environment** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Account** | [**ApplianceSettingsEnabledZoneTypesInner**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | [**ActivityActivityInnerUser**](ActivityActivityInnerUser.md) |  | [optional] 
**SiteId** | **Int64** |  | [optional] 
**Group** | [**ApplianceSettingsEnabledZoneTypesInner**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Blueprint** | [**AppBlueprint**](AppBlueprint.md) |  | [optional] 
**Type** | **String** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**RemovalDate** | **System.DateTime** |  | [optional] 
**AppContext** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**AppStatus** | **String** |  | [optional] 
**InstanceCount** | **Int64** |  | [optional] 
**ContainerCount** | **Int64** |  | [optional] 
**AppTiers** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Instances** | [**ApplianceSettingsEnabledZoneTypesInner[]**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Stats** | [**AppStats**](AppStats.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$App = Initialize-PSOpenAPIToolsApp  -Id null `
 -Name null `
 -Description null `
 -Labels null `
 -Environment null `
 -AccountId null `
 -Account null `
 -Owner null `
 -SiteId null `
 -Group null `
 -Blueprint null `
 -Type null `
 -DateCreated null `
 -LastUpdated null `
 -RemovalDate null `
 -AppContext null `
 -Status null `
 -AppStatus null `
 -InstanceCount null `
 -ContainerCount null `
 -AppTiers null `
 -Instances null `
 -Stats null
```

- Convert the resource to JSON
```powershell
$App | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

