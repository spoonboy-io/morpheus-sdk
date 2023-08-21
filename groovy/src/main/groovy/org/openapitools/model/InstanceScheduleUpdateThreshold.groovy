package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class InstanceScheduleUpdateThreshold {
    /* Source Scale Threshold to apply as a template. All threshold settings with be copied from this threshold, and can be overridden by also passing each setting explicitly. */
    Long sourceThresholdId
    /* Auto Upscale */
    Boolean autoUp
    /* Auto Downscale */
    Boolean autoDown
    /* The minimum number of nodes to scale down to */
    Integer minCount
    /* The maximum number of nodes to scale up to */
    Integer maxCount
    /* Enable CPU Threshold */
    Boolean cpuEnabled
    /* Min CPU (%) */
    Double minCpu
    /* Max CPU (%) */
    Double maxCpu
    /* Enable Memory Threshold */
    Boolean memoryEnabled
    /* Min Memory (%) */
    Double minMemory
    /* Max Memory (%) */
    Double maxMemory
    /* Enable Disk Threshold */
    Boolean diskEnabled
    /* Min Disk (%) */
    Double minDisk
    /* Max Disk (%) */
    Double maxDisk
}
