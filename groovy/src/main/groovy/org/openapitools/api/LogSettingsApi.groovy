package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject139
import org.openapitools.model.InlineObject140
import org.openapitools.model.InlineResponse20084
import org.openapitools.model.Model200Success

class LogSettingsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addLogSettingsSyslogRules ( InlineObject140 inlineObject140, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/log-settings/syslog-rules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject140


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def deleteLogSettingsSyslogRules ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/log-settings/syslog-rules/${id}"

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

    def listLogSettings ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/log-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20084.class )

    }

    def updateLogSettings ( InlineObject139 inlineObject139, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/log-settings"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject139


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

}
