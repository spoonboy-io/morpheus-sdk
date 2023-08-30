# GetNetworkRouters200ResponseNetworkRoutersInnerType
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Code** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Enabled** | **Boolean** |  | [optional] 
**Creatable** | **Boolean** |  | [optional] 
**Selectable** | **Boolean** |  | [optional] 
**HasFirewall** | **Boolean** |  | [optional] 
**HasDhcp** | **Boolean** |  | [optional] 
**HasRouting** | **Boolean** |  | [optional] 
**HasNetworkServer** | **Boolean** |  | [optional] 
**OptionTypes** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**RuleOptionTypes** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkRouters200ResponseNetworkRoutersInnerType = Initialize-PSOpenAPIToolsGetNetworkRouters200ResponseNetworkRoutersInnerType  -Id null `
 -Code null `
 -Name null `
 -Description null `
 -Enabled null `
 -Creatable null `
 -Selectable null `
 -HasFirewall null `
 -HasDhcp null `
 -HasRouting null `
 -HasNetworkServer null `
 -OptionTypes null `
 -RuleOptionTypes null
```

- Convert the resource to JSON
```powershell
$GetNetworkRouters200ResponseNetworkRoutersInnerType | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

