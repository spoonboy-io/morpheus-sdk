package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.IdentitySourcesLDAPConfigDefaultAccountRole;

@Canonical
class IdentitySourcesLDAPConfigRoleMappings {
    
    String sourceRoleName
    
    String sourceRoleFqn
    
    IdentitySourcesLDAPConfigDefaultAccountRole mappedRole
}
