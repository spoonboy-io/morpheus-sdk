package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IntegrationSaltMasterConfigIntegrationConfig;

@Canonical
class IntegrationSaltMasterConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Topology */
    String serviceMode = "single"
    /* Salt Master */
    String serviceUrl
    /* Salt Syndic */
    String secondary
    /* SSH Port */
    Integer servicePort = 22
    /* Username */
    String serviceUsername
    /* Password */
    String servicePassword
    /* Master Key Pair */
    String serviceKey
    /* Signing Key */
    String authKey
    /* Working Directory */
    String servicePath
    /* Salt Version */
    String serviceVersion
    /* Salt Version (Windows) */
    String serviceWindowsVersion
    /* Repo URL */
    String repoUrl
    /* Minion Config */
    String serviceConfig
    /* Post Provision Commands */
    String serviceCommand
    
    IntegrationSaltMasterConfigIntegrationConfig config
}
