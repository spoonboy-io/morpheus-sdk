package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import org.openapitools.model.ApiServersIdMakeManagedServer;

@Canonical
class InlineObject223 {
    
    ApiServersIdMakeManagedServer server
    /* Install agent. Set to false to manually install agent instead. */
    Boolean installAgent = true
    /* Instance Type ID for the new Instance */
    Long instanceTypeId
    /* Layout ID for the new Instance */
    Long layout
}
