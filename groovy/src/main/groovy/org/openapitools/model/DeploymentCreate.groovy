package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class DeploymentCreate {
    /* Name, a unique identifier for the deployment */
    String name
    /* Description */
    String description
}
