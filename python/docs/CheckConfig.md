# CheckConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ignore_ssl** | **bool** | Ignore SSL Errors | [optional]  if omitted the server will use the default value of False
**check_user** | **str** |  | [optional] 
**check_password** | **str** |  | [optional] 
**text_check_on** | **str** |  | [optional] 
**web_text_match** | **str** |  | [optional] 
**tunnel_on** | **str** | Set to on to turn on tunneling | [optional]  if omitted the server will use the default value of "off"
**ssh_host** | **str** | Hostname or IP address of the proxy host | [optional] 
**ssh_port** | **int** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**ssh_user** | **str** | SSH user on the proxy host to login as | [optional] 
**ssh_password** | **str** | Password for user, if not using key based authentication | [optional] 
**db_password_hash** | **str** | Database password hash | [optional] 
**check_operator** | **str** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional] 
**check_result** | **float** |  | [optional] 
**check_password_hash** | **str** |  | [optional] 
**external_id** | **str, none_type** |  | [optional] 
**web_method** | **str** | HTTP method to use for testing | [optional] 
**web_url** | **str** | Web URL you wish to use to run a check on | [optional] 
**db_host** | **str** | Hostname or IP address of the database | [optional] 
**db_port** | **str** | Database Port (defaults to default port of DB type selected) | [optional] 
**db_user** | **str** | Database username | [optional] 
**db_password** | **str** | Database password, (all check data is encrypted inside the database) | [optional] 
**db_name** | **str** | Database name you would like to connect to | [optional] 
**db_query** | **str** | Query to test | [optional] 
**es_host** | **str** | Hostname or IP address of the Elasticsearch server | [optional] 
**es_port** | **str** | Port to connect to the HTTP service | [optional] 
**host** | **str** | Hostname or IP address of the socket server | [optional] 
**port** | **str** | TCP port | [optional] 
**send** | **str** | Connection string you might want to send to the service | [optional] 
**response_match** | **str** | Response from the service to match against | [optional] 
**container_name** | **str** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


