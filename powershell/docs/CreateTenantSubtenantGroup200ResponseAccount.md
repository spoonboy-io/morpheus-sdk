# CreateTenantSubtenantGroup200ResponseAccount
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Location** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**Zones** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 
**Stats** | [**GroupStats**](GroupStats.md) |  | [optional] 
**ServerCount** | **Int64** |  | [optional] 
**Success** | **Boolean** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$CreateTenantSubtenantGroup200ResponseAccount = Initialize-PSOpenAPIToolsCreateTenantSubtenantGroup200ResponseAccount  -Id null `
 -Name null `
 -Code null `
 -Location null `
 -AccountId null `
 -Visibility null `
 -Active null `
 -DateCreated null `
 -LastUpdated null `
 -Zones null `
 -Stats null `
 -ServerCount null `
 -Success null
```

- Convert the resource to JSON
```powershell
$CreateTenantSubtenantGroup200ResponseAccount | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

