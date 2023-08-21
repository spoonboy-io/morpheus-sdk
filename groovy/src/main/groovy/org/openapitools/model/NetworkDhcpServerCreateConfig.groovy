package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@Canonical
class NetworkDhcpServerCreateConfig {
    /* Edge Cluster */
    String edgeCluster
    /* Active Edge Node Options obtained by calling option source with :optionSource = nsxtEdgeNodes and networkServerId param */
    String preferredEdgeNode1
    /* Standby Edge Node Options obtained by calling option source with optionSource = nsxtEdgeNodes and networkServerId param */
    String preferredEdgeNode2
}
