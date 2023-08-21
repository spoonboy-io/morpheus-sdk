# # InstancesConfigAWS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**no_agent** | **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**is_ec2** | **string** | Amazon Cloud Type | [optional] [default to 'false']
**availability_id** | **string** | Amazon Zone | [optional]
**security_id** | **string** | Security Group | [optional]
**public_ip_type** | **string** | Public IP | [optional]
**instance_profile** | **string** | IAM Profile | [optional]
**kms_key_id** | **string** | KMS Key ID | [optional]

[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)
