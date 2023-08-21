package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.model.PriceSetVolumeType;

@Canonical
class PriceSetPrices {
    
    Long id
    
    String name
    
    String code
    
    Boolean active
    
    String priceType
    
    String priceUnit
    
    String additionalPriceUnit
    
    BigDecimal price
    
    BigDecimal customPrice
    
    String markupType
    
    Long markup
    
    BigDecimal markupPercent
    
    BigDecimal cost
    
    String currency
    
    String incurCharges
    
    String platform
    
    String software
    
    PriceSetVolumeType volumeType
    
    String datastore
    
    Boolean crossCloudApply
    
    String account
}
