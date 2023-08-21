package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.LogData;
import org.openapitools.model.LogSort;

@Canonical
class Log {
    
    LogSort sort
    
    Long offset
    
    Date start
    
    Date end
    
    List<LogData> data = new ArrayList<LogData>()
    
    Long max
    
    Long grandTotal
    
    Long total
    
    Boolean success
    
    Long count
}
