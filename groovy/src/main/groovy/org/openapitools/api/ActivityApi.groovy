package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.ListActivity200Response

class ActivityApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def listActivity ( Long max, Long offset, String sort, String order, String phrase, String name, Long userId, BigDecimal tenantId, String timeframe, Date start, Date end, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/activity"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (userId != null) {
            queryParams.put("userId", userId)
        }
        if (tenantId != null) {
            queryParams.put("tenantId", tenantId)
        }
        if (timeframe != null) {
            queryParams.put("timeframe", timeframe)
        }
        if (start != null) {
            queryParams.put("start", start)
        }
        if (end != null) {
            queryParams.put("end", end)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListActivity200Response.class )

    }

}
