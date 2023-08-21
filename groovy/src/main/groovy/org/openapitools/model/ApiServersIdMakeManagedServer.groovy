package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersIdInstallAgentServerServerOs;
import org.openapitools.model.ApiServersIdMakeManagedServerAccount;
import org.openapitools.model.ApiServersIdMakeManagedServerConfig;
import org.openapitools.model.ApiServersIdMakeManagedServerPlan;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;

@Canonical
class ApiServersIdMakeManagedServer {
    /* SSH username to use when provisioning */
    String sshUsername
    /* SSH password to use, if not specified the account public key can be used */
    String sshPassword
    
    ApiServersIdInstallAgentServerServerOs serverOs
    
    ApiServersIdMakeManagedServerPlan plan
    
    ApiServersIdMakeManagedServerAccount account
    /* Specific group to assign the server */
    Long provisionSiteId
    /* Metadata tags, Array of objects having a name and value, this adds or updates the specified tags and removes any tags not specified. */
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    
    ApiServersIdMakeManagedServerConfig config
}
