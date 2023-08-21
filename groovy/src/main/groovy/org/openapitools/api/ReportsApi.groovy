package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject205
import org.openapitools.model.InlineResponse200129
import org.openapitools.model.InlineResponse200130
import org.openapitools.model.Model200Success

class ReportsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def downloadReports ( Long id, String format, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/reports/download/${id}${format}"

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
        if (format == null) {
            throw new RuntimeException("missing required params format")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Model200Success.class )

    }

    def getReportTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, String category, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/report-types"

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
        if (category != null) {
            queryParams.put("category", category)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200129.class )

    }

    def getReports ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/reports/${id}"

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
                    InlineResponse200130.class )

    }

    def listReports ( Long max, Long offset, String sort, String direction, String phrase, String name, String reportType, String category, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/reports"

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
        if (reportType != null) {
            queryParams.put("reportType", reportType)
        }
        if (category != null) {
            queryParams.put("category", category)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def removeReports ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/reports/${id}"

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

    def runReports ( InlineObject205 inlineObject205, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/reports"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject205


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

}
