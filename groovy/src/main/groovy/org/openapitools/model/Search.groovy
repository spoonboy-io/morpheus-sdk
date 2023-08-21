package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.MetaObject;
import org.openapitools.model.SearchHits;

@Canonical
class Search {
    
    List<SearchHits> hits = new ArrayList<SearchHits>()
    
    String query
    
    Long took
    
    MetaObject meta
}
