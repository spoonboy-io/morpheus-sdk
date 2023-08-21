package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class CheckType {
    
    Long id
    
    String code
    
    String name
    
    Long defaultInterval
    
    String metricName
    
    Boolean inUptime
    
    Boolean createIncident
    
    Boolean pushOnly
    
    Boolean tunnelSupported
}
