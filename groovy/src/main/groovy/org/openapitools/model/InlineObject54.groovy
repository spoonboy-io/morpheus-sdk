package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InlineObject54 {
    /* Url of desired template to apply to cluster */
    String serviceUrl
    /* Name or ID of desired Spec Template to apply to cluster */
    String specTemplate
    /* Yaml of template to apply to cluster */
    String specYaml
}
