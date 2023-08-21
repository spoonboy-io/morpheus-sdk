package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.Check;
import org.openapitools.model.CheckGroup;

@Canonical
class InlineResponse20019 {
    
    CheckGroup checkGroup
    
    List<Check> checks = new ArrayList<Check>()
}
