package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject220
import org.openapitools.model.InlineObject221
import org.openapitools.model.InlineObject222
import org.openapitools.model.InlineObject223
import org.openapitools.model.InlineObject224
import org.openapitools.model.InlineObject225
import org.openapitools.model.InlineObject226
import org.openapitools.model.InlineObject271
import org.openapitools.model.InlineResponse200137
import org.openapitools.model.InlineResponse200138
import org.openapitools.model.InlineResponse200141
import org.openapitools.model.InlineResponse200168
import org.openapitools.model.InlineResponse20050
import org.openapitools.model.Model200Success
import org.openapitools.model.NetworkInterfaceUpdate

class HostsApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def getHost ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}"

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
                    InlineResponse200137.class )

    }

    def getHostSnpshots ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/snapshots"

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
                    InlineResponse200138.class )

    }

    def getHostType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/server-types/${id}"

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
                    InlineResponse20050.class )

    }

    def getWikiServer ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/wiki"

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
                    InlineResponse200168.class )

    }

    def listHostTypes ( Long max, Long offset, String sort, String direction, String name, String code, String phrase, String provisionType, String zoneType, Boolean creatable, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/server-types"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (provisionType != null) {
            queryParams.put("provisionType", provisionType)
        }
        if (zoneType != null) {
            queryParams.put("zoneType", zoneType)
        }
        if (creatable != null) {
            queryParams.put("creatable", creatable)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listHosts ( String name, String phrase, Long zoneId, Long siteId, Long clusterId, Boolean managed, String serverType, String powerState, String ip, Boolean vm, Boolean vmHypervisor, Boolean bareMetalHost, String status, Boolean agentInstalled, Long max, Long offset, Date lastUpdated, Long createdBy, String labels, String allLabels, String tags, String metadata, String uuid, String externalId, String internalId, String externalUniquelId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (siteId != null) {
            queryParams.put("siteId", siteId)
        }
        if (clusterId != null) {
            queryParams.put("clusterId", clusterId)
        }
        if (managed != null) {
            queryParams.put("managed", managed)
        }
        if (serverType != null) {
            queryParams.put("serverType", serverType)
        }
        if (powerState != null) {
            queryParams.put("powerState", powerState)
        }
        if (ip != null) {
            queryParams.put("ip", ip)
        }
        if (vm != null) {
            queryParams.put("vm", vm)
        }
        if (vmHypervisor != null) {
            queryParams.put("vmHypervisor", vmHypervisor)
        }
        if (bareMetalHost != null) {
            queryParams.put("bareMetalHost", bareMetalHost)
        }
        if (status != null) {
            queryParams.put("status", status)
        }
        if (agentInstalled != null) {
            queryParams.put("agentInstalled", agentInstalled)
        }
        if (max != null) {
            queryParams.put("max", max)
        }
        if (offset != null) {
            queryParams.put("offset", offset)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (createdBy != null) {
            queryParams.put("createdBy", createdBy)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (tags != null) {
            queryParams.put("tags", tags)
        }
        if (metadata != null) {
            queryParams.put("metadata", metadata)
        }
        if (uuid != null) {
            queryParams.put("uuid", uuid)
        }
        if (externalId != null) {
            queryParams.put("externalId", externalId)
        }
        if (internalId != null) {
            queryParams.put("internalId", internalId)
        }
        if (externalUniquelId != null) {
            queryParams.put("externalUniquelId", externalUniquelId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listServerServicePlans ( Long zoneId, Long serverTypeId, Long siteId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/service-plans"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (zoneId == null) {
            throw new RuntimeException("missing required params zoneId")
        }

        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (serverTypeId != null) {
            queryParams.put("serverTypeId", serverTypeId)
        }
        if (siteId != null) {
            queryParams.put("siteId", siteId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200141.class )

    }

    def removeHost ( Long id, String removeResources, String removeInstances, String preserveVolumes, String releaseFloatingIps, String releaseEIPs, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (removeResources != null) {
            queryParams.put("removeResources", removeResources)
        }
        if (removeInstances != null) {
            queryParams.put("removeInstances", removeInstances)
        }
        if (preserveVolumes != null) {
            queryParams.put("preserveVolumes", preserveVolumes)
        }
        if (releaseFloatingIps != null) {
            queryParams.put("releaseFloatingIps", releaseFloatingIps)
        }
        if (releaseEIPs != null) {
            queryParams.put("releaseEIPs", releaseEIPs)
        }
        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def restartHost ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/restart"

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
                    Object.class )

    }

    def startHost ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/start"

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

    def stopHost ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/stop"

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

    def updateHost ( Long id, InlineObject220 inlineObject220, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}"

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
        bodyParams = inlineObject220


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200137.class )

    }

    def updateHostAssignTenant ( Long id, Long accountId, InlineObject221 inlineObject221, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/assign-account"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (accountId != null) {
            queryParams.put("accountId", accountId)
        }


        contentType = 'application/json';
        bodyParams = inlineObject221


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateHostCloud ( InlineObject226 inlineObject226, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/change-cloud"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject226


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateHostExecuteWorkflow ( Long id, Long workflowId, String workflowName, InlineObject225 inlineObject225, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/workflow"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (workflowId != null) {
            queryParams.put("workflowId", workflowId)
        }
        if (workflowName != null) {
            queryParams.put("workflowName", workflowName)
        }


        contentType = 'application/json';
        bodyParams = inlineObject225


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateHostInstallAgent ( Long id, InlineObject222 inlineObject222, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/install-agent"

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
        bodyParams = inlineObject222


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateHostManaged ( Long id, InlineObject223 inlineObject223, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/make-managed"

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
        bodyParams = inlineObject223


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateHostResize ( Long id, InlineObject224 inlineObject224, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/resize"

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
        bodyParams = inlineObject224


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateHostUpgradeAgent ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/upgrade"

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

    def updateServerNetworkInterface ( Long id, BigDecimal networkInterfaceId, NetworkInterfaceUpdate networkInterfaceUpdate, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/networkInterfaces/${networkInterfaceId}"

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
        if (networkInterfaceId == null) {
            throw new RuntimeException("missing required params networkInterfaceId")
        }



        contentType = 'application/json';
        bodyParams = networkInterfaceUpdate


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateWikiServer ( Long id, InlineObject271 inlineObject271, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/servers/${id}/wiki"

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
        bodyParams = inlineObject271


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
