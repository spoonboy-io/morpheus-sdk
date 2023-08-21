package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IntegrationGitHubConfigIntegrationConfig;

@Canonical
class IntegrationGitHubConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Username */
    String serviceUsername
    /* Password */
    String servicePassword
    /* Access Token */
    String serviceToken
    /* Key Pair ID */
    Long serviceKey
    
    IntegrationGitHubConfigIntegrationConfig config
}
