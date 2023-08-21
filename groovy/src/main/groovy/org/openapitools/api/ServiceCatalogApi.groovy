package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.CatalogCartItemCreate
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject227
import org.openapitools.model.InlineResponse200139
import org.openapitools.model.InlineResponse200140
import org.openapitools.model.Model200Success

class ServiceCatalogApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCatalogCart ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/checkout"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addCatalogCartItem ( Boolean validate, CatalogCartItemCreate catalogCartItemCreate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/cart/items"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (validate != null) {
            queryParams.put("validate", validate)
        }


        contentType = 'application/json';
        bodyParams = catalogCartItemCreate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def addCatalogOrder ( Boolean validate, InlineObject227 inlineObject227, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/orders"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (validate != null) {
            queryParams.put("validate", validate)
        }


        contentType = 'application/json';
        bodyParams = inlineObject227


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deleteCatalogCart ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/cart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteCatalogCartItem ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/cart/items/${id}"

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

    def deleteCatalogItem ( Long id, String preserveVolumes, String keepBackups, String releaseFloatingIps, String releaseEIPs, String removeInstances, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/items/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (preserveVolumes != null) {
            queryParams.put("preserveVolumes", preserveVolumes)
        }
        if (keepBackups != null) {
            queryParams.put("keepBackups", keepBackups)
        }
        if (releaseFloatingIps != null) {
            queryParams.put("releaseFloatingIps", releaseFloatingIps)
        }
        if (releaseEIPs != null) {
            queryParams.put("releaseEIPs", releaseEIPs)
        }
        if (removeInstances != null) {
            queryParams.put("removeInstances", removeInstances)
        }
        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getCatalogItem ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/items/${id}"

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
                    InlineResponse200140.class )

    }

    def getCatalogType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/types/${id}"

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

    def listCatalogCart ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/cart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200139.class )

    }

    def listCatalogItems ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/items"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCatalogTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, Boolean featured, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog/types"

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
        if (featured != null) {
            queryParams.put("featured", featured)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

}
