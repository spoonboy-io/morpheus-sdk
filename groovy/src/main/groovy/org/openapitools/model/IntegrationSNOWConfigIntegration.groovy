package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IntegrationSNOWConfigIntegrationConfig;

@Canonical
class IntegrationSNOWConfigIntegration {
    /* Name, a unique identifier for the integration */
    String name
    /* Integration Type Code */
    String type
    /* Set `true` to enable integration */
    Boolean enabled
    /* Pass `false` to skip refresh.  By default, refresh is done on update, when it is supported by the integration type.  */
    Boolean refresh = true
    /* ServiceNow Host */
    String serviceUrl
    /* ServiceNow Username */
    String serviceUsername
    /* ServiceNow Password */
    String servicePassword
    
    IntegrationSNOWConfigIntegrationConfig config
}
