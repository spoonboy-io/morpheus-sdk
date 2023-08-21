package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.InstanceEnvsEnvs;

@Canonical
class InstanceEnvs {
    
    List<InstanceEnvsEnvs> envs = new ArrayList<InstanceEnvsEnvs>()
    
    List<InstanceEnvsEnvs> readOnlyEnvs = new ArrayList<InstanceEnvsEnvs>()
    
    List<InstanceEnvsEnvs> importedEnvs = new ArrayList<InstanceEnvsEnvs>()
}
