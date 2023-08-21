package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionType;

@Canonical
class InlineResponse20077LoadBalancerType {
    
    Long id
    
    String name
    
    String code
    
    Boolean enabled
    
    Boolean internal
    
    Boolean creatable
    
    String createType
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    List<OptionType> vipOptionTypes = new ArrayList<OptionType>()
}
