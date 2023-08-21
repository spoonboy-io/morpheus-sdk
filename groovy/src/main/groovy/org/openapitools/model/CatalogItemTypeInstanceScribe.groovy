package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.model.CatalogItemTypeInstanceScribeCloud;
import org.openapitools.model.CatalogItemTypeInstanceScribeGroup;
import org.openapitools.model.CatalogItemTypeInstanceScribeLayout;
import org.openapitools.model.CatalogItemTypeInstanceScribePlan;
import org.openapitools.model.CatalogItemTypeInstanceScribePorts;
import org.openapitools.model.CatalogItemTypeInstanceScribeSecurityGroups;
import org.openapitools.model.InstanceCreateNetwork;
import org.openapitools.model.InstanceCreateVolume;
import org.openapitools.model.ServicePlanOptions;

@Canonical
class CatalogItemTypeInstanceScribe {
    
    CatalogItemTypeInstanceScribeGroup group
    
    CatalogItemTypeInstanceScribeCloud cloud
    /* The type of instance by code we want to fetch. */
    String type
    /* Name of the instance to be created. */
    String name
    
    AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject config = null
    /* The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of */
    List<InstanceCreateVolume> volumes = new ArrayList<InstanceCreateVolume>()
    /* Hostname of the instance to be created.  Can be the same as the instance name. */
    String hostName
    /* Environment code */
    String environment
    
    CatalogItemTypeInstanceScribeLayout layout
    
    CatalogItemTypeInstanceScribePlan plan
    /* Version of the layout to create. */
    String version
    /* Environment Variables, an array of objects that have name and value. */
    List<ApiServersIdMakeManagedServerTags> evars = new ArrayList<ApiServersIdMakeManagedServerTags>()
    
    ServicePlanOptions servicePlanOptions
    /* Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. */
    List<CatalogItemTypeInstanceScribeSecurityGroups> securityGroups = new ArrayList<CatalogItemTypeInstanceScribeSecurityGroups>()
    /* The networkInterfaces parameter is for network configuration.  The Options API `/api/options/zoneNetworkOptions?zoneId=5&provisionTypeId=10` can be used to see which options are available.  */
    List<InstanceCreateNetwork> networkInterfaces = new ArrayList<InstanceCreateNetwork>()
    /* Array of strings (keywords). */
    List<String> labels = new ArrayList<String>()
    /* Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Alias for `tags`. */
    List<ApiServersIdMakeManagedServerTags> metadata = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  */
    List<CatalogItemTypeInstanceScribePorts> ports = new ArrayList<CatalogItemTypeInstanceScribePorts>()
    /* The Workflow ID to execute. */
    Long taskSetId
    /* The Workflow Name to execute. */
    String taskSetName
}
