package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ForgotPasswordRequest {
    /* Specified as \"username\" or \"tenantId\\username\" with proper HTML encoding and used in conjunction with `password`.  */
    String username
}
