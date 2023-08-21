package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterDatastoresPermissionsResourcePermissions;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.SecurityGroupLocations;
import org.openapitools.model.SecurityGroupRules;
import org.openapitools.model.SecurityGroupTenants;

@Canonical
class SecurityGroup {
    
    Long id
    
    String name
    
    String description
    
    Long accountId
    
    String groupSource
    
    String externalId
    
    String enabled
    
    String syncSource
    
    String visibility
    
    Boolean active
    
    InlineResponse20082LoadBalancerInstanceSslCert zone
    
    List<SecurityGroupLocations> locations = new ArrayList<SecurityGroupLocations>()
    
    List<SecurityGroupRules> rules = new ArrayList<SecurityGroupRules>()
    
    List<SecurityGroupTenants> tenants = new ArrayList<SecurityGroupTenants>()
    
    ClusterDatastoresPermissionsResourcePermissions resourcePermission
}
