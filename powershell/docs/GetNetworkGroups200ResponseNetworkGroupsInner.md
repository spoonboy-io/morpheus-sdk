# GetNetworkGroups200ResponseNetworkGroupsInner
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Description** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**Networks** | **Int64[]** |  | [optional] 
**Subnets** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 
**Tenants** | [**ListDeploys200ResponseAllOfAppDeploysInnerInstance[]**](ListDeploys200ResponseAllOfAppDeploysInnerInstance.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$GetNetworkGroups200ResponseNetworkGroupsInner = Initialize-PSOpenAPIToolsGetNetworkGroups200ResponseNetworkGroupsInner  -Id null `
 -Name null `
 -Description null `
 -Visibility null `
 -Active null `
 -Networks null `
 -Subnets null `
 -Tenants null
```

- Convert the resource to JSON
```powershell
$GetNetworkGroups200ResponseNetworkGroupsInner | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

