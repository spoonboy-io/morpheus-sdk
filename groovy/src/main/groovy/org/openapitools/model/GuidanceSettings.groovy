package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class GuidanceSettings {
    /* Power Shutdown Average CPU (%). Lower limit for average CPU usage */
    Integer cpuAvgCutoffPower
    /* Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage */
    Integer cpuMaxCutoffPower
    /* Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth */
    Integer networkCutoffPower
    /* CPU Up-size Average CPU (%). Upper limit for CPU usage */
    Integer cpuUpAvgStandardCutoffRightSize
    /* CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage */
    Integer cpuUpMaxStandardCutoffRightSize
    /* Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage */
    Integer memoryUpAvgStandardCutoffRightSize
    /* Memory Down-size Maximum Free Memory (%). Upper limit for average free memory */
    Integer memoryDownAvgStandardCutoffRightSize
    /* Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage */
    Integer memoryDownMaxStandardCutoffRightSize
}
