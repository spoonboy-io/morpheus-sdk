package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiPluginsIdPlugin {
    /* Can be used to disable the plugin */
    Boolean enabled
    /* Configuration object that contains settings for the applicable option types. */
    Object config
}
