package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import org.openapitools.model.ApiPricesPriceAccount;
import org.openapitools.model.ApiPricesPriceDatastore;
import org.openapitools.model.ApiPricesPriceVolumeType;

@Canonical
class ApiPricesIdPrice {
    /* Price name */
    String name
    /* Price code, must be unique */
    String code
    
    ApiPricesPriceAccount account
    /* Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software * `load_balancer` - Load Balancer * `load_balancer_virtual_server` - Load Balancer Virtual Server  */
    String priceType
    /* The unit of pricing */
    String priceUnit
    /* Indicates when to incur charge */
    String incurCharges
    /* ISO Currency code */
    String currency
    /* Cost */
    BigDecimal cost
    /* Price adjustment type */
    String markupType
    /* Amount for `fixed` price adjustment type */
    BigDecimal markup
    /* Percent for `percent` price adjustment type */
    BigDecimal markupPercent
    /* Custom price for `custom` price adjustment type */
    BigDecimal customPrice
    /* Platform.  Required for `platform` price type */
    String platform
    /* Software.  Required for software price type */
    String software
    
    ApiPricesPriceVolumeType volumeType
    
    ApiPricesPriceDatastore datastore
    /* Apply price across clouds, optional true/false flag for datastore price type */
    Boolean crossCloudApply
}
