package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.CatalogItemTypeInstanceScribe;

@Canonical
class CatalogItemTypeInstanceCreate {
    /* Catalog Item Type name */
    String name
    /* Useful shortcode for provisioning naming schemes and export reference. */
    String code
    /* Catalog Item Type category */
    String category
    /* Catalog Item Type description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Type, `instance`, `blueprint` or `workflow`. This determines whether an Instance or App will be provisioned. Instance types require a config and blueprint requires a blueprint and appSpec, while workflow types requires a workflow and context. */
    String type
    /* Visibility - Set to public to allow all tenants */
    String visibility = "private"
    /* Identifier primarily used for Plugin Catalog Item Types */
    String layoutCode
    /* Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. */
    String iconPath
    /* Can be used to enable / disable the catalog item type. */
    Boolean enabled = true
    /* Can be used to feature the catalog item type. */
    Boolean featured = false
    /* Can users order more than one of this item at a time. */
    Boolean allowQuantity = false
    
    CatalogItemTypeInstanceScribe config
    /* Array of option type IDs. Only applies to type instance and blueprint. */
    List<Long> optionTypes = new ArrayList<Long>()
    /* Documentation content for this Catalog Item. Markdown-formatted text is accepted and displayed appropriately when the item is ordered from the Service Catalog. A new Catalog Item-type Wiki entry will also be added containing this information. */
    String content
}
