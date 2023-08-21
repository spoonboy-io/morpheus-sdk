package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.model.SecurityScanSecurityPackage;

@Canonical
class SecurityScan {
    
    Long id
    
    SecurityScanSecurityPackage securityPackage
    
    String status
    
    Date scanDate
    
    Long scanDuration
    
    Long testCount
    
    Long runCount
    
    Long passCount
    
    Long failCount
    
    Long otherCount
    
    BigDecimal scanScore
    
    String externalId
    
    String createdBy
    
    String updatedBy
    
    Date dateCreated
    
    Date lastUpdated
    /* Results Summary (only returned when using query parameter results=true) */
    Object results
}
