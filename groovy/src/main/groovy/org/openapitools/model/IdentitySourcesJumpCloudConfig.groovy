package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.IdentitySourcesJumpCloudConfigConfig;
import org.openapitools.model.IdentitySourcesJumpCloudConfigRoleMappings;
import org.openapitools.model.IdentitySourcesLDAPConfigDefaultAccountRole;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class IdentitySourcesJumpCloudConfig {
    
    Long id
    
    String name
    
    String description
    
    String code
    
    String type
    
    Boolean active
    
    Boolean deleted
    
    Boolean autoSyncOnLogin
    
    Boolean externalLogin
    
    Boolean allowCustomMappings
    
    Boolean manualRoleAssignment
    
    InlineResponse20040AppDeployInstance account
    
    IdentitySourcesLDAPConfigDefaultAccountRole defaultAccountRole
    
    IdentitySourcesJumpCloudConfigConfig config
    
    List<IdentitySourcesJumpCloudConfigRoleMappings> roleMappings = new ArrayList<IdentitySourcesJumpCloudConfigRoleMappings>()
    
    String subdomain
    
    String loginURL
    
    Object providerSettings
    
    Date dateCreated
    
    Date lastUpdated
}
