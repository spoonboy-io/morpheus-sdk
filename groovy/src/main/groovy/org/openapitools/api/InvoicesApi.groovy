package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject102

class InvoicesApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getInvoiceLineItems ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/invoice-line-items/${id}"

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

    def getInvoices ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/invoices/${id}"

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

    def listInvoiceLineItems ( Long max, Long offset, String sort, String direction, String phrase, String name, String startDate, String endDate, String period, String refType, Long refId, Long zoneId, Long siteId, Long instanceId, Long containerId, Long serverId, Long userId, Long projectId, Boolean active, Long accountId, Boolean includeTotals, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/invoice-line-items"

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
        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (period != null) {
            queryParams.put("period", period)
        }
        if (refType != null) {
            queryParams.put("refType", refType)
        }
        if (refId != null) {
            queryParams.put("refId", refId)
        }
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (siteId != null) {
            queryParams.put("siteId", siteId)
        }
        if (instanceId != null) {
            queryParams.put("instanceId", instanceId)
        }
        if (containerId != null) {
            queryParams.put("containerId", containerId)
        }
        if (serverId != null) {
            queryParams.put("serverId", serverId)
        }
        if (userId != null) {
            queryParams.put("userId", userId)
        }
        if (projectId != null) {
            queryParams.put("projectId", projectId)
        }
        if (active != null) {
            queryParams.put("active", active)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }
        if (includeTotals != null) {
            queryParams.put("includeTotals", includeTotals)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listInvoices ( Long max, Long offset, String sort, String direction, String phrase, String name, String startDate, String endDate, String period, String refType, Long refId, String refStatus, Long zoneId, Long siteId, Long instanceId, Long containerId, Long serverId, Long userId, Long projectId, Boolean active, Long accountId, Boolean includeLineItems, Boolean includeTotals, String tags, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/invoices"

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
        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (period != null) {
            queryParams.put("period", period)
        }
        if (refType != null) {
            queryParams.put("refType", refType)
        }
        if (refId != null) {
            queryParams.put("refId", refId)
        }
        if (refStatus != null) {
            queryParams.put("refStatus", refStatus)
        }
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (siteId != null) {
            queryParams.put("siteId", siteId)
        }
        if (instanceId != null) {
            queryParams.put("instanceId", instanceId)
        }
        if (containerId != null) {
            queryParams.put("containerId", containerId)
        }
        if (serverId != null) {
            queryParams.put("serverId", serverId)
        }
        if (userId != null) {
            queryParams.put("userId", userId)
        }
        if (projectId != null) {
            queryParams.put("projectId", projectId)
        }
        if (active != null) {
            queryParams.put("active", active)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }
        if (includeLineItems != null) {
            queryParams.put("includeLineItems", includeLineItems)
        }
        if (includeTotals != null) {
            queryParams.put("includeTotals", includeTotals)
        }
        if (tags != null) {
            queryParams.put("tags", tags)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def updateInvoices ( Long id, InlineObject102 inlineObject102, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/invoices/${id}"

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
        bodyParams = inlineObject102


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
