# CheckSqlConfig


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_host** | **str** | Hostname or IP address of the database | 
**db_port** | **str** | Database Port (defaults to default port of DB type selected) | 
**db_user** | **str** | Database username | 
**db_password** | **str** | Database password, (all check data is encrypted inside the database) | 
**db_name** | **str** | Database name you would like to connect to | 
**db_query** | **str** | Query to test | 
**db_password_hash** | **str** | Database password hash | [optional] 
**check_operator** | **str** | Can be set to &#x60;lt&#x60; (less than), &#x60;gt&#x60; (greater than), &#x60;equal&#x60; (Equal to) for comparison | [optional] 
**check_result** | **float** |  | [optional] 
**check_user** | **str** |  | [optional] 
**text_check_on** | **str** |  | [optional] 
**check_password** | **str** |  | [optional] 
**web_text_match** | **str** |  | [optional] 
**check_password_hash** | **str** |  | [optional] 
**tunnel_on** | **str** |  | [optional] 
**ssh_host** | **str** |  | [optional] 
**ssh_port** | **int** |  | [optional] 
**ssh_user** | **str** |  | [optional] 
**ssh_password** | **str** | Password for user, if not using key based authentication | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


