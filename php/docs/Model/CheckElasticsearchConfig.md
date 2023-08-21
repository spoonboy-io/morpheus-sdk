# # CheckElasticsearchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**es_host** | **string** | Hostname or IP address of the Elasticsearch server |
**es_port** | **string** | Port to connect to the HTTP service |
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
