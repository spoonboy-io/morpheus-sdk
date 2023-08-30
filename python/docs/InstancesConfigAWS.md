# InstancesConfigAWS


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_agent** | **bool, none_type** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional]  if omitted the server will use the default value of False
**is_ec2** | **str** | Amazon Cloud Type | [optional]  if omitted the server will use the default value of "false"
**availability_id** | **str** | Amazon Zone | [optional] 
**security_id** | **str** | Security Group | [optional] 
**public_ip_type** | **str** | Public IP | [optional] 
**instance_profile** | **str** | IAM Profile | [optional] 
**kms_key_id** | **str** | KMS Key ID | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


