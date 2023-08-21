package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.IdentitySourcesAzureADConfigConfig;
import org.openapitools.model.IdentitySourcesLDAPConfigDefaultAccountRole;
import org.openapitools.model.IdentitySourcesSAMLConfigProviderSettings;
import org.openapitools.model.IdentitySourcesSAMLConfigRoleMappings;
import org.openapitools.model.InlineResponse20040AppDeployInstance;

@Canonical
class IdentitySourcesAzureADConfig {
    
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
    
    IdentitySourcesAzureADConfigConfig config
    
    List<IdentitySourcesSAMLConfigRoleMappings> roleMappings = new ArrayList<IdentitySourcesSAMLConfigRoleMappings>()
    
    String subdomain
    
    String loginURL
    
    IdentitySourcesSAMLConfigProviderSettings providerSettings
    
    Date dateCreated
    
    Date lastUpdated
}
