package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ImageBuild;
import org.openapitools.model.ImageBuildExecution;

@Canonical
class InlineResponse20053 {
    
    ImageBuild imageBuild
    
    List<ImageBuildExecution> imageBuildExecutions = new ArrayList<ImageBuildExecution>()
}
