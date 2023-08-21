package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionType;

@Canonical
class CatalogType {
    
    Long id
    
    String name
    
    String description
    
    String type
    
    String context
    
    Boolean featured
    /* Can users order more than one of this item at a time. */
    Boolean allowQuantity
    
    String imagePath
    
    String darkImagePath
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
}
