package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateActiveDirectory {
    /* AD Server IP/FQDN */
    String url
    /* Domain */
    String domain
    /* Use SSL */
    Boolean useSSL = false
    /* Binding Username */
    String bindingUsername
    /* Binding Password */
    String bindingPassword
    /* Required Group */
    String requiredGroup
    /* Include Member Groups */
    Boolean searchMemberGroups = false
}
