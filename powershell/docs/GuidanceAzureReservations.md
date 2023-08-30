# GuidanceAzureReservations
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**DateCreated** | **System.DateTime** |  | [optional] 
**LastUpdated** | **System.DateTime** |  | [optional] 
**ActionCategory** | **String** |  | [optional] 
**ActionMessage** | **String** |  | [optional] 
**ActionTitle** | **String** |  | [optional] 
**ActionType** | **String** |  | [optional] 
**ActionValue** | **String** |  | [optional] 
**ActionValueType** | **String** |  | [optional] 
**ActionPlanId** | **String** |  | [optional] 
**StatusMessage** | **String** |  | [optional] 
**AccountId** | **Int64** |  | [optional] 
**UserId** | **String** |  | [optional] 
**SiteId** | **Int64** |  | [optional] 
**Zone** | [**GuidanceVmwareSizingZone**](GuidanceVmwareSizingZone.md) |  | [optional] 
**State** | **String** |  | [optional] 
**StateMessage** | **String** |  | [optional] 
**Severity** | **String** |  | [optional] 
**Resolved** | **Boolean** |  | [optional] 
**ResolvedMessage** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**RefName** | **String** |  | [optional] 
**Type** | [**GuidanceVmwareSizingType**](GuidanceVmwareSizingType.md) |  | [optional] 
**Savings** | [**GuidanceVmwareSizingSavings**](GuidanceVmwareSizingSavings.md) |  | [optional] 
**Config** | [**GuidanceAzureReservationsConfig**](GuidanceAzureReservationsConfig.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GuidanceAzureReservations = Initialize-PSOpenAPIToolsGuidanceAzureReservations  -Id null `
 -DateCreated null `
 -LastUpdated null `
 -ActionCategory null `
 -ActionMessage null `
 -ActionTitle null `
 -ActionType null `
 -ActionValue null `
 -ActionValueType null `
 -ActionPlanId null `
 -StatusMessage null `
 -AccountId null `
 -UserId null `
 -SiteId null `
 -Zone null `
 -State null `
 -StateMessage null `
 -Severity null `
 -Resolved null `
 -ResolvedMessage null `
 -RefType null `
 -RefId null `
 -RefName null `
 -Type null `
 -Savings null `
 -Config null
```

- Convert the resource to JSON
```powershell
$GuidanceAzureReservations | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

