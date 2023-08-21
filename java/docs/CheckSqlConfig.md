

# CheckSqlConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dbHost** | **String** | Hostname or IP address of the database | 
**dbPort** | **String** | Database Port (defaults to default port of DB type selected) | 
**dbUser** | **String** | Database username | 
**dbPassword** | **String** | Database password, (all check data is encrypted inside the database) | 
**dbPasswordHash** | **String** | Database password hash |  [optional]
**dbName** | **String** | Database name you would like to connect to | 
**dbQuery** | **String** | Query to test | 
**checkOperator** | [**CheckOperatorEnum**](#CheckOperatorEnum) | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison |  [optional]
**checkResult** | **BigDecimal** |  |  [optional]
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



## Enum: CheckOperatorEnum

Name | Value
---- | -----
LT | &quot;lt&quot;
GT | &quot;gt&quot;
EQUAL | &quot;equal&quot;



## Enum: TunnelOnEnum

Name | Value
---- | -----
ON | &quot;on&quot;
OFF | &quot;off&quot;



