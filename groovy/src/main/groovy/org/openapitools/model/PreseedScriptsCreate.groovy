package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class PreseedScriptsCreate {
    /* A name for the preseed script */
    String fileName
    /* The script content */
    String content
}
