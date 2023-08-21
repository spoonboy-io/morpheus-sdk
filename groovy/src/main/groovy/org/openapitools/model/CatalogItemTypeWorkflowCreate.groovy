package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;

@Canonical
class CatalogItemTypeWorkflowCreate {
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
    
    ApiBlueprintsIdUpdatePermissionsResourcePermissionSites workflow
    /* Context for running the workflow, determines if a target resource must be selected. */
    String context
    /* Configuration object that contains settings for the workflow. */
    String workflowConfig
}
