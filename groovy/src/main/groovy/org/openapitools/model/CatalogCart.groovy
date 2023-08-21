package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.CatalogCartStats;
import org.openapitools.model.CatalogItem;

@Canonical
class CatalogCart {
    
    Long id
    
    String name
    
    List<CatalogItem> items = new ArrayList<CatalogItem>()
    
    CatalogCartStats stats
}
