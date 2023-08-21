package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.VdiPoolConfig;
import org.openapitools.model.VdiPoolOwner;

@Canonical
class VdiPool {
    
    Long id
    
    String name
    
    String description
    
    Long minIdle
    
    Long maxIdle
    
    Long initialPoolSize
    
    Long maxPoolSize
    
    Long allocationTimeoutMinutes
    
    Boolean persistentUser
    
    Boolean recyclable
    
    Boolean enabled
    
    Boolean autoCreateLocalUserOnReservation
    
    Boolean allowHypervisorConsole
    
    Boolean allowCopy
    
    Boolean allowPrinter
    
    Boolean allowFileshare
    
    String guestConsoleJumpHost
    
    String guestConsoleJumpPort
    
    String guestConsoleJumpUsername
    
    String guestConsoleJumpPassword
    
    String guestConsoleJumpKeypair
    
    InlineResponse20082LoadBalancerInstanceSslCert gateway
    
    String iconPath
    
    String logo
    
    List<InlineResponse20040AppDeployInstance> apps = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    VdiPoolOwner owner
    
    VdiPoolConfig config
    
    InlineResponse20082LoadBalancerInstanceSslCert group
    
    InlineResponse20082LoadBalancerInstanceSslCert cloud
    
    Long usedCount
    
    Long reservedCount
    
    Long preparingCount
    
    Long idleCount
    
    String status
    
    Date dateCreated
    
    Date lastUpdated
}
