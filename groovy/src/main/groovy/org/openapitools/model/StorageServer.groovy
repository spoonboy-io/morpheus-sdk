package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20079LoadBalancerMonitorLoadBalancerType;

@Canonical
class StorageServer {
    
    Long id
    
    String name
    
    InlineResponse20079LoadBalancerMonitorLoadBalancerType type
    
    String chassis
    
    String visibility
    
    String description
    
    String internalId
    
    String externalId
    
    String serviceUrl
    
    String serviceHost
    
    String servicePath
    
    String serviceToken
    
    String serviceTokenHash
    
    String serviceVersion
    
    String serviceUsername
    
    String servicePassword
    
    String servicePasswordHash
    
    String internalIp
    
    String externalIp
    
    String apiPort
    
    String adminPort
    
    Object config
    
    String refType
    
    Long refId
    
    String category
    
    String serverVendor
    
    String serverModel
    
    String serialNumber
    
    String status
    
    String statusMessage
    
    Date statusDate
    
    String errorMessage
    
    String maxStorage
    
    String usedStorage
    
    String diskCount
    
    Date dateCreated
    
    Date lastUpdated
    
    Boolean enabled
    
    List<Object> groups = new ArrayList<Object>()
    
    List<Object> hostGroups = new ArrayList<Object>()
    
    List<Object> hosts = new ArrayList<Object>()
    
    List<Object> tenants = new ArrayList<Object>()
    
    InlineResponse20040AppDeployInstance owner
    
    Object credential
}
