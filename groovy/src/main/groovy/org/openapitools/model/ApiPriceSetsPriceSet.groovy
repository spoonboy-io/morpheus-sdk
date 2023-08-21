package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiPriceSetsPriceSetZone;
import org.openapitools.model.ApiPriceSetsPriceSetZonePool;

@Canonical
class ApiPriceSetsPriceSet {
    /* Price set name */
    String name
    /* Price set code. Must be unique. */
    String code
    /* Price set region code */
    String regionCode
    
    ApiPriceSetsPriceSetZone zone
    
    ApiPriceSetsPriceSetZonePool zonePool
    /* Price Unit */
    String priceUnit
    /* Price set type */
    String type
    
    List<Long> prices = new ArrayList<Long>()
}
