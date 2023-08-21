package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject274
import org.openapitools.model.InlineObject41
import org.openapitools.model.InlineObject42
import org.openapitools.model.InlineObject43
import org.openapitools.model.InlineObject44
import org.openapitools.model.InlineObject45
import org.openapitools.model.InlineObject46
import org.openapitools.model.InlineObject47
import org.openapitools.model.InlineObject49
import org.openapitools.model.InlineResponse200168
import org.openapitools.model.InlineResponse20020
import org.openapitools.model.InlineResponse20021
import org.openapitools.model.InlineResponse20022
import org.openapitools.model.InlineResponse20023
import org.openapitools.model.InlineResponse20024
import org.openapitools.model.Model200Success

class CloudsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCloudResourcePool ( BigDecimal zoneId, InlineObject45 inlineObject45, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/resource-pools"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject45


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse20024.class )

    }

    def addClouds ( InlineObject41 inlineObject41, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject41


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getCloudDatastores ( BigDecimal zoneId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/data-stores/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getCloudFolders ( BigDecimal zoneId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/folders/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getCloudResourcePools ( BigDecimal zoneId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/resource-pools/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getCloudTypes ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zone-types/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20020.class )

    }

    def getClouds ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20021.class )

    }

    def getWikiCloud ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200168.class )

    }

    def listCloudDatastores ( BigDecimal zoneId, String name, String phrase, Long max, String sort, String direction, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/data-stores"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }

        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (max != null) {
            queryParams.put("max", max)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCloudFolders ( BigDecimal zoneId, String name, String phrase, Long max, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/folders"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }

        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (max != null) {
            queryParams.put("max", max)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCloudResourcePools ( BigDecimal zoneId, String name, String phrase, Long max, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/resource-pools"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }

        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (max != null) {
            queryParams.put("max", max)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCloudSecurityGroups ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/security-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCloudTypes ( Long max, Long offset, String sort, String direction, String name, String code, String phrase, String provisionType, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zone-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (provisionType != null) {
            queryParams.put("provisionType", provisionType)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClouds ( Date lastUpdated, String type, Long groupId, Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (type != null) {
            queryParams.put("type", type)
        }
        if (groupId != null) {
            queryParams.put("groupId", groupId)
        }
        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (sort != null) {
            queryParams.put("sort", sort)
        }
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def refreshClouds ( Long id, InlineObject47 inlineObject47, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/refresh"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject47


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def removeCloudResourcePools ( BigDecimal zoneId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/resource-pools/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeClouds ( Long id, Boolean removeResources, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (removeResources != null) {
            queryParams.put("removeResources", removeResources)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def updateCloudDatastores ( BigDecimal zoneId, Long id, InlineObject43 inlineObject43, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/data-stores/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject43


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20022.class )

    }

    def updateCloudFolders ( BigDecimal zoneId, Long id, InlineObject44 inlineObject44, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/folders/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject44


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20023.class )

    }

    def updateCloudLogo ( Long id, File logo, File darkLogo, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/update-logo"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }




        contentType = 'multipart/form-data';
        bodyParams = [:]
        bodyParams.put("logo", logo)
        bodyParams.put("darkLogo", darkLogo)

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateCloudResourcePool ( BigDecimal zoneId, Long id, InlineObject46 inlineObject46, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${zoneId}/resource-pools/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject46


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20024.class )

    }

    def updateCloudSecurityGroups ( Long id, InlineObject49 inlineObject49, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/security-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject49


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def updateClouds ( Long id, InlineObject42 inlineObject42, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject42


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiCloud ( Long id, InlineObject274 inlineObject274, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/zones/${id}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject274


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
