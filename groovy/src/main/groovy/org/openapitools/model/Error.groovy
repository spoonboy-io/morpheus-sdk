package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class Error {
    /* Success indicator, true when the request succeeded and false when an error occurred */
    Boolean success = true
    /* Message containing a description of the result, usually a message about the error that occurred */
    String msg
    /* Validation errors, with a key for Object containing error messages for each invalid parameter (key) */
    Object errors
}
