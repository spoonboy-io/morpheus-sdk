

# ImageBuildCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the image build |  [optional]
**description** | **String** | A description for the image build |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | The image builder type. |  [optional]
**site** | [**ImageBuildCreateSite**](ImageBuildCreateSite.md) |  |  [optional]
**zone** | [**ImageBuildCreateZone**](ImageBuildCreateZone.md) |  |  [optional]
**config** | **Object** | A map of config values. This is the instance config that is used for provisioning. See Provisioning Types. |  [optional]
**bootScript** | [**ImageBuildCreateBootScript**](ImageBuildCreateBootScript.md) |  |  [optional]
**preseedScript** | [**ImageBuildCreatePreseedScript**](ImageBuildCreatePreseedScript.md) |  |  [optional]
**sshUsername** | **String** | SSH Username |  [optional]
**sshPassword** | **String** | SSH Password |  [optional]
**storageProvider** | **String** | Storage Provider |  [optional]
**isCloudInit** | **String** | Cloud Init |  [optional]
**buildOutputName** | **String** | Build Output Name |  [optional]
**conversionFormats** | [**ConversionFormatsEnum**](#ConversionFormatsEnum) | Conversion Formats |  [optional]
**keepResults** | **Long** | Keep Results - Keep only the most recent builds. Older executions will be deleted along with their associated Virtual Images. The value 0 disables this functionality. |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
VMWARE | &quot;vmware&quot;



## Enum: ConversionFormatsEnum

Name | Value
---- | -----
NULL | &quot;null&quot;
OVF | &quot;ovf&quot;
QCOW2 | &quot;qcow2&quot;
VHD | &quot;vhd&quot;



