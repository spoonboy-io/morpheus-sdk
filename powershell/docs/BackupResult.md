# BackupResult
## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **Int64** | Backup Result ID | [optional] 
**Backup** | [**BackupJobBackups**](BackupJobBackups.md) |  | [optional] 
**BackupSetId** | **String** |  | [optional] 
**InstanceId** | **Int64** |  | [optional] 
**ContainerId** | **Int64** |  | [optional] 
**ServerId** | **Int64** |  | [optional] 
**Status** | **String** |  | [optional] 
**ErrorMessage** | **String** |  | [optional] 
**StartDate** | **System.DateTime** |  | [optional] 
**EndDate** | **System.DateTime** |  | [optional] 
**DurationMillis** | **Int64** |  | [optional] 
**SizeInBytes** | **Int64** |  | [optional] 
**SizeInMb** | **Int64** |  | [optional] 
**VolumePath** | **String** |  | [optional] 
**ResultArchive** | **String** |  | [optional] 
**ResultPath** | **String** |  | [optional] 
**ExternalId** | **String** |  | [optional] 
**SnapshotId** | **String** |  | [optional] 
**SnapshotExternalId** | **String** |  | [optional] 
**CreatedBy** | [**InlineResponse200108NetworkFloatingIpCreatedBy**](InlineResponse200108NetworkFloatingIpCreatedBy.md) |  | [optional] 
**DateCreated** | **System.DateTime** | Date Created | [optional] 
**LastUpdated** | **System.DateTime** | Last Updated | [optional] 

## Examples

- Prepare the resource
```powershell
$BackupResult = Initialize-PSOpenAPIToolsBackupResult  -Id null `
 -Backup null `
 -BackupSetId null `
 -InstanceId null `
 -ContainerId null `
 -ServerId null `
 -Status null `
 -ErrorMessage null `
 -StartDate null `
 -EndDate null `
 -DurationMillis null `
 -SizeInBytes null `
 -SizeInMb null `
 -VolumePath null `
 -ResultArchive null `
 -ResultPath null `
 -ExternalId null `
 -SnapshotId null `
 -SnapshotExternalId null `
 -CreatedBy null `
 -DateCreated null `
 -LastUpdated null
```

- Convert the resource to JSON
```powershell
$BackupResult | ConvertTo-JSON
```

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

