package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.InstanceTypesInstanceTypeLayouts;

@Canonical
class InstanceTypes {
    
    Long id
    
    InlineResponse20082LoadBalancerInstanceSslCert account
    
    String name
    
    List<String> labels = new ArrayList<String>()
    
    String code
    
    String description
    
    String provisionTypeCode
    
    String category
    
    Boolean active
    
    String environmentPrefix
    
    String visibility
    
    Boolean featured
    
    List<String> versions = new ArrayList<String>()
    
    List<InstanceTypesInstanceTypeLayouts> instanceTypeLayouts = new ArrayList<InstanceTypesInstanceTypeLayouts>()
    /* Logo image URL */
    String imagePath
    /* Dark logo image URL */
    String darkImagePath
}
