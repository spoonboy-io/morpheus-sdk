package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BlueprintKubernetesCreateKubernetesGit {
    /* Morpheus SCM Repository ID */
    Long repoId
    /* Path to kubernetes Files in the Repository */
    String path
    /* Morpheus SCM Integration ID */
    Long integrationId
    /* Branch Name */
    String branch
}
