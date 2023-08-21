package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.LoadBalancerCreateResourcePermission;
import org.openapitools.model.LoadBalancerCreateTenants;

@Canonical
class LoadBalancerCreate {
    /* Name */
    String name
    /* Description */
    String description
    /* Network Server ID */
    Long networkServerId
    /* Configuration object with parameters that vary by load balancer type. */
    Object config
    /* private or public */
    String visibility = "public"
    /* Array of tenant account ids that are allowed access */
    List<LoadBalancerCreateTenants> tenants = new ArrayList<LoadBalancerCreateTenants>()
    
    LoadBalancerCreateResourcePermission resourcePermission
}
