package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IntegrationGitRepoConfigIntegrationConfig;

@Canonical
class IntegrationGitRepoConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Git URL */
    String serviceUrl
    /* Username */
    String serviceUsername
    /* Password */
    String servicePassword
    /* Access Token */
    String serviceToken
    /* Key Pair ID */
    Long serviceKey
    
    IntegrationGitRepoConfigIntegrationConfig config
}
