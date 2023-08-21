package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class ApiServicePlansServicePlanConfigRanges {
    /* Custom min storage in GB */
    String minStorage
    /* Custom max storage in GB */
    String maxStorage
    /* Custom min per disk size in GB */
    String minPerDiskSize
    /* Custom max per disk size in GB */
    String maxPerDiskSize
    /* Custom min memory in bytes */
    Long minMemory
    /* Custom max memory in bytes */
    Long maxMemory
    /* Custom min cores */
    String minCores
    /* Custom max cores */
    String maxCores
    /* Custom min sockets */
    String minSockets
    /* Custom max sockets */
    String maxSockets
    /* Custom min cores allowed per socket */
    String minCoresPerSocket
    /* Custom max cores allowed per socket */
    String maxCoresPerSocket
}
