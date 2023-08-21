

# AwsResourcePoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cidrBlock** | **String** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) |  [optional]
**tenancy** | [**TenancyEnum**](#TenancyEnum) | default or dedicated |  [optional]



## Enum: TenancyEnum

Name | Value
---- | -----
DEFAULT | &quot;default&quot;
DEDICATED | &quot;dedicated&quot;



