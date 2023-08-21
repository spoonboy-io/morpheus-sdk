# MorpheusApi.CheckWebConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**webMethod** | **String** | HTTP method to use for testing | 
**webUrl** | **String** | Web URL you wish to use to run a check on | 
**ignoreSSL** | **Boolean** | Ignore SSL Errors | [optional] [default to false]
**checkUser** | **String** | If you want to use HTTP Basic Authentication, populate this field with the username | [optional] 
**checkPassword** | **String** | If you want to use HTTP basic Authentication, populate this field with the password | [optional] 
**textCheckOn** | **String** | Set value to &#x60;on&#x60; if you want to turn on text matching | [optional] 
**webTextMatch** | **String** | Set the string you want to look for in the page source | [optional] 
**tunnelOn** | **String** | Set to on to turn on tunneling | [optional] 
**sshHost** | **String** | Hostname or IP address of the proxy host | [optional] 
**sshPort** | **Number** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**sshUser** | **String** | SSH user on the proxy host to login as | [optional] 
**sshPassword** | **String** | Password for user, if not using key based authentication | [optional] 



## Enum: WebMethodEnum


* `GET` (value: `"GET"`)

* `POST` (value: `"POST"`)





## Enum: TunnelOnEnum


* `on` (value: `"on"`)

* `off` (value: `"off"`)




