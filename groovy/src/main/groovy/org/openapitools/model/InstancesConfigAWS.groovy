package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstancesConfigAWS {
    /* Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. */
    Boolean noAgent = false
    /* Amazon Cloud Type */
    String isEC2 = "false"
    /* Amazon Zone */
    String availabilityId
    /* Security Group */
    String securityId
    /* Public IP */
    String publicIpType
    /* IAM Profile */
    String instanceProfile
    /* KMS Key ID */
    String kmsKeyId
}
