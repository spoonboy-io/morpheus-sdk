# CatalogItemTypeInstanceScribe
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**CatalogItemTypeInstanceScribeGroup**](CatalogItemTypeInstanceScribeGroup.md) |  | 
**Cloud** | [**CatalogItemTypeInstanceScribeCloud**](CatalogItemTypeInstanceScribeCloud.md) |  | 
**Type** | **String** | The type of instance by code we want to fetch. | 
**Name** | **String** | Name of the instance to be created. | 
**Config** | [**CatalogItemTypeInstanceScribeConfig**](CatalogItemTypeInstanceScribeConfig.md) |  | 
**Volumes** | [**InstanceCreateVolume[]**](InstanceCreateVolume.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**HostName** | **String** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 
**Environment** | **String** | Environment code | [optional] 
**Layout** | [**CatalogItemTypeInstanceScribeLayout**](CatalogItemTypeInstanceScribeLayout.md) |  | 
**Plan** | [**CatalogItemTypeInstanceScribePlan**](CatalogItemTypeInstanceScribePlan.md) |  | 
**Version** | **String** | Version of the layout to create. | [optional] 
**Evars** | [**UpdateHostManagedRequestServerTagsInner[]**](UpdateHostManagedRequestServerTagsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**ServicePlanOptions** | [**ServicePlanOptions**](ServicePlanOptions.md) |  | [optional] 
**SecurityGroups** | [**CatalogItemTypeInstanceScribeSecurityGroupsInner[]**](CatalogItemTypeInstanceScribeSecurityGroupsInner.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**NetworkInterfaces** | [**InstanceCreateNetwork[]**](InstanceCreateNetwork.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Labels** | **String[]** | Array of strings (keywords). | [optional] 
**Tags** | [**UpdateHostManagedRequestServerTagsInner[]**](UpdateHostManagedRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | [**UpdateHostManagedRequestServerTagsInner[]**](UpdateHostManagedRequestServerTagsInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | [**CatalogItemTypeInstanceScribePortsInner[]**](CatalogItemTypeInstanceScribePortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | **Int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | **String** | The Workflow Name to execute. | [optional] 

## Examples

- Prepare the resource
```powershell
$CatalogItemTypeInstanceScribe = Initialize-PSOpenAPIToolsCatalogItemTypeInstanceScribe  -Group null `
 -Cloud null `
 -Type ubuntu `
 -Name myinstance `
 -Config null `
 -Volumes null `
 -HostName myinstance `
 -Environment null `
 -Layout null `
 -Plan null `
 -Version null `
 -Evars null `
 -ServicePlanOptions null `
 -SecurityGroups null `
 -NetworkInterfaces null `
 -Labels null `
 -Tags null `
 -Metadata null `
 -Ports null `
 -TaskSetId null `
 -TaskSetName null
```

- Convert the resource to JSON
```powershell
$CatalogItemTypeInstanceScribe | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

