# # CheckSqlConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_host** | **string** | Hostname or IP address of the database |
**db_port** | **string** | Database Port (defaults to default port of DB type selected) |
**db_user** | **string** | Database username |
**db_password** | **string** | Database password, (all check data is encrypted inside the database) |
**db_password_hash** | **string** | Database password hash | [optional]
**db_name** | **string** | Database name you would like to connect to |
**db_query** | **string** | Query to test |
**check_operator** | **string** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional]
**check_result** | **float** |  | [optional]
**check_user** | **string** |  | [optional]
**text_check_on** | **string** |  | [optional]
**check_password** | **string** |  | [optional]
**web_text_match** | **string** |  | [optional]
**check_password_hash** | **string** |  | [optional]
**tunnel_on** | **string** | Set to on to turn on tunneling | [optional] [default to 'off']
**ssh_host** | **string** | Hostname or IP address of the proxy host | [optional]
**ssh_port** | **int** | Port for SSH on the proxy host, defaults to 22 | [optional]
**ssh_user** | **string** | SSH user on the proxy host to login as | [optional]
**ssh_password** | **string** | Password for user, if not using key based authentication | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
