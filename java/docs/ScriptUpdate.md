

# ScriptUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Script name |  [optional]
**labels** | **List&lt;String&gt;** | Array of label strings, can be used for filtering. |  [optional]
**category** | **String** | Script category |  [optional]
**scriptVersion** | **String** | Version of the script |  [optional]
**scriptPhase** | **String** | Phase for the script, provision, start, etc. |  [optional]
**scriptType** | [**ScriptTypeEnum**](#ScriptTypeEnum) | Type for the script |  [optional]
**script** | **String** | Script content, that is, the code itself. |  [optional]
**runAsUser** | **String** | Run as a specific user. |  [optional]
**sudoUser** | **Boolean** | Sudo, whether or not to run with sudo. |  [optional]



## Enum: ScriptTypeEnum

Name | Value
---- | -----
BASH | &quot;bash&quot;
POWERSHELL | &quot;powershell&quot;



