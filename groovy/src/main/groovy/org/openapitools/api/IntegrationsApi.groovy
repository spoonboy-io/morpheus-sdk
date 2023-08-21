package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AnyOfobjectmeta
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject100
import org.openapitools.model.InlineObject101
import org.openapitools.model.InlineResponse20061
import org.openapitools.model.InlineResponse20062
import org.openapitools.model.InlineResponse20063
import org.openapitools.model.InlineResponse20064
import org.openapitools.model.InlineResponse20065
import org.openapitools.model.Model200Success
import org.openapitools.model.OneOfintegrationConfigintegrationAnsibleConfigintegrationSNOWConfigintegrationSaltMasterConfigintegrationDockerRepoConfigintegrationGitRepoConfigintegrationGitHubConfig
import org.openapitools.model.UNKNOWN_BASE_TYPE

class IntegrationsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addIntegrationSnowObjects ( Long id, InlineObject100 inlineObject100, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/objects"

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
        bodyParams = inlineObject100


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addIntegrations ( UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = UNKNOWN_BASE_TYPE


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def getIntegrationInventory ( Long id, Long inventoryId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/inventory/${inventoryId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (inventoryId == null) {
            throw new RuntimeException("missing required params inventoryId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20065.class )

    }

    def getIntegrationObjects ( Long id, Long objectId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/objects/${objectId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (objectId == null) {
            throw new RuntimeException("missing required params objectId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20064.class )

    }

    def getIntegrationTypeOptionTypes ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integration-types/${id}/option-types"

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
                    InlineResponse20062.class )

    }

    def getIntegrationTypes ( Long id, Boolean optiontypes, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integration-types/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (optiontypes != null) {
            queryParams.put("optiontypes", optiontypes)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20061.class )

    }

    def getIntegrations ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}"

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

    def listIntegrationInventory ( Long id, Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/inventory"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
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

    def listIntegrationObjects ( Long id, Long max, Long offset, String sort, String direction, String phrase, String name, String value, Long refId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/objects"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
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
        if (value != null) {
            queryParams.put("value", value)
        }
        if (refId != null) {
            queryParams.put("refId", refId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20063.class )

    }

    def listIntegrationTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, Boolean optiontypes, String description, String category, Boolean creatable, Boolean enabled, Boolean hasCMDB, Boolean hasCMDBDiscovery, Boolean hasCM, Boolean hasDNS, Boolean hasApprovals, Boolean isPlugin, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integration-types"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }
        if (optiontypes != null) {
            queryParams.put("optiontypes", optiontypes)
        }
        if (description != null) {
            queryParams.put("description", description)
        }
        if (category != null) {
            queryParams.put("category", category)
        }
        if (creatable != null) {
            queryParams.put("creatable", creatable)
        }
        if (enabled != null) {
            queryParams.put("enabled", enabled)
        }
        if (hasCMDB != null) {
            queryParams.put("hasCMDB", hasCMDB)
        }
        if (hasCMDBDiscovery != null) {
            queryParams.put("hasCMDBDiscovery", hasCMDBDiscovery)
        }
        if (hasCM != null) {
            queryParams.put("hasCM", hasCM)
        }
        if (hasDNS != null) {
            queryParams.put("hasDNS", hasDNS)
        }
        if (hasApprovals != null) {
            queryParams.put("hasApprovals", hasApprovals)
        }
        if (isPlugin != null) {
            queryParams.put("isPlugin", isPlugin)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listIntegrations ( Long max, Long offset, String sort, String direction, String phrase, String name, Long id, String url, String host, String port, String username, Long version, String windowsVersion, String status, Boolean objects, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (id != null) {
            queryParams.put("id", id)
        }
        if (url != null) {
            queryParams.put("url", url)
        }
        if (host != null) {
            queryParams.put("host", host)
        }
        if (port != null) {
            queryParams.put("port", port)
        }
        if (username != null) {
            queryParams.put("username", username)
        }
        if (version != null) {
            queryParams.put("version", version)
        }
        if (windowsVersion != null) {
            queryParams.put("windowsVersion", windowsVersion)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (objects != null) {
            queryParams.put("objects", objects)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    AnyOfobjectmeta.class )

    }

    def refreshIntegrations ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/refresh"

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
                    "POST", "",
                    Object.class )

    }

    def removeIntegrationObjects ( Long id, Long objectId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/objects/${objectId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (objectId == null) {
            throw new RuntimeException("missing required params objectId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeIntegrations ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}"

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
                    "DELETE", "",
                    Model200Success.class )

    }

    def updateIntegrationInventory ( Long id, Long inventoryId, InlineObject101 inlineObject101, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}/inventory/${inventoryId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }
        // verify required params are set
        if (inventoryId == null) {
            throw new RuntimeException("missing required params inventoryId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject101


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateIntegrations ( Long id, UNKNOWN_BASE_TYPE UNKNOWN_BASE_TYPE, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/integrations/${id}"

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
        bodyParams = UNKNOWN_BASE_TYPE


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
