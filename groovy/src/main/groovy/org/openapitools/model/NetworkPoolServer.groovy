package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.Creds2;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.NetworkPoolServerAccount;
import org.openapitools.model.NetworkPoolServerIntegration;
import org.openapitools.model.NetworkPoolServerType;

@Canonical
class NetworkPoolServer {
    /* Network Pool Server ID */
    Long id
    
    NetworkPoolServerType type
    /* Name */
    String name
    
    Boolean enabled
    /* Service URL */
    String serviceUrl
    /* Service Host */
    String serviceHost
    /* Service Port */
    Integer servicePort
    /* Service Mode */
    String serviceMode
    /* Service Username */
    String serviceUsername
    /* Service Password */
    String servicePassword
    
    String servicePasswordHash
    /* Throttle Rate */
    Long serviceThrottleRate = 0l
    /* Disable SSL SNI Verification */
    Boolean ignoreSsl = true
    /* Status */
    String status
    /* Status Message */
    String statusMessage
    
    Date statusDate
    /* Config object varies with pool server type. */
    Object config
    /* Network Filter */
    String networkFilter
    /* Zone Filter */
    String zoneFilter
    /* Tenant Match */
    String tenantMatch
    
    Date dateCreated
    
    Date lastUpdated
    
    NetworkPoolServerAccount account
    
    NetworkPoolServerIntegration integration
    
    List<InlineResponse20040AppDeployInstance> pools = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    Creds2 credential
}
