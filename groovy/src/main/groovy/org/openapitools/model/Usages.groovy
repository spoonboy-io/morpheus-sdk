package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.MetaObject;
import org.openapitools.model.UsagesActivity;

@Canonical
class Usages {
    
    List<UsagesActivity> activity = new ArrayList<UsagesActivity>()
    
    MetaObject meta = null
}
