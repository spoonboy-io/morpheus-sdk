package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class DeploymentVersionCreate {
    /* Version number (userVersion), a unique version identifier for the deployment version. */
    String version
    /* Alias for version */
    String userVersion
    /* Deploy Type, eg. file, git, fetch */
    String deployType
    
    String gitUrl
    
    String gitRef
    
    String fetchUrl
}
