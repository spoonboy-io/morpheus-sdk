package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfobjectobjectobject;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.ResourcePermissions;

@Canonical
class ZoneResourcePool {
    
    Long id
    
    String description
    
    InlineResponse20040AppDeployInstance zone
    
    InlineResponse20082LoadBalancerInstanceSslCert parent
    
    String type
    
    String externalId
    
    String regionCode
    
    String iacId
    
    String visibility
    
    Boolean readOnly
    
    Boolean defaultPool
    
    Boolean active
    
    String status
    
    Boolean inventory
    
    AnyOfobjectobjectobject config = null
    
    String name
    
    String displayName
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    ResourcePermissions resourcePermission
    
    Long depth
}
