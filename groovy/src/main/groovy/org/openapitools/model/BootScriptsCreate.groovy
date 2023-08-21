package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BootScriptsCreate {
    /* A name for the boot script */
    String fileName
    /* The script content */
    String content
}
