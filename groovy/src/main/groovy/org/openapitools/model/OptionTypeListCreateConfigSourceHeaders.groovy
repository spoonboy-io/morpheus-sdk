package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class OptionTypeListCreateConfigSourceHeaders {
    /* Header name */
    String name
    /* Header value */
    String value
    /* Can be used to enable / disable masking of value */
    Boolean masked = false
}
