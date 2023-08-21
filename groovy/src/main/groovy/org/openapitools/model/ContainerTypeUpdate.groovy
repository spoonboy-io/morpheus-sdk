package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ClusterLayoutCreateEnvironmentVariables;
import org.openapitools.model.ContainerTypeCreateContainerPorts;

@Canonical
class ContainerTypeUpdate {
    /* Node type name */
    String name
    
    List<String> labels = new ArrayList<String>()
    /* The short name is a name with no spaces used for display in your container list. */
    String shortName
    /* Node type description */
    String description
    /* Version of the node type */
    String containerVersion
    /* Provision type code, eg. `amazon`, etc. */
    String provisionTypeCode
    /* Array of script IDs. */
    List<Long> scripts = new ArrayList<Long>()
    /* Array of file template IDs. */
    List<Long> templates = new ArrayList<Long>()
    /* Virtual image ID */
    Long virtualImageId
    /* Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. */
    String statTypeCode
    /* Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. */
    String logTypeCode
    /* Server type.  Always pass \"vm\". */
    String serverType
    /* List of exposed port definitions in the format NAME=PORT|PROTOCOL */
    List<ContainerTypeCreateContainerPorts> containerPorts = new ArrayList<ContainerTypeCreateContainerPorts>()
    /* The environmentVariables parameter is array of env objects. */
    List<ClusterLayoutCreateEnvironmentVariables> environmentVariables = new ArrayList<ClusterLayoutCreateEnvironmentVariables>()
    /* Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. */
    Object config
}
