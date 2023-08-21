# DashboardActivity
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **String** |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**ActivityType** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Message** | **String** |  | [optional] 
**ObjectType** | **String** |  | [optional] 
**ObjectId** | **Decimal** |  | [optional] 
**User** | [**DashboardUser**](DashboardUser.md) |  | [optional] 
**Ts** | **String** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$DashboardActivity = Initialize-PSOpenAPIToolsDashboardActivity  -Id null `
 -Success null `
 -ActivityType null `
 -Name null `
 -Message null `
 -ObjectType null `
 -ObjectId null `
 -User null `
 -Ts null
```

- Convert the resource to JSON
```powershell
$DashboardActivity | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

