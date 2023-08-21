package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IntegrationDockerRepoConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Repository URL */
    String serviceUrl
    /* Username */
    String serviceUsername
    /* Password */
    String servicePassword
}
