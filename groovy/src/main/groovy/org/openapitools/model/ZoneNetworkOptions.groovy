package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ZoneNetworkOptionsNetworkTypes;
import org.openapitools.model.ZoneNetworkOptionsNetworks;

@Canonical
class ZoneNetworkOptions {
    
    List<ZoneNetworkOptionsNetworks> networks = new ArrayList<ZoneNetworkOptionsNetworks>()
    
    List<Object> networkGroups = new ArrayList<Object>()
    
    List<ZoneNetworkOptionsNetworkTypes> networkTypes = new ArrayList<ZoneNetworkOptionsNetworkTypes>()
    
    List<Object> networkSubnets = new ArrayList<Object>()
    
    Boolean hasNetworks
    
    Long maxNetworks
    
    String enableNetworkTypeSelection
    
    Boolean supportsNetworkSelection
}
