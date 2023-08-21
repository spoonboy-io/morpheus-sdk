package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20040AppDeployInstance;
import org.openapitools.model.InlineResponse20082LoadBalancerInstanceSslCert;
import org.openapitools.model.OptionType;

@Canonical
class CatalogItemType {
    
    Long id
    
    String name
    /* Useful shortcode for provisioning naming schemes and export reference. */
    String code
    /* Catalog Item Type category */
    String category
    
    String description
    
    List<String> labels = new ArrayList<String>()
    
    String type
    
    Boolean enabled
    
    Boolean featured
    /* Can users order more than one of this item at a time. */
    Boolean allowQuantity
    
    String iconPath
    
    String imagePath
    
    String darkImagePath
    
    String visibility
    
    String layoutCode
    
    Object blueprint
    
    String appSpec
    
    Object config
    
    InlineResponse20082LoadBalancerInstanceSslCert workflow
    
    String content
    
    List<OptionType> optionTypes = new ArrayList<OptionType>()
    
    String createdBy
    
    InlineResponse20040AppDeployInstance owner
    
    Date dateCreated
    
    Date lastUpdated
}
