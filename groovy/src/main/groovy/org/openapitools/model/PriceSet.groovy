package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.PriceSetPrices;

@Canonical
class PriceSet {
    
    Long id
    
    String name
    
    String code
    
    Boolean active
    
    String priceUnit
    
    String type
    
    String regionCode
    
    Boolean systemCreated
    
    String zone
    
    String zonePool
    
    String account
    
    List<PriceSetPrices> prices = new ArrayList<PriceSetPrices>()
}
