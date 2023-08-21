package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.CatalogCartStats;
import org.openapitools.model.CatalogOrderCreateSuccessItems;

@Canonical
class CatalogOrderCreateSuccess {
    
    Long id
    
    String name
    
    List<CatalogOrderCreateSuccessItems> items = new ArrayList<CatalogOrderCreateSuccessItems>()
    
    CatalogCartStats stats
}
