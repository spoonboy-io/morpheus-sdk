package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class AddAppInstanceRequest {
    /* The ID of the instance being added */
    Long instanceId
    /* The Name of the Tier */
    String tierName
}
