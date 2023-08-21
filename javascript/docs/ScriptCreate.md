# MorpheusApi.ScriptCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Script name | 
**labels** | **[String]** | Array of label strings, can be used for filtering. | [optional] 
**category** | **String** | Script category | [optional] 
**scriptVersion** | **String** | Version of the script | [optional] [default to &#39;1&#39;]
**scriptPhase** | **String** | Phase for the script, provision, start, etc. | [optional] 
**scriptType** | **String** | Type for the script | [optional] [default to &#39;bash&#39;]
**script** | **String** | Script content, that is, the code itself. | [optional] 
**runAsUser** | **String** | Run as a specific user. | [optional] 
**sudoUser** | **Boolean** | Sudo, whether or not to run with sudo. | [optional] [default to false]



## Enum: ScriptTypeEnum


* `bash` (value: `"bash"`)

* `powershell` (value: `"powershell"`)




