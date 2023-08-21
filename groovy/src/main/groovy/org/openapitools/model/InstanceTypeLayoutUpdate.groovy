package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions;
import org.openapitools.model.ClusterLayoutCreateEnvironmentVariables;
import org.openapitools.model.InstanceTypeCreatePriceSets;

@Canonical
class InstanceTypeLayoutUpdate {
    /* Layout name */
    String name
    
    List<String> labels = new ArrayList<String>()
    /* Version of the layout */
    String instanceVersion
    /* Layout description */
    String description
    /* Can be used to enable / disable the creatability of the layout. */
    Boolean creatable = true
    /* Provision type code */
    String provisionTypeCode
    /* Memory requirement in megabytes */
    String memoryRequirement
    /* Can be used to enable / disable the horizontal scaling. */
    Boolean hasAutoScale = false
    /* Can be used to enable / disable the supports convert to managed. */
    Boolean supportsConvertToManaged = false
    /* Array of layout node type IDs */
    List<Long> containerTypes = new ArrayList<Long>()
    /* Array of layout option type IDs */
    List<Long> optionTypes = new ArrayList<Long>()
    /* Array of layout spec template IDs */
    List<Long> specTemplates = new ArrayList<Long>()
    /* The environmentVariables parameter is array of env objects */
    List<ClusterLayoutCreateEnvironmentVariables> environmentVariables = new ArrayList<ClusterLayoutCreateEnvironmentVariables>()
    /* Array of price set objects */
    List<InstanceTypeCreatePriceSets> priceSets = new ArrayList<InstanceTypeCreatePriceSets>()
    
    ApiLibraryLayoutsIdPermissionsInstanceTypeLayoutPermissions permissions
}
