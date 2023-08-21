package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiBlueprintsIdUpdatePermissionsResourcePermissionSites;
import org.openapitools.model.ApiServicePlansServicePlanProvisionType;
import org.openapitools.model.ClusterLayoutCreateEnvironmentVariables;
import org.openapitools.model.ClusterLayoutCreateGroupType;
import org.openapitools.model.ClusterLayoutCreateMasters;

@Canonical
class ClusterLayoutCreate {
    /* Cluster layout name */
    String name
    /* Cluster layout description */
    String description
    
    List<String> labels = new ArrayList<String>()
    /* Version of the cluster layout */
    String computeVersion
    /* Can be used to enable / disable the creatability of the cluster layout. */
    Boolean creatable = true
    /* Can be used to enable / disable the horizontal scaling. */
    Boolean hasAutoScale = false
    /* Install Docker (container runtime). */
    Boolean installContainerRuntime = false
    /* Memory requirement in bytes */
    Long memoryRequirement
    
    ClusterLayoutCreateGroupType groupType
    
    ApiServicePlansServicePlanProvisionType provisionType
    /* Array of cluster layout option types */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> optionTypes = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    /* Array of cluster layout task sets */
    List<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites> taskSets = new ArrayList<ApiBlueprintsIdUpdatePermissionsResourcePermissionSites>()
    /* Array of cluster layout env variables */
    List<ClusterLayoutCreateEnvironmentVariables> environmentVariables = new ArrayList<ClusterLayoutCreateEnvironmentVariables>()
    /* Array of cluster layout master nodes */
    List<ClusterLayoutCreateMasters> masters = new ArrayList<ClusterLayoutCreateMasters>()
    /* Array of cluster layout worker nodes */
    List<ClusterLayoutCreateMasters> workers = new ArrayList<ClusterLayoutCreateMasters>()
}
