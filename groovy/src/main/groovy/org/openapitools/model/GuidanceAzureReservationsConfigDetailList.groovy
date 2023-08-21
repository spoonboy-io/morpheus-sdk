package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class GuidanceAzureReservationsConfigDetailList {
    
    String id
    
    String apiName
    
    String apiType
    
    String externalId
    
    String period
    
    String name
    
    String type
    
    String category
    
    String size
    
    String region
    
    String term
    
    String meterId
    
    Long onDemandCount
    
    BigDecimal onDemandCost
    
    Long reservedCount
    
    Long reservedCost
    
    Long recommendedCount
    
    BigDecimal recommendedCost
    
    BigDecimal totalSavings
    
    BigDecimal totalSavingsPercent
}
