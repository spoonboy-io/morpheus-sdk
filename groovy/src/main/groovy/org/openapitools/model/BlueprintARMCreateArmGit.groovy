package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BlueprintARMCreateArmGit {
    /* Morpheus SCM Repository ID */
    Long repoId
    /* Path to ARM Files in the Repository */
    String path
    /* Morpheus SCM Integration ID */
    Long integrationId
    /* Branch Name */
    String branch
}
