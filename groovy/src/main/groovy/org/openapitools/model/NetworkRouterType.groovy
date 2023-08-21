package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.OptionType;

@Canonical
class NetworkRouterType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    Boolean enabled
    
    Boolean creatable
    
    Boolean selectable
    
    Boolean hasFirewall
    
    Boolean hasDhcp
    
    Boolean hasRouting
    
    Boolean hasNetworkServer
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    List<OptionType> ruleOptionTypes = new ArrayList<OptionType>()
    
    List<OptionType> ruleGroupOptionTypes = new ArrayList<OptionType>()
    
    List<OptionType> natOptionTypes = new ArrayList<OptionType>()
    
    List<OptionType> bgpOptionTypes = new ArrayList<OptionType>()
}
