package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ClusterLayoutCreateEnvironmentVariables {
    /* Environment variable name */
    String name
    /* Sets fixed value for variable */
    String value
    /* Can be used to enable / disable masking of variable */
    Boolean masked = false
    /* Can be used to enable / disable export of variable */
    Boolean export = false
}
