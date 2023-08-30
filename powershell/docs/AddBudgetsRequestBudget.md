# AddBudgetsRequestBudget
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **String** |  | 
**Description** | **String** |  | [optional] 
**Scope** | **String** |  | [optional] [default to "account"]
**Period** | **String** |  | [optional] [default to "year"]
**Year** | **Int64** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**Interval** | **String** |  | [optional] [default to "year"]
**ScopeTenantId** | **Int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;tenant  | [optional] 
**ScopeGroupId** | **Int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;group   | [optional] 
**ScopeCloudId** | **Int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;cloud  | [optional] 
**ScopeUserId** | **Int64** | The Tenant ID to scope to, for use with &#x60;&#x60;scope&#x60;&#x60;&#x3D;user  | [optional] 
**Costs** | **Int64[]** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] [default to $true]

## Examples

- Prepare the resource
```powershell
$AddBudgetsRequestBudget = Initialize-PSOpenAPIToolsAddBudgetsRequestBudget  -Name null `
 -Description null `
 -Scope null `
 -Period null `
 -Year 2020 `
 -StartDate null `
 -EndDate null `
 -Interval null `
 -ScopeTenantId null `
 -ScopeGroupId null `
 -ScopeCloudId null `
 -ScopeUserId null `
 -Costs [100,100,110,120] `
 -Enabled null
```

- Convert the resource to JSON
```powershell
$AddBudgetsRequestBudget | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

