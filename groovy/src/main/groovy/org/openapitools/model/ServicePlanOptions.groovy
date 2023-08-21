package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ServicePlanOptions {
    /* Core Count */
    Long maxCores
    /* Cores Per Socket */
    Long coresPerSocket
    /* Memory in bytes For backwards compatability, values less than 1048576 are treated as being in MB and will be converted to bytes */
    Long maxMemory
}
