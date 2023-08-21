package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;

@Canonical
class SearchHits {
    /* ID */
    String id
    /* UUID */
    String uuid
    /* Name */
    String name
    
    String description
    
    String type
    
    Date dateCreated
    
    BigDecimal score
}
