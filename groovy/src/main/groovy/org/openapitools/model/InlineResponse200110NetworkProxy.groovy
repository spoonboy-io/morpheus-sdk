package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class InlineResponse200110NetworkProxy {
    
    Long id
    
    String name
    
    String proxyHost
    
    Long proxyPort
    
    String proxyUser
    
    String proxyPassword
    
    String proxyDomain
    
    String proxyWorkstation
    
    String visibility
    
    InlineResponse20040AppDeployInstance account
    
    InlineResponse20040AppDeployInstance owner
}
