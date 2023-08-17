package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.Arrays;
import org.openapitools.jackson.nullable.JsonNullable;
import org.openapitools.model.ApplianceSettingsEnabledZoneTypesInner;
import org.openapitools.model.CheckGroupInstance;

@Canonical
class CatalogItemType {
    
    Long id
    
    String name
    /* Useful shortcode for provisioning naming schemes and export reference. */
    String code
    /* Catalog Item Type category */
    String category
    
    String description
    
    List<String> labels
    
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
    
    CheckGroupInstance workflow
    
    String content
    
    List<Object> optionTypes
    
    String createdBy
    
    ApplianceSettingsEnabledZoneTypesInner owner
    
    Date dateCreated
    
    Date lastUpdated
}
