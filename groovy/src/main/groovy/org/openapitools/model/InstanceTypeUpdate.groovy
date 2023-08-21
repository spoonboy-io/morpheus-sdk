package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterLayoutCreateEnvironmentVariables;
import org.openapitools.model.InstanceTypeCreatePriceSets;

@Canonical
class InstanceTypeUpdate {
    /* Instance type name */
    String name
    /* Instance type description */
    String description
    /* Array of label strings, can be used for filtering. */
    List<String> labels = new ArrayList<String>()
    /* Instance type code */
    String code
    /* Category */
    String category
    /* Visibility */
    String visibility = VisibilityEnum.PRIVATE
    /* Featured */
    Boolean featured
    /* Enable Settings */
    Boolean hasSettings
    /* Enable Scaling (Horizontal) */
    Boolean hasAutoScale
    /* Supports Deployments */
    Boolean hasDeployment
    /* Environment Prefix, can be used to make exported evars unique. */
    String environmentPrefix
    /* Array of instance type env variables. */
    List<ClusterLayoutCreateEnvironmentVariables> environmentVariables = new ArrayList<ClusterLayoutCreateEnvironmentVariables>()
    /* Array of price set objects */
    List<InstanceTypeCreatePriceSets> priceSets = new ArrayList<InstanceTypeCreatePriceSets>()
    /* Array of instance type option type IDs */
    List<Long> optionTypes = new ArrayList<Long>()
}
