# ActivityActivityInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **String** |  | [optional] 
**Success** | **Boolean** |  | [optional] 
**ActivityType** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Message** | **String** |  | [optional] 
**ObjectType** | **String** |  | [optional] 
**ObjectId** | **Int64** |  | [optional] 
**User** | [**ActivityActivityInnerUser**](ActivityActivityInnerUser.md) |  | [optional] 
**Ts** | **System.DateTime** |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ActivityActivityInner = Initialize-PSOpenAPIToolsActivityActivityInner  -Id null `
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
$ActivityActivityInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

