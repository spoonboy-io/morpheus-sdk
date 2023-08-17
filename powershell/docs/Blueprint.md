# Blueprint
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Labels** | **String[]** |  | [optional] 
**Type** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Category** | **String** |  | [optional] 
**Config** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Visibility** | **String** |  | [optional] 
**ResourcePermission** | [**SystemCollectionsHashtable**](.md) |  | [optional] 
**Owner** | [**ActivityActivityInnerUser**](ActivityActivityInnerUser.md) |  | [optional] 
**Tenant** | [**ApplianceSettingsEnabledZoneTypesInner**](ApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Blueprint = Initialize-PSOpenAPIToolsBlueprint  -Id null `
 -Name null `
 -Labels null `
 -Type null `
 -Description null `
 -Category null `
 -Config null `
 -Visibility null `
 -ResourcePermission null `
 -Owner null `
 -Tenant null
```

- Convert the resource to JSON
```powershell
$Blueprint | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

