# CheckElasticsearchConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**es_host** | **str** | Hostname or IP address of the Elasticsearch server | 
**es_port** | **str** | Port to connect to the HTTP service | 
**check_user** | **str** |  | [optional] 
**text_check_on** | **str** |  | [optional] 
**check_password** | **str** |  | [optional] 
**web_text_match** | **str** |  | [optional] 
**check_password_hash** | **str** |  | [optional] 
**tunnel_on** | **str** | Set to on to turn on tunneling | [optional]  if omitted the server will use the default value of "off"
**ssh_host** | **str** | Hostname or IP address of the proxy host | [optional] 
**ssh_port** | **int** | Port for SSH on the proxy host, defaults to 22 | [optional] 
**ssh_user** | **str** | SSH user on the proxy host to login as | [optional] 
**ssh_password** | **str** | Password for user, if not using key based authentication | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


