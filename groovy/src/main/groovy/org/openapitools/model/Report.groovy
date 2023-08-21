package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse200107NetworkPoolCreatedBy;
import org.openapitools.model.ReportConfig;
import org.openapitools.model.ReportRows;
import org.openapitools.model.ReportType;

@Canonical
class Report {
    
    Long id
    
    ReportType type
    
    String reportTitle
    
    String filterTitle
    
    String status
    
    Date dateCreated
    
    Date lastUpdated
    
    Date startDate
    
    Date endDate
    
    ReportConfig config
    
    InlineResponse200107NetworkPoolCreatedBy createdBy
    
    List<ReportRows> rows = new ArrayList<ReportRows>()
}
