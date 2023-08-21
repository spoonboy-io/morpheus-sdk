# MorpheusApi.CheckVmConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**containerName** | **String** |  | 
**externalId** | **String** |  | [optional] 
**checkUser** | **String** |  | [optional] 
**textCheckOn** | **String** |  | [optional] 
**checkPassword** | **String** |  | [optional] 
**webTextMatch** | **String** |  | [optional] 
**checkPasswordHash** | **String** |  | [optional] 
**tunnelOn** | **String** | Set to on to turn on tunneling | [optional] [default to &#39;off&#39;]
**sshHost** | **String** | Hostname or IP address of the proxy host | [optional] 
**sshPort** | **Number** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**sshUser** | **String** | SSH user on the proxy host to login as | [optional] 
**sshPassword** | **String** | Password for user, if not using key based authentication | [optional] 



## Enum: TunnelOnEnum


* `on` (value: `"on"`)

* `off` (value: `"off"`)




