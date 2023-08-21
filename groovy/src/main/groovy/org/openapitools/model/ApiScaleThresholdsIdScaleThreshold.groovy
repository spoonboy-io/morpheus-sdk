package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class ApiScaleThresholdsIdScaleThreshold {
    /* A name for the scale threshold */
    String name
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
    BigDecimal minCpu = new BigDecimal("0")
    /* Max CPU (%) */
    BigDecimal maxCpu = new BigDecimal("0")
    /* Enable Memory Threshold */
    Boolean memoryEnabled = false
    /* Min Memory (%) */
    BigDecimal minMemory = new BigDecimal("0")
    /* Max Memory (%) */
    BigDecimal maxMemory = new BigDecimal("0")
    /* Enable Disk Threshold */
    Boolean diskEnabled = false
    /* Min Disk (%) */
    BigDecimal minDisk = new BigDecimal("0")
    /* Max Disk (%) */
    BigDecimal maxDisk = new BigDecimal("0")
}
