package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IntegrationAnsibleConfigIntegrationConfig;

@Canonical
class IntegrationAnsibleConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Set `true` to enable integration */
    Boolean enabled
    /* Pass `false` to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  */
    Boolean refresh = true
    /* Ansible Git URL */
    String serviceUrl
    /* Git Username */
    String serviceUsername
    /* Git Password or Token depending on the Git host */
    String servicePassword
    /* Git Token */
    String serviceToken
    /* Keypair ID */
    Long serviceKey
    
    IntegrationAnsibleConfigIntegrationConfig config
}
