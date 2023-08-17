

# BlueprintARMCreateArm


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**configType** | [**ConfigTypeEnum**](#ConfigTypeEnum) | Configuration Type |  |
|**json** | **String** | ARM Template in JSON |  [optional] |
|**yaml** | **String** | ARM Template in YAML |  [optional] |
|**git** | [**BlueprintARMCreateArmGit**](BlueprintARMCreateArmGit.md) |  |  [optional] |
|**osType** | [**OsTypeEnum**](#OsTypeEnum) | OS Type |  [optional] |
|**installAgent** | [**BlueprintARMCreateArmInstallAgent**](BlueprintARMCreateArmInstallAgent.md) |  |  [optional] |
|**cloudInitEnabled** | [**BlueprintARMCreateArmCloudInitEnabled**](BlueprintARMCreateArmCloudInitEnabled.md) |  |  [optional] |



## Enum: ConfigTypeEnum

| Name | Value |
|---- | -----|
| JSON | &quot;json&quot; |
| YAML | &quot;yaml&quot; |
| GIT | &quot;git&quot; |



## Enum: OsTypeEnum

| Name | Value |
|---- | -----|
| LINUX | &quot;linux&quot; |
| WINDOWS | &quot;windows&quot; |



