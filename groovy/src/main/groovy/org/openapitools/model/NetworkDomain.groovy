package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class NetworkDomain {
    
    Long id
    
    String name
    
    Boolean active
    
    String fqdn
    
    String description
    
    String visibility
    
    Boolean domainController
    
    Boolean publicZone
    
    String domainUsername
    
    String domainPassword
    
    String refType
    
    Long refId
    
    String refSource
    
    String internalId
    
    String ouPath
    
    String dcServer
    
    String zoneType
    
    String dnssec
    
    String domainSerial
    
    InlineResponse20040AppDeployInstance account
    
    InlineResponse20040AppDeployInstance owner
}
