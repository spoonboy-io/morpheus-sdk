# MorpheusApi.VirtualImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** |  | [optional] 
**name** | **String** |  | [optional] 
**description** | **String** |  | [optional] 
**labels** | **[String]** |  | [optional] 
**ownerId** | **Number** |  | [optional] 
**tenant** | [**InlineResponse20040AppDeployInstance**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**imageType** | **String** |  | [optional] 
**userUploaded** | **Boolean** |  | [optional] 
**userDefined** | **Boolean** |  | [optional] 
**systemImage** | **Boolean** |  | [optional] 
**isCloudInit** | **Boolean** |  | [optional] 
**sshUsername** | **String** |  | [optional] 
**sshPassword** | **String** |  | [optional] 
**sshPasswordHash** | **String** |  | [optional] 
**sshKey** | **String** |  | [optional] 
**osType** | [**VirtualImageOsType**](VirtualImageOsType.md) |  | [optional] 
**minRam** | **Number** |  | [optional] 
**minRamGB** | **Number** |  | [optional] 
**minDisk** | **String** |  | [optional] 
**minDiskGB** | **String** |  | [optional] 
**rawSize** | **Number** |  | [optional] 
**rawSizeGB** | **Number** |  | [optional] 
**trialVersion** | **Boolean** |  | [optional] 
**virtioSupported** | **Boolean** |  | [optional] 
**uefi** | **String** |  | [optional] 
**isAutoJoinDomain** | **Boolean** |  | [optional] 
**vmToolsInstalled** | **Boolean** |  | [optional] 
**installAgent** | **Boolean** |  | [optional] 
**isForceCustomization** | **Boolean** |  | [optional] 
**isSysprep** | **Boolean** |  | [optional] 
**fipsEnabled** | **Boolean** |  | [optional] 
**userData** | **String** |  | [optional] 
**consoleKeymap** | **String** |  | [optional] 
**storageProvider** | **String** |  | [optional] 
**externalId** | **String** |  | [optional] 
**visibility** | **String** |  | [optional] 
**accounts** | [**[InlineResponse20040AppDeployInstance]**](InlineResponse20040AppDeployInstance.md) |  | [optional] 
**config** | [**VirtualImageConfig**](VirtualImageConfig.md) |  | [optional] 
**volumes** | **[Object]** |  | [optional] 
**storageControllers** | **[Object]** |  | [optional] 
**networkInterfaces** | **[Object]** |  | [optional] 
**tags** | **[Object]** |  | [optional] 
**locations** | **[Object]** |  | [optional] 
**dateCreated** | **Date** |  | [optional] 
**lastUpdated** | **Date** |  | [optional] 
**status** | **String** |  | [optional] 

