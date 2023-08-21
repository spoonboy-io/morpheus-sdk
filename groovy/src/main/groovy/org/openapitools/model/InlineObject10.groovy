package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InlineObject10 {
    /* The secret Reset Password token that was included in the **Forgot Password Email**. */
    String token
    /* User new password. This is the new password for your user. */
    String password
}
