# MorpheusApi.BackupTypeServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**locationType** | **String** |  | 
**name** | **String** | A name for the backup | 
**serverId** | **Number** | The ID of the server to backup | [optional] 
**backupType** | **String** | The backup type code, options vary by the type of cloud and source | 
**jobAction** | **String** | Create a new backup job, clone an existing job or add the new backup to an existing job | 
**jobId** | **Number** | The ID of the job to clone or add to. Only applies to jobAction &#x60;clone&#x60; and &#x60;addTo&#x60;. | [optional] 
**jobName** | **String** | Name for new job. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**jobSchedule** | **Number** | The ID of the execute schedule for new job. See Execute Schedules. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 
**retentionCount** | **Number** | Retention Count for new job. By default the backup settings value will be used. Only applies to jobAction &#x60;new&#x60; and &#x60;clone&#x60;. | [optional] 



## Enum: LocationTypeEnum


* `instance` (value: `"instance"`)





## Enum: BackupTypeEnum


* `alibabaSnapshot` (value: `"alibabaSnapshot"`)

* `amazonSnapshot` (value: `"amazonSnapshot"`)

* `avamarVMWareBackup` (value: `"avamarVMWareBackup"`)

* `azureSnapshot` (value: `"azureSnapshot"`)

* `bluemixSnapshot` (value: `"bluemixSnapshot"`)

* `commvaultFileBackup` (value: `"commvaultFileBackup"`)

* `commvaultOpenstackBackup` (value: `"commvaultOpenstackBackup"`)

* `commvaultVMWareBackup` (value: `"commvaultVMWareBackup"`)

* `digitaloceanSnapshot` (value: `"digitaloceanSnapshot"`)

* `directoryBackup` (value: `"directoryBackup"`)

* `esxiSnapshot` (value: `"esxiSnapshot"`)

* `fileBackup` (value: `"fileBackup"`)

* `fusionSnapshot` (value: `"fusionSnapshot"`)

* `googleSnapshot` (value: `"googleSnapshot"`)

* `huaweiSnapshot` (value: `"huaweiSnapshot"`)

* `hypervSnapshot` (value: `"hypervSnapshot"`)

* `kvm` (value: `"kvm"`)

* `lvmImage` (value: `"lvmImage"`)

* `lvmMigration` (value: `"lvmMigration"`)

* `lvmSnapshot` (value: `"lvmSnapshot"`)

* `MongoDB` (value: `"MongoDB"`)

* `morpheusAppliance` (value: `"morpheusAppliance"`)

* `morpheusContainerBackup` (value: `"morpheusContainerBackup"`)

* `morpheusStorageBackup` (value: `"morpheusStorageBackup"`)

* `morpheusVmBackup` (value: `"morpheusVmBackup"`)

* `MySQL` (value: `"MySQL"`)

* `nutanixSnapshot` (value: `"nutanixSnapshot"`)

* `openstackSnapshot` (value: `"openstackSnapshot"`)

* `opentelekomSnapshot` (value: `"opentelekomSnapshot"`)

* `oracleCloudSnapshot` (value: `"oracleCloudSnapshot"`)

* `Postgres` (value: `"Postgres"`)

* `powervcSnapshot` (value: `"powervcSnapshot"`)

* `rubrikVMWareBackup` (value: `"rubrikVMWareBackup"`)

* `scvmmSnapshot` (value: `"scvmmSnapshot"`)

* `softlayerSnapshot` (value: `"softlayerSnapshot"`)

* `SqlServer` (value: `"SqlServer"`)

* `tarDirectoryBackup` (value: `"tarDirectoryBackup"`)

* `upCloudSnapshot` (value: `"upCloudSnapshot"`)

* `vcdSnapshot` (value: `"vcdSnapshot"`)

* `veeamHypervBackup` (value: `"veeamHypervBackup"`)

* `veeamScvmmBackup` (value: `"veeamScvmmBackup"`)

* `veeamVcdBackup` (value: `"veeamVcdBackup"`)

* `veeamVMWareBackup` (value: `"veeamVMWareBackup"`)

* `virtustreamSnapshot` (value: `"virtustreamSnapshot"`)

* `vmwareSnapshot` (value: `"vmwareSnapshot"`)

* `winMigration` (value: `"winMigration"`)

* `xenSnapshot` (value: `"xenSnapshot"`)





## Enum: JobActionEnum


* `new` (value: `"new"`)

* `clone` (value: `"clone"`)

* `addTo` (value: `"addTo"`)




