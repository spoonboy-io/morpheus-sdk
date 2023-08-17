package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.DefaultError
import org.openapitools.model.GetBillingInstancesIdentifier200Response
import org.openapitools.model.GetBillingServersIdentifier200Response
import org.openapitools.model.GetBillingZoneIdentifier200Response
import org.openapitools.model.ListBillingAccount200Response
import org.openapitools.model.ListBillingInstances200Response
import org.openapitools.model.ListBillingServers200Response
import org.openapitools.model.ListBillingZone200Response

class BillingApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getBillingAccount ( Long id, String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeComputeServers, Boolean includeInstances, Boolean includeDiscoveredServers, Boolean includeLoadBalancers, Boolean includeVirtualImages, Boolean includeSnapshots, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/account/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeComputeServers != null) {
            queryParams.put("includeComputeServers", includeComputeServers)
        }
        if (includeInstances != null) {
            queryParams.put("includeInstances", includeInstances)
        }
        if (includeDiscoveredServers != null) {
            queryParams.put("includeDiscoveredServers", includeDiscoveredServers)
        }
        if (includeLoadBalancers != null) {
            queryParams.put("includeLoadBalancers", includeLoadBalancers)
        }
        if (includeVirtualImages != null) {
            queryParams.put("includeVirtualImages", includeVirtualImages)
        }
        if (includeSnapshots != null) {
            queryParams.put("includeSnapshots", includeSnapshots)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBillingAccount200Response.class )

    }

    def getBillingInstancesIdentifier ( String identifier, String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeTenants, Long accountId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/instances/${identifier}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (identifier == null) {
            throw new RuntimeException("missing required params identifier")
        }

        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeTenants != null) {
            queryParams.put("includeTenants", includeTenants)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    GetBillingInstancesIdentifier200Response.class )

    }

    def getBillingServersIdentifier ( String identifier, String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeTenants, Long accountId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/servers/${identifier}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (identifier == null) {
            throw new RuntimeException("missing required params identifier")
        }

        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeTenants != null) {
            queryParams.put("includeTenants", includeTenants)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    GetBillingServersIdentifier200Response.class )

    }

    def getBillingZoneIdentifier ( String identifier, String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeComputeServers, Boolean includeInstances, Boolean includeDiscoveredServers, Boolean includeLoadBalancers, Boolean includeVirtualImages, Boolean includeSnapshots, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/zones/${identifier}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (identifier == null) {
            throw new RuntimeException("missing required params identifier")
        }

        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeComputeServers != null) {
            queryParams.put("includeComputeServers", includeComputeServers)
        }
        if (includeInstances != null) {
            queryParams.put("includeInstances", includeInstances)
        }
        if (includeDiscoveredServers != null) {
            queryParams.put("includeDiscoveredServers", includeDiscoveredServers)
        }
        if (includeLoadBalancers != null) {
            queryParams.put("includeLoadBalancers", includeLoadBalancers)
        }
        if (includeVirtualImages != null) {
            queryParams.put("includeVirtualImages", includeVirtualImages)
        }
        if (includeSnapshots != null) {
            queryParams.put("includeSnapshots", includeSnapshots)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    GetBillingZoneIdentifier200Response.class )

    }

    def listBillingAccount ( String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeComputeServers, Boolean includeInstances, Boolean includeDiscoveredServers, Boolean includeLoadBalancers, Boolean includeVirtualImages, Boolean includeSnapshots, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/account"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeComputeServers != null) {
            queryParams.put("includeComputeServers", includeComputeServers)
        }
        if (includeInstances != null) {
            queryParams.put("includeInstances", includeInstances)
        }
        if (includeDiscoveredServers != null) {
            queryParams.put("includeDiscoveredServers", includeDiscoveredServers)
        }
        if (includeLoadBalancers != null) {
            queryParams.put("includeLoadBalancers", includeLoadBalancers)
        }
        if (includeVirtualImages != null) {
            queryParams.put("includeVirtualImages", includeVirtualImages)
        }
        if (includeSnapshots != null) {
            queryParams.put("includeSnapshots", includeSnapshots)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBillingAccount200Response.class )

    }

    def listBillingInstances ( String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeTenants, Long accountId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/instances"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeTenants != null) {
            queryParams.put("includeTenants", includeTenants)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBillingInstances200Response.class )

    }

    def listBillingServers ( String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeTenants, Long accountId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeTenants != null) {
            queryParams.put("includeTenants", includeTenants)
        }
        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBillingServers200Response.class )

    }

    def listBillingZone ( String startDate, String endDate, Boolean includeUsages, Long maxUsages, Long offsetUsages, Boolean includeComputeServers, Boolean includeInstances, Boolean includeDiscoveredServers, Boolean includeLoadBalancers, Boolean includeVirtualImages, Boolean includeSnapshots, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/billing/zones"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (startDate != null) {
            queryParams.put("startDate", startDate)
        }
        if (endDate != null) {
            queryParams.put("endDate", endDate)
        }
        if (includeUsages != null) {
            queryParams.put("includeUsages", includeUsages)
        }
        if (maxUsages != null) {
            queryParams.put("maxUsages", maxUsages)
        }
        if (offsetUsages != null) {
            queryParams.put("offsetUsages", offsetUsages)
        }
        if (includeComputeServers != null) {
            queryParams.put("includeComputeServers", includeComputeServers)
        }
        if (includeInstances != null) {
            queryParams.put("includeInstances", includeInstances)
        }
        if (includeDiscoveredServers != null) {
            queryParams.put("includeDiscoveredServers", includeDiscoveredServers)
        }
        if (includeLoadBalancers != null) {
            queryParams.put("includeLoadBalancers", includeLoadBalancers)
        }
        if (includeVirtualImages != null) {
            queryParams.put("includeVirtualImages", includeVirtualImages)
        }
        if (includeSnapshots != null) {
            queryParams.put("includeSnapshots", includeSnapshots)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ListBillingZone200Response.class )

    }

}
