package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject80
import org.openapitools.model.InlineObject81
import org.openapitools.model.Model200Success
import org.openapitools.model.Ping

class HealthApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def acknowledgeHealthAlarm ( Long id, InlineObject81 inlineObject81, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health/alarms/${id}/acknowledge"

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
        bodyParams = inlineObject81


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def acknowledgeHealthAlarms ( InlineObject80 inlineObject80, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health/alarms/acknowledge"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject80


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def exportHealthLogs ( Long max, Long offset, String sort, String direction, String phrase, String name, Boolean acknowledged, String startDate, String endDate, Boolean reverse, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health/logs/export"

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
        if (acknowledged != null) {
            queryParams.put("acknowledged", acknowledged)
        }
        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (reverse != null) {
            queryParams.put("reverse", reverse)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    File.class )

    }

    def getHealthAlarms ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health/alarms/${id}"

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

    def listHealth ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listHealthAlarms ( Long max, Long offset, String sort, String direction, String phrase, String name, Boolean acknowledged, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health/alarms"

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
        if (acknowledged != null) {
            queryParams.put("acknowledged", acknowledged)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listHealthLogs ( Long max, Long offset, String sort, String direction, String phrase, String name, Boolean acknowledged, String startDate, String endDate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/health/logs"

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
        if (acknowledged != null) {
            queryParams.put("acknowledged", acknowledged)
        }
        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def ping ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/ping"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Ping.class )

    }

}
