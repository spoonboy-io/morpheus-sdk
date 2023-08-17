# MorpheusApi.BlueprintARMCreateArm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**configType** | **String** | Configuration Type | 
**json** | **String** | ARM Template in JSON | [optional] 
**yaml** | **String** | ARM Template in YAML | [optional] 
**git** | [**BlueprintARMCreateArmGit**](BlueprintARMCreateArmGit.md) |  | [optional] 
**osType** | **String** | OS Type | [optional] 
**installAgent** | [**BlueprintARMCreateArmInstallAgent**](BlueprintARMCreateArmInstallAgent.md) |  | [optional] 
**cloudInitEnabled** | [**BlueprintARMCreateArmCloudInitEnabled**](BlueprintARMCreateArmCloudInitEnabled.md) |  | [optional] 



## Enum: ConfigTypeEnum


* `json` (value: `"json"`)

* `yaml` (value: `"yaml"`)

* `git` (value: `"git"`)





## Enum: OsTypeEnum


* `linux` (value: `"linux"`)

* `windows` (value: `"windows"`)




