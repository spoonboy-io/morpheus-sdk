package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.ResourcePermissions;
import org.openapitools.model.ZoneDatastoreTenants;

@Canonical
class ZoneDatastore {
    
    Long id
    
    String name
    
    InlineResponse20040AppDeployInstance zone
    
    String type
    
    Long freeSpace
    
    Boolean online
    
    Boolean active
    
    String visibility
    
    List<ZoneDatastoreTenants> tenants = new ArrayList<ZoneDatastoreTenants>()
    
    ResourcePermissions resourcePermission
}
