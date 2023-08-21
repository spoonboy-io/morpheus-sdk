package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.ClusterDatastoresPermissions;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ClusterDatastores {
    
    Long id
    
    String name
    
    String code
    
    String type
    
    String visibility
    
    Long storageSize
    
    Long freeSpace
    
    Boolean drsEnabled
    
    Boolean active
    
    Boolean allowWrite
    
    Boolean defaultStore
    
    Boolean online
    
    Boolean allowRead
    
    Boolean allowProvision
    
    String refType
    
    Long refId
    
    String externalId
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites zone
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites zonePool
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites owner
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
    
    List<Object> datastores = new ArrayList<Object>()
    
    ClusterDatastoresPermissions permissions
}
