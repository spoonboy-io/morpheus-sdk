package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.Log

class LogsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def listLogs ( Long max, Long offset, String sort, String order, String query, String message, String sourceType, String typeCode, Long objectId, String token, String level, Long startMs, Long endMs, Date startDateTime, Date endDateTime, Long containers, Long servers, Long clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/logs"

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
        if (order != null) {
            queryParams.put("order", order)
        }
        if (query != null) {
            queryParams.put("query", query)
        }
        if (message != null) {
            queryParams.put("message", message)
        }
        if (sourceType != null) {
            queryParams.put("sourceType", sourceType)
        }
        if (typeCode != null) {
            queryParams.put("typeCode", typeCode)
        }
        if (objectId != null) {
            queryParams.put("objectId", objectId)
        }
        if (token != null) {
            queryParams.put("token", token)
        }
        if (level != null) {
            queryParams.put("level", level)
        }
        if (startMs != null) {
            queryParams.put("startMs", startMs)
        }
        if (endMs != null) {
            queryParams.put("endMs", endMs)
        }
        if (startDateTime != null) {
            queryParams.put("startDateTime", startDateTime)
        }
        if (endDateTime != null) {
            queryParams.put("endDateTime", endDateTime)
        }
        if (containers != null) {
            queryParams.put("containers", containers)
        }
        if (servers != null) {
            queryParams.put("servers", servers)
        }
        if (clusterId != null) {
            queryParams.put("clusterId", clusterId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Log.class )

    }

}
