package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.Creds;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20094Network;
import org.openapitools.model.Owner;
import org.openapitools.model.ResourcePermissions;

@Canonical
class LoadBalancer {
    
    Long id
    
    String uuid
    
    String name
    
    Long accountId
    
    InlineResponse20040AppDeployInstance cloud
    
    InlineResponse20094Network type
    
    Owner owner
    
    String visibility
    
    String description
    
    String host
    
    Long port
    
    String username
    
    String ip
    
    String internalIp
    
    String externalIp
    
    String apiPort
    
    String adminPort
    
    Boolean sslEnabled
    
    String sslCert
    
    Object config
    
    Date dateCreated
    
    Date lastUpdated
    
    Creds credential
    
    List<Object> tenants = new ArrayList<Object>()
    
    ResourcePermissions resourcePermission
}
