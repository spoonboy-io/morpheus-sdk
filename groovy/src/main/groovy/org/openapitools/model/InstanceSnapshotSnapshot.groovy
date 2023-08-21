package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceSnapshotSnapshot {
    /* Optional name for the snapshot being created. */
    String name = "{serverName}.{timestamp}"
    /* Optional description for the snapshot */
    String description
}
