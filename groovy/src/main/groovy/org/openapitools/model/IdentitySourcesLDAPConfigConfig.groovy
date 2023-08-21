package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class IdentitySourcesLDAPConfigConfig {
    
    String url
    
    String bindingUsername
    
    String bindingPassword
    
    String userFqnExpression
    
    String requiredRoleFqn
    
    String usernameAttribute
    
    String commonNameAttribute
    
    String firstNameAttribute
    
    String lastNameAttribute
    
    String emailAttribute
    
    String uniqueMemberAttribute
    
    String memberOfAttribute
    
    String bindingPasswordHash
}
