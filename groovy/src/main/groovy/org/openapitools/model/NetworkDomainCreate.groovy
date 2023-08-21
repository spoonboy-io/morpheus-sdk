package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkDomainCreate {
    /* Name */
    String name
    /* Description */
    String description
    /* Overrides displayed name in domain selection components. Useful if using many OU Paths. */
    String displayName
    /* Public Zone */
    Boolean publicZone = false
    /* Workflow ID. Workflows can be applied to an instance when associated with a domain. Useful for custom domain related scripting. (Important if wanting to ensure the computer is removed from the domain using teardown phased workflows.)  */
    Long taskSetId
    /* Active */
    Boolean active
    /* Join Domain Controller */
    Boolean domainController = true
    /* Domain Username */
    String domainUsername
    /* Domain Password */
    String domainPassword
    /* DC Server. If specified, must be the server name and not an IP Address */
    String dcServer
    /* OU Path. (i.e. 'OU=staging,DC=ad,DC=yourdomain,DC=com') */
    String ouPath
    /* Guest Username. If set, will change the instances RPC Service User after joining a Domain. */
    String guestUsername
    /* Guest Password */
    String guestPassword
}
