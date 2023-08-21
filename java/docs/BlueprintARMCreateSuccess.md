

# BlueprintARMCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | A name for the blueprint |  [optional]
**image** | **String** | Path to display image. Defaults to an internal Morpheus image. |  [optional]
**type** | [**TypeEnum**](#TypeEnum) | Blueprint Type |  [optional]
**arm** | [**BlueprintARMCreateArm**](BlueprintARMCreateArm.md) |  |  [optional]
**visibility** | [**VisibilityEnum**](#VisibilityEnum) | Private or Public Access |  [optional]
**resourcePermission** | **Object** | Resource Permission Block |  [optional]
**owner** | **Object** | Owner |  [optional]
**tenant** | **Object** | Tenant |  [optional]



## Enum: TypeEnum

Name | Value
---- | -----
ARM | &quot;arm&quot;



## Enum: VisibilityEnum

Name | Value
---- | -----
PRIVATE | &quot;private&quot;
PUBLIC | &quot;public&quot;



