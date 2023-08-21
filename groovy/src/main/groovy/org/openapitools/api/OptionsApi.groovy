package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.Model200Success
import org.openapitools.model.ZoneNetworkOptions

class OptionsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getOptionSourceData ( String optionSource, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/options/${optionSource}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (optionSource == null) {
            throw new RuntimeException("missing required params optionSource")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listCodeRepositories ( Long integrationId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/options/codeRepositories"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (integrationId != null) {
            queryParams.put("integrationId", integrationId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Model200Success.class )

    }

    def listOptionNetworkOptions ( Long zoneId, Long provisionTypeId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/options/zoneNetworkOptions"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (provisionTypeId != null) {
            queryParams.put("provisionTypeId", provisionTypeId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ZoneNetworkOptions.class )

    }

    def listOptionValues ( Long optionTypeId, Object config, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/options/list"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (optionTypeId == null) {
            throw new RuntimeException("missing required params optionTypeId")
        }

        if (optionTypeId != null) {
            queryParams.put("optionTypeId", optionTypeId)
        }
        if (config != null) {
            queryParams.put("config", config)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

}
