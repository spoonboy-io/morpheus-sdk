package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class AwsResourcePoolConfig {
    /* Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) */
    String cidrBlock
    /* default or dedicated */
    String tenancy = TenancyEnum.DEFAULT
}
