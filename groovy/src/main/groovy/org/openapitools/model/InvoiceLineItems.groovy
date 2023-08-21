package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class InvoiceLineItems {
    
    Long id
    
    Long invoiceId
    
    String refType
    
    Long refId
    
    String refName
    
    Date startDate
    
    Date endDate
    
    String itemId
    
    String itemType
    
    String itemName
    
    String itemDescription
    
    String productId
    
    String productCode
    
    String productName
    
    String itemSeller
    
    String itemAction
    
    String externalId
    
    String rateId
    
    String rateClass
    
    String rateUnit
    
    String rateTerm
    
    String usageType
    
    String usageCategory
    
    String usageService
    
    BigDecimal itemUsage
    
    BigDecimal itemRate
    
    BigDecimal itemCost
    
    BigDecimal itemPrice
    
    Long itemTax
    
    String itemTerm
    
    String taxType
    
    String regionCode
    
    String currency
    
    Long conversionRate
    
    Date dateCreated
    
    Date lastUpdated
}
