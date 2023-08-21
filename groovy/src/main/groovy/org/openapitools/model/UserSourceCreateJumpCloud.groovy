package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateJumpCloud {
    /* Organization ID */
    Boolean organizationId = false
    /* Binding Username */
    String bindingUsername
    /* Binding Password */
    String bindingPassword
    /* Required group name */
    String requiredRole
}
