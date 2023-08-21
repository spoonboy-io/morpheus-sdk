# ClusterDatastore
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**Name** | **String** |  | [optional] 
**Code** | **String** |  | [optional] 
**Type** | **String** |  | [optional] 
**Visibility** | **String** |  | [optional] 
**StorageSize** | **Int64** |  | [optional] 
**FreeSpace** | **Int64** |  | [optional] 
**DrsEnabled** | **Boolean** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**AllowWrite** | **Boolean** |  | [optional] 
**DefaultStore** | **Boolean** |  | [optional] 
**Online** | **Boolean** |  | [optional] 
**AllowRead** | **Boolean** |  | [optional] 
**AllowProvision** | **Boolean** |  | [optional] 
**RefType** | **String** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**Zone** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**ZonePool** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Owner** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Tenants** | [**InlineResponse20040AppDeployInstance[]**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**Permissions** | [**ClusterDatastorePermissions**](ClusterDatastorePermissions.md) |  | [optional] 
**Datastores** | [**SystemCollectionsHashtable[]**](SystemCollectionsHashtable.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterDatastore = Initialize-PSOpenAPIToolsClusterDatastore  -Id null `
 -Name null `
 -Code null `
 -Type null `
 -Visibility null `
 -StorageSize null `
 -FreeSpace null `
 -DrsEnabled null `
 -Active null `
 -AllowWrite null `
 -DefaultStore null `
 -Online null `
 -AllowRead null `
 -AllowProvision null `
 -RefType null `
 -RefId null `
 -ExternalId null `
 -Zone null `
 -ZonePool null `
 -Owner null `
 -Tenants null `
 -Permissions null `
 -Datastores null
```

- Convert the resource to JSON
```powershell
$ClusterDatastore | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

