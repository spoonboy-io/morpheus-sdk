package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.ResourcePermissions;
import org.openapitools.model.ZoneDatastoreTenants;

@Canonical
class ZoneFolder {
    
    Long id
    
    String name
    
    InlineResponse20040AppDeployInstance zone
    
    InlineResponse20082LoadBalancerInstanceSslCert parent
    
    String type
    
    String externalId
    
    String visibility
    
    Boolean readOnly
    
    Boolean defaultFolder
    
    Boolean defaultStore
    
    Boolean active
    
    List<ZoneDatastoreTenants> tenants = new ArrayList<ZoneDatastoreTenants>()
    
    ResourcePermissions resourcePermissions
    
    Long depth
}
