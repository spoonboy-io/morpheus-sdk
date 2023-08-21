package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;

@Canonical
class LogSettings {
    
    Boolean enabled
    
    String retentionDays
    
    List<Object> syslogRules = new ArrayList<Object>()
    
    List<Object> integrations = new ArrayList<Object>()
}
