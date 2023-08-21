# MorpheusApi.ImageBuildCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the image build | [optional] 
**description** | **String** | A description for the image build | [optional] 
**type** | **String** | The image builder type. | [optional] 
**site** | [**ImageBuildCreateSite**](ImageBuildCreateSite.md) |  | [optional] 
**zone** | [**ImageBuildCreateZone**](ImageBuildCreateZone.md) |  | [optional] 
**config** | **Object** | A map of config values. This is the instance config that is used for provisioning. See Provisioning Types. | [optional] 
**bootScript** | [**ImageBuildCreateBootScript**](ImageBuildCreateBootScript.md) |  | [optional] 
**preseedScript** | [**ImageBuildCreatePreseedScript**](ImageBuildCreatePreseedScript.md) |  | [optional] 
**sshUsername** | **String** | SSH Username | [optional] 
**sshPassword** | **String** | SSH Password | [optional] 
**storageProvider** | **String** | Storage Provider | [optional] 
**isCloudInit** | **String** | Cloud Init | [optional] 
**buildOutputName** | **String** | Build Output Name | [optional] 
**conversionFormats** | **String** | Conversion Formats | [optional] 
**keepResults** | **Number** | Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality. | [optional] [default to 0]



## Enum: TypeEnum


* `vmware` (value: `"vmware"`)





## Enum: ConversionFormatsEnum


* `null` (value: `"null"`)

* `ovf` (value: `"ovf"`)

* `qcow2` (value: `"qcow2"`)

* `vhd` (value: `"vhd"`)




