package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceScheduleCreateThreshold {
    /* Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. */
    Long sourceThresholdId
    /* Auto Upscale */
    Boolean autoUp = false
    /* Auto Downscale */
    Boolean autoDown = false
    /* The minimum number of nodes to scale down to */
    Integer minCount
    /* The maximum number of nodes to scale up to */
    Integer maxCount
    /* Enable CPU Threshold */
    Boolean cpuEnabled = false
    /* Min CPU (%) */
    Double minCpu = 0d
    /* Max CPU (%) */
    Double maxCpu = 0d
    /* Enable Memory Threshold */
    Boolean memoryEnabled = false
    /* Min Memory (%) */
    Double minMemory = 0d
    /* Max Memory (%) */
    Double maxMemory = 0d
    /* Enable Disk Threshold */
    Boolean diskEnabled = false
    /* Min Disk (%) */
    Double minDisk = 0d
    /* Max Disk (%) */
    Double maxDisk = 0d
}
