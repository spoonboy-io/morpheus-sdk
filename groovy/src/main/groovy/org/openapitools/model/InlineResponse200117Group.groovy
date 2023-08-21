package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.NetworkServerGroupMember;
import org.openapitools.model.Permissions;
import org.openapitools.model.Tag;

@Canonical
class InlineResponse200117Group {
    
    Long id
    
    String name
    
    String description
    
    String internalId
    
    String externalId
    
    String visibility
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites account
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites owner
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites networkServer
    
    Permissions permissions
    
    List<Tag> tags = new ArrayList<Tag>()
    
    List<NetworkServerGroupMember> members = new ArrayList<NetworkServerGroupMember>()
}
