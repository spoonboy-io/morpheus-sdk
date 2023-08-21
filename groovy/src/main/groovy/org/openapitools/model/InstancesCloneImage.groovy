package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstancesCloneImage {
    /* Image Template Name */
    String templateName = "{server.name}-{timestamp}"
    /* Zone Folder externalId. This is required for VMware */
    String zoneFolder
}
