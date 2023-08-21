package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class UserSourceCreateLDAP {
    /* URL of Endpoint */
    String url
    /* Binding Username */
    String bindingUsername
    /* Binding Password */
    String bindingPassword
    /* User DN Expression */
    String requiredGroup
}
