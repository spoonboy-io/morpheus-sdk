package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiEnvironmentsEnvironment {
    /* A unique name for the environment */
    String name
    /* A unique code for the environment */
    String code
    /* A description of the environment */
    String description
    /* private or public */
    String visibility = "private"
    /* Sort order */
    Long sortOrder = 0l
}