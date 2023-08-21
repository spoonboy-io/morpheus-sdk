package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IntegrationGitRepoConfigIntegrationConfig {
    /* Default Branch */
    String defaultBranch
    /* Enable Git Repository Caching */
    Boolean cacheEnabled
}
