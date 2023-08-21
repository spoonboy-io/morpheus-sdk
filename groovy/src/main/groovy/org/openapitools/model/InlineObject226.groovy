package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiServersChangeCloudServers;

@Canonical
class InlineObject226 {
    /* The cloud/zone ID we are moving the set of servers to */
    Long cloudId
    /* A JSON array of source: and target: server ids to be moved. If the target is blank Morpheus will automatically try to match by the servers unique or externalId */
    List<ApiServersChangeCloudServers> servers = new ArrayList<ApiServersChangeCloudServers>()
}
