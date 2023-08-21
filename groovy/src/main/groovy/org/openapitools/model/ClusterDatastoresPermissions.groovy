package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterDatastoresPermissionsResourcePermissions;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class ClusterDatastoresPermissions {
    
    ClusterDatastoresPermissionsResourcePermissions resourcePermissions
    
    List<InlineResponse20040AppDeployInstance> tenants = new ArrayList<InlineResponse20040AppDeployInstance>()
}
