package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject200
import org.openapitools.model.InlineObject201
import org.openapitools.model.InlineResponse200124
import org.openapitools.model.Model200Success

class PricesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addPrices ( InlineObject200 inlineObject200, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/prices"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject200


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def deactivatePrices ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/prices/${id}/deactivate"

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
                    "PUT", "",
                    Model200Success.class )

    }

    def getPrices ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/prices/${id}"

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
                    InlineResponse200124.class )

    }

    def listPrices ( Long max, Long offset, String sort, String direction, String phrase, String name, Boolean includeInactive, String priceType, String platform, String currency, String type, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/prices"

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
        if (includeInactive != null) {
            queryParams.put("includeInactive", includeInactive)
        }
        if (priceType != null) {
            queryParams.put("priceType", priceType)
        }
        if (platform != null) {
            queryParams.put("platform", platform)
        }
        if (currency != null) {
            queryParams.put("currency", currency)
        }
        if (type != null) {
            queryParams.put("type", type)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updatePrices ( Long id, InlineObject201 inlineObject201, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/prices/${id}"

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
        bodyParams = inlineObject201


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
