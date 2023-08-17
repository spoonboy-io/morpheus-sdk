package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.AddCatalogItemType200Response
import org.openapitools.model.AddCatalogItemTypeRequest
import org.openapitools.model.AddCatalogItemTypeRequest1
import org.openapitools.model.DefaultError
import org.openapitools.model.GetCatalogItemType200Response
import org.openapitools.model.ListCatalogItemTypes200Response
import org.openapitools.model.Model200Success
import org.openapitools.model.UpdateCatalogItemType200Response
import org.openapitools.model.UpdateCatalogItemTypeRequest

class CatalogItemsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCatalogItemType ( AddCatalogItemTypeRequest addCatalogItemTypeRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog-item-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = addCatalogItemTypeRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    AddCatalogItemType200Response.class )

    }

    def getCatalogItemType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog-item-types/${id}"

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
                    GetCatalogItemType200Response.class )

    }

    def listCatalogItemTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String description, Boolean enabled, Boolean featured, String labels, String allLabels, String code, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog-item-types"

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
        if (description != null) {
            queryParams.put("description", description)
        }
        if (enabled != null) {
            queryParams.put("enabled", enabled)
        }
        if (featured != null) {
            queryParams.put("featured", featured)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (code != null) {
            queryParams.put("code", code)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListCatalogItemTypes200Response.class )

    }

    def removeCatalogItemType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog-item-types/${id}"

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

    def updateCatalogItemType ( Long id, UpdateCatalogItemTypeRequest updateCatalogItemTypeRequest, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog-item-types/${id}"

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
        bodyParams = updateCatalogItemTypeRequest


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    UpdateCatalogItemType200Response.class )

    }

    def updateCatalogItemTypeLogo ( Long id, File catalogItemTypeLogo, File catalogItemTypeDarkLogo, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/catalog-item-types/${id}/update-logo"

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
        bodyParams.put("catalogItemType.logo", catalogItemTypeLogo)
        bodyParams.put("catalogItemType.darkLogo", catalogItemTypeDarkLogo)

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    UpdateCatalogItemType200Response.class )

    }

}
