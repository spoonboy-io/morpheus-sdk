# Group
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Uuid** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Location** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**Config** | [**GroupConfig**](GroupConfig.md) |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Zones** | [**InlineResponse20040AppDeployInstance[]**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Stats** | [**GroupStats**](GroupStats.md) |  | [optional] 
**ServerCount** | **Int64** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$Group = Initialize-PSOpenAPIToolsGroup  -Id null `
 -Uuid null `
 -Name null `
 -Code null `
 -Location null `
 -AccountId null `
 -Active null `
 -Config null `
 -DateCreated null `
 -LastUpdated null `
 -Zones null `
 -Stats null `
 -ServerCount null
```

- Convert the resource to JSON
```powershell
$Group | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

