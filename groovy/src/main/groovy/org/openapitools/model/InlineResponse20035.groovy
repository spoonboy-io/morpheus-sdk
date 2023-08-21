package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InlineResponse20035Actions;

@Canonical
class InlineResponse20035 {
    
    List<Long> containerIds = new ArrayList<Long>()
    
    List<InlineResponse20035Actions> actions = new ArrayList<InlineResponse20035Actions>()
}
