

# BackupTypeProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**locationType** | [**LocationTypeEnum**](#LocationTypeEnum) |  | 
**name** | **String** | A name for the backup | 
**sourceProviderId** | **Long** | Source Object Store. The ID of the storage provider (bucket) to be backed up. |  [optional]
**storageProviderId** | **Long** | Target Object Store. The ID of the storage provider (bucket) the backup will be copied to. |  [optional]
**backupType** | [**BackupTypeEnum**](#BackupTypeEnum) | The backup type code, options vary by the type of cloud and source | 
**jobAction** | [**JobActionEnum**](#JobActionEnum) | Create a new backup job, clone an existing job or add the new backup to an existing job | 
**jobId** | **Long** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. |  [optional]
**jobName** | **String** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. |  [optional]
**jobSchedule** | **Long** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. |  [optional]
**retentionCount** | **Long** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. |  [optional]



## Enum: LocationTypeEnum

Name | Value
---- | -----
STORAGE | &quot;storage&quot;



## Enum: BackupTypeEnum

Name | Value
---- | -----
ALIBABASNAPSHOT | &quot;alibabaSnapshot&quot;
AMAZONSNAPSHOT | &quot;amazonSnapshot&quot;
AVAMARVMWAREBACKUP | &quot;avamarVMWareBackup&quot;
AZURESNAPSHOT | &quot;azureSnapshot&quot;
BLUEMIXSNAPSHOT | &quot;bluemixSnapshot&quot;
COMMVAULTFILEBACKUP | &quot;commvaultFileBackup&quot;
COMMVAULTOPENSTACKBACKUP | &quot;commvaultOpenstackBackup&quot;
COMMVAULTVMWAREBACKUP | &quot;commvaultVMWareBackup&quot;
DIGITALOCEANSNAPSHOT | &quot;digitaloceanSnapshot&quot;
DIRECTORYBACKUP | &quot;directoryBackup&quot;
ESXISNAPSHOT | &quot;esxiSnapshot&quot;
FILEBACKUP | &quot;fileBackup&quot;
FUSIONSNAPSHOT | &quot;fusionSnapshot&quot;
GOOGLESNAPSHOT | &quot;googleSnapshot&quot;
HUAWEISNAPSHOT | &quot;huaweiSnapshot&quot;
HYPERVSNAPSHOT | &quot;hypervSnapshot&quot;
KVM | &quot;kvm&quot;
LVMIMAGE | &quot;lvmImage&quot;
LVMMIGRATION | &quot;lvmMigration&quot;
LVMSNAPSHOT | &quot;lvmSnapshot&quot;
MONGODB | &quot;MongoDB&quot;
MORPHEUSAPPLIANCE | &quot;morpheusAppliance&quot;
MORPHEUSCONTAINERBACKUP | &quot;morpheusContainerBackup&quot;
MORPHEUSSTORAGEBACKUP | &quot;morpheusStorageBackup&quot;
MORPHEUSVMBACKUP | &quot;morpheusVmBackup&quot;
MYSQL | &quot;MySQL&quot;
NUTANIXSNAPSHOT | &quot;nutanixSnapshot&quot;
OPENSTACKSNAPSHOT | &quot;openstackSnapshot&quot;
OPENTELEKOMSNAPSHOT | &quot;opentelekomSnapshot&quot;
ORACLECLOUDSNAPSHOT | &quot;oracleCloudSnapshot&quot;
POSTGRES | &quot;Postgres&quot;
POWERVCSNAPSHOT | &quot;powervcSnapshot&quot;
RUBRIKVMWAREBACKUP | &quot;rubrikVMWareBackup&quot;
SCVMMSNAPSHOT | &quot;scvmmSnapshot&quot;
SOFTLAYERSNAPSHOT | &quot;softlayerSnapshot&quot;
SQLSERVER | &quot;SqlServer&quot;
TARDIRECTORYBACKUP | &quot;tarDirectoryBackup&quot;
UPCLOUDSNAPSHOT | &quot;upCloudSnapshot&quot;
VCDSNAPSHOT | &quot;vcdSnapshot&quot;
VEEAMHYPERVBACKUP | &quot;veeamHypervBackup&quot;
VEEAMSCVMMBACKUP | &quot;veeamScvmmBackup&quot;
VEEAMVCDBACKUP | &quot;veeamVcdBackup&quot;
VEEAMVMWAREBACKUP | &quot;veeamVMWareBackup&quot;
VIRTUSTREAMSNAPSHOT | &quot;virtustreamSnapshot&quot;
VMWARESNAPSHOT | &quot;vmwareSnapshot&quot;
WINMIGRATION | &quot;winMigration&quot;
XENSNAPSHOT | &quot;xenSnapshot&quot;



## Enum: JobActionEnum

Name | Value
---- | -----
NEW | &quot;new&quot;
CLONE | &quot;clone&quot;
ADDTO | &quot;addTo&quot;



