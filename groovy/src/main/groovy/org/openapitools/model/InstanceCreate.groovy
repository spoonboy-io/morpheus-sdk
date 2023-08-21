package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.model.InstanceCreateInstance;
import org.openapitools.model.InstanceCreateNetwork;
import org.openapitools.model.InstanceCreatePorts;
import org.openapitools.model.InstanceCreateVolume;

@Canonical
class InstanceCreate {
    
    InstanceCreateInstance instance
    /* The Cloud ID to provision the instance onto. */
    Long zoneId
    /* Environment Variables, an array of objects that have name and value. */
    List<ApiServersIdMakeManagedServerTags> evars = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Number of copies to provision. */
    Long copies = 1l
    /* Apply a multiply factor of containers/vms within the instance. */
    Long layoutSize = 1l
    /* Map of custom options depending on selected service plan. */
    Object servicePlanOptions
    /* Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. */
    List<Object> securityGroups = new ArrayList<Object>()
    /* The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of */
    List<InstanceCreateVolume> volumes = new ArrayList<InstanceCreateVolume>()
    /* The networkInterfaces parameter is for network configuration.  The Options API `/api/options/zoneNetworkOptions?zoneId=5&provisionTypeId=10` can be used to see which options are available.  */
    List<InstanceCreateNetwork> networkInterfaces = new ArrayList<InstanceCreateNetwork>()
    
    AnyOfinstancesConfigAzureinstancesConfigVMWareinstancesConfigGCPinstancesConfigAWSobject config = null
    /* Array of strings (keywords). */
    List<String> labels = new ArrayList<String>()
    /* Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Alias for `tags`. */
    List<ApiServersIdMakeManagedServerTags> metadata = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  */
    List<InstanceCreatePorts> ports = new ArrayList<InstanceCreatePorts>()
    /* The Workflow ID to execute. */
    Long taskSetId
    /* The Workflow Name to execute. */
    String taskSetName
}
