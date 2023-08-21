package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject265
import org.openapitools.model.InlineResponse200166
import org.openapitools.model.Model200Success

class WhitelabelSettingsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getWhitelabelImage ( String imageType, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/whitelabel-settings/images/${imageType}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (imageType == null) {
            throw new RuntimeException("missing required params imageType")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    File.class )

    }

    def listWhitelabelSettings ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/whitelabel-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200166.class )

    }

    def removeWhitelabelImage ( String imageType, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/whitelabel-settings/images/${imageType}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (imageType == null) {
            throw new RuntimeException("missing required params imageType")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def updateWhitelabelImages ( File headerLogoFile, Boolean resetHeaderLogo, File footerLogoFile, Boolean resetFooterLogo, File loginLogoFile, Boolean resetLoginLogo, File faviconFile, Boolean resetFaviconLogo, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/whitelabel-settings/images"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType





        contentType = 'multipart/form-data';
        bodyParams = [:]
        bodyParams.put("headerLogo.file", headerLogoFile)
        bodyParams.put("resetHeaderLogo", resetHeaderLogo)
        bodyParams.put("footerLogo.file", footerLogoFile)
        bodyParams.put("resetFooterLogo", resetFooterLogo)
        bodyParams.put("loginLogo.file", loginLogoFile)
        bodyParams.put("resetLoginLogo", resetLoginLogo)
        bodyParams.put("favicon.file", faviconFile)
        bodyParams.put("resetFaviconLogo", resetFaviconLogo)

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateWhitelabelSettings ( InlineObject265 inlineObject265, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/whitelabel-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject265


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
