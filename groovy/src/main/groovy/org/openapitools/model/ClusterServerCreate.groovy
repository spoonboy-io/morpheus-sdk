package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersIdMakeManagedServerTags;
import org.openapitools.model.ClusterServerCreateNetworkInterfaces;
import org.openapitools.model.ClusterServerCreatePlan;
import org.openapitools.model.ClusterServerCreateServerType;
import org.openapitools.model.ClusterServerCreateUserGroup;
import org.openapitools.model.ClusterServerCreateVolumes;

@Canonical
class ClusterServerCreate {
    /* Key for specific host type configuration  The config parameter is for configuration options that are specific to each Provision Type. The Provision Types api can be used to see which options are available.  */
    Object config
    
    ClusterServerCreateServerType serverType
    /* Name to be used for host(s) created in the cluster */
    String name
    
    ClusterServerCreatePlan plan
    /* The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects */
    List<ClusterServerCreateVolumes> volumes = new ArrayList<ClusterServerCreateVolumes>()
    /* The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  */
    List<ClusterServerCreateNetworkInterfaces> networkInterfaces = new ArrayList<ClusterServerCreateNetworkInterfaces>()
    /* Key for security group configuration. */
    List<String> securityGroups = new ArrayList<String>()
    /* Visibility for server host */
    String visibility = "private"
    
    ClusterServerCreateUserGroup userGroup
    /* Network domain */
    String networkDomain
    /* Hostname for server host */
    String hostname
    /* Number of workers or hosts */
    Long nodeCount
    /* Metadata tags, Array of objects having a name and value. */
    List<ApiServersIdMakeManagedServerTags> tags = new ArrayList<ApiServersIdMakeManagedServerTags>()
    /* Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. */
    List<String> labels = new ArrayList<String>()
}
