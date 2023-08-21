package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.StorageServerTypeGroupOptionTypes;
import org.openapitools.model.StorageServerTypeOptionTypes;
import org.openapitools.model.StorageServerTypeStorageVolumeTypes;

@Canonical
class StorageServerType {
    
    Long id
    
    String code
    
    String name
    
    String description
    
    Boolean enabled
    
    Boolean creatable
    
    Boolean hasNamespaces
    
    Boolean hasGroups
    
    Boolean hasBlock
    
    Boolean hasObject
    
    Boolean hasFile
    
    Boolean hasDatastore
    
    Boolean hasDisks
    
    Boolean hasHosts
    
    Boolean createNamespaces
    
    Boolean createGroup
    
    Boolean createBlock
    
    Boolean createObject
    
    Boolean createFile
    
    Boolean createDatastore
    
    Boolean createDisk
    
    Boolean createHost
    
    String iconCode
    
    Boolean hasFileBrowser
    
    List<StorageServerTypeOptionTypes> optionTypes = new ArrayList<StorageServerTypeOptionTypes>()
    
    List<StorageServerTypeGroupOptionTypes> groupOptionTypes = new ArrayList<StorageServerTypeGroupOptionTypes>()
    
    List<Object> bucketOptionTypes = new ArrayList<Object>()
    
    List<Object> shareOptionTypes = new ArrayList<Object>()
    
    List<Object> shareAccessOptionTypes = new ArrayList<Object>()
    
    List<StorageServerTypeStorageVolumeTypes> storageVolumeTypes = new ArrayList<StorageServerTypeStorageVolumeTypes>()
}
