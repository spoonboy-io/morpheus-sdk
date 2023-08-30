# SecurityScan


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | [optional] 
**security_package** | [**SecurityScanSecurityPackage**](SecurityScanSecurityPackage.md) |  | [optional] 
**status** | **str** |  | [optional] 
**scan_date** | **datetime** |  | [optional] 
**scan_duration** | **int** |  | [optional] 
**test_count** | **int** |  | [optional] 
**run_count** | **int** |  | [optional] 
**pass_count** | **int** |  | [optional] 
**fail_count** | **int** |  | [optional] 
**other_count** | **int** |  | [optional] 
**scan_score** | **float** |  | [optional] 
**external_id** | **str, none_type** |  | [optional] 
**created_by** | **str** |  | [optional] 
**updated_by** | **str** |  | [optional] 
**date_created** | **datetime** |  | [optional] 
**last_updated** | **datetime** |  | [optional] 
**results** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** | Results Summary (only returned when using query parameter results&#x3D;true) | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


