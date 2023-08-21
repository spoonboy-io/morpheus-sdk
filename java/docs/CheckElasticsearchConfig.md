

# CheckElasticsearchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**esHost** | **String** | Hostname or IP address of the Elasticsearch server | 
**esPort** | **String** | Port to connect to the HTTP service | 
**checkUser** | **String** |  |  [optional]
**textCheckOn** | **String** |  |  [optional]
**checkPassword** | **String** |  |  [optional]
**webTextMatch** | **String** |  |  [optional]
**checkPasswordHash** | **String** |  |  [optional]
**tunnelOn** | [**TunnelOnEnum**](#TunnelOnEnum) | Set to on to turn on tunneling |  [optional]
**sshHost** | **String** | Hostname or IP address of the proxy host |  [optional]
**sshPort** | **Long** | Port for SSH on the proxy host, defaults to 22 |  [optional]
**sshUser** | **String** | SSH user on the proxy host to login as |  [optional]
**sshPassword** | **String** | Password for user, if not using key based authentication |  [optional]



## Enum: TunnelOnEnum

Name | Value
---- | -----
ON | &quot;on&quot;
OFF | &quot;off&quot;



