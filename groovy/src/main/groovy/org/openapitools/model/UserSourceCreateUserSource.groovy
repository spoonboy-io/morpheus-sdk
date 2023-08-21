package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.HashMap;
import java.util.List;
import org.openapitools.model.OneOfarraymap;
import org.openapitools.model.OneOfuserSourceCreateLDAPuserSourceCreateJumpClouduserSourceCreateActiveDirectoryuserSourceCreateOktauserSourceCreateOneLoginuserSourceCreateSamluserSourceCreateCustomExternaluserSourceCreateCustomApi;
import org.openapitools.model.UserSourceCreateUserSourceDefaultAccountRole;

@Canonical
class UserSourceCreateUserSource {
    /* A name for the Identity Source */
    String name
    /* IDM type code */
    String type
    /* description */
    String description
    
    UserSourceCreateUserSourceDefaultAccountRole defaultAccountRole
    
    OneOfarraymap roleMappings = null
    /* Map of Morpheus '`Role ID`':'`Role Name`'.  */
    Map<String, String> roleMappingNames = new HashMap<String, String>()
    /* Enable Role Mapping Permission */
    Boolean allowCustomMappings
    /* Manual Role Assignment */
    Boolean manualRoleAssignment
    
    OneOfuserSourceCreateLDAPuserSourceCreateJumpClouduserSourceCreateActiveDirectoryuserSourceCreateOktauserSourceCreateOneLoginuserSourceCreateSamluserSourceCreateCustomExternaluserSourceCreateCustomApi config = null
}
