package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class BackupLastResult {
    /* Last Result ID */
    Long id
    /* Last Result Status */
    String name
    /* Last Result Date Created */
    Date dateCreated
}
