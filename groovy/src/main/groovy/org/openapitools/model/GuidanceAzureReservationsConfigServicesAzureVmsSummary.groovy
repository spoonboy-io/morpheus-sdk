package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class GuidanceAzureReservationsConfigServicesAzureVmsSummary {
    
    BigDecimal totalSavings
    
    String currencyCode
    
    BigDecimal totalSavingsPercent
    
    String term
    
    String paymentOption
    
    String service
    
    Long onDemandCount
    
    BigDecimal onDemandCost
    
    Long reservedCount
    
    Long reservedCost
    
    Long recommendedCount
    
    BigDecimal recommendedCost
}
