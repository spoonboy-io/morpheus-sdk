# ClusterVolumes
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** |  | [optional] 
**InternalId** | **String** |  | [optional] 
**DisplayOrder** | **Int64** |  | [optional] 
**Active** | **Boolean** |  | [optional] 
**UsedStorage** | **Int64** |  | [optional] 
**ProvisionType** | **String** |  | [optional] 
**Resizeable** | **Boolean** |  | [optional] 
**Online** | **Boolean** |  | [optional] 
**DeviceDisplayName** | **String** |  | [optional] 
**RefType** | **String** |  | [optional] 
**Name** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**DatastoreOption** | **String** |  | [optional] 
**ClaimName** | **String** |  | [optional] 
**VolumeType** | **String** |  | [optional] 
**DeviceName** | **String** |  | [optional] 
**Removable** | **Boolean** |  | [optional] 
**PoolName** | **String** |  | [optional] 
**ReadOnly** | **Boolean** |  | [optional] 
**SourceId** | **String** |  | [optional] 
**ZoneId** | **Int64** |  | [optional] 
**RootVolume** | **Boolean** |  | [optional] 
**RefId** | **Int64** |  | [optional] 
**Category** | **String** |  | [optional] 
**Status** | **String** |  | [optional] 
**RawData** | **String** |  | [optional] 
**MaxStorage** | **Int64** |  | [optional] 
**Account** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Type** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 
**Datastore** | [**ApiBlueprintsIdUpdatePermissionsResourcePermissionSites**](ApiBlueprintsIdUpdatePermissionsResourcePermissionSites.md) |  | [optional] 

## Examples

- Prepare the resource
```powershell
$ClusterVolumes = Initialize-PSOpenAPIToolsClusterVolumes  -Id null `
 -InternalId null `
 -DisplayOrder null `
 -Active null `
 -UsedStorage null `
 -ProvisionType null `
 -Resizeable null `
 -Online null `
 -DeviceDisplayName null `
 -RefType null `
 -Name null `
 -ExternalId null `
 -DatastoreOption null `
 -ClaimName null `
 -VolumeType null `
 -DeviceName null `
 -Removable null `
 -PoolName null `
 -ReadOnly null `
 -SourceId null `
 -ZoneId null `
 -RootVolume null `
 -RefId null `
 -Category null `
 -Status null `
 -RawData null `
 -MaxStorage null `
 -Account null `
 -Type null `
 -Datastore null
```

- Convert the resource to JSON
```powershell
$ClusterVolumes | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

