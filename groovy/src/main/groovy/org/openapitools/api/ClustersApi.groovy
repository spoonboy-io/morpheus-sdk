package org.openapitools.api;

import org.openapitools.api.ApiUtils
import org.openapitools.model.ClusterApiConfig
import org.openapitools.model.ClusterApplyTemplate
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject268
import org.openapitools.model.InlineObject52
import org.openapitools.model.InlineObject53
import org.openapitools.model.InlineObject54
import org.openapitools.model.InlineObject55
import org.openapitools.model.InlineObject56
import org.openapitools.model.InlineObject57
import org.openapitools.model.InlineObject58
import org.openapitools.model.InlineObject59
import org.openapitools.model.InlineResponse200168
import org.openapitools.model.InlineResponse20026
import org.openapitools.model.InlineResponse20027
import org.openapitools.model.InlineResponse20028
import org.openapitools.model.InlineResponse20029
import org.openapitools.model.InlineResponse20030
import org.openapitools.model.InlineResponse20031
import org.openapitools.model.InlineResponse20032
import org.openapitools.model.Model200Success
import org.openapitools.model.SuccessError

class ClustersApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addCluster ( InlineObject52 inlineObject52, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject52


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addClusterNamespace ( Integer clusterId, InlineObject56 inlineObject56, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/namespaces"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject56


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addClusterWorker ( Integer clusterId, InlineObject59 inlineObject59, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject59


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def applyTemplate ( Integer clusterId, InlineObject54 inlineObject54, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/apply-template"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject54


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    ClusterApplyTemplate.class )

    }

    def deleteCluster ( Integer clusterId, String removeInstances, String removeResources, String preserveVolumes, String releaseFloatingIps, String releaseEIPs, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

        if (removeInstances != null) {
            queryParams.put("removeInstances", removeInstances)
        }
        if (removeResources != null) {
            queryParams.put("removeResources", removeResources)
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

    def deleteClusterContainer ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/containers/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteClusterDeployment ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/deployments/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteClusterJob ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/jobs/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteClusterNamespace ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/namespaces/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteClusterService ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/services/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteClusterStatefulSet ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/statefulsets/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    SuccessError.class )

    }

    def deleteClusterVolume ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/volumes/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteClusterWorker ( Integer clusterId, Long id, String force, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/servers/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }

        if (force != null) {
            queryParams.put("force", force)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getCluster ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20026.class )

    }

    def getClusterApiConfig ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/api-config"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    ClusterApiConfig.class )

    }

    def getClusterDatastore ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/datastores/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20027.class )

    }

    def getClusterHistory ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/history"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getClusterHistoryDetail ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/history/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20028.class )

    }

    def getClusterHistoryEventDetail ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/history/events/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20029.class )

    }

    def getClusterMasters ( Integer clusterId, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/masters"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20030.class )

    }

    def getClusterNamespace ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/namespaces/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20031.class )

    }

    def getClusterNamespaces ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/namespaces"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getClusterUpgradeVersions ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/upgrade-cluster"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20032.class )

    }

    def getWikiCluster ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200168.class )

    }

    def listClusterContainers ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, String resourceLevel, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/containers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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
        if (resourceLevel != null) {
            queryParams.put("resourceLevel", resourceLevel)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterDatastores ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, String name, String code, Boolean hideInactive, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/datastores"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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
        if (code != null) {
            queryParams.put("code", code)
        }
        if (hideInactive != null) {
            queryParams.put("hideInactive", hideInactive)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterDeployments ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, String resourceLevel, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/deployments"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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
        if (resourceLevel != null) {
            queryParams.put("resourceLevel", resourceLevel)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterJobs ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/jobs"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterPods ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, String resourceLevel, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/pods"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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
        if (resourceLevel != null) {
            queryParams.put("resourceLevel", resourceLevel)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterServices ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/services"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterStatefulSets ( Integer clusterId, Long max, Long offset, String sort, String order, String phrase, String resourceLevel, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/statefulsets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }

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
        if (resourceLevel != null) {
            queryParams.put("resourceLevel", resourceLevel)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, String providerType, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/cluster-types"

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
        if (providerType != null) {
            queryParams.put("providerType", providerType)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterVolumes ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/volumes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusterWorkers ( Integer clusterId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/workers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listClusters ( Long max, Long offset, String sort, String direction, String phrase, String name, Long zoneId, Long typeId, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters"

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
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (typeId != null) {
            queryParams.put("typeId", typeId)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def restartClusterContainer ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/containers/${id}/restart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessError.class )

    }

    def restartClusterDeployment ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/deployments/${id}/restart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessError.class )

    }

    def restartClusterPod ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/pods/${id}/restart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessError.class )

    }

    def restartClusterStatefulSet ( Integer clusterId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/statefulsets/${id}/restart"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessError.class )

    }

    def updateCluster ( Integer clusterId, InlineObject53 inlineObject53, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject53


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateClusterDatastore ( Integer clusterId, Long id, InlineObject55 inlineObject55, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/datastores/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject55


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateClusterNamespace ( Integer clusterId, Long id, InlineObject57 inlineObject57, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/namespaces/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject57


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateClusterPermissions ( Integer clusterId, InlineObject58 inlineObject58, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/permissions"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject58


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateClusterUpgradeVersions ( Integer clusterId, String targetVersion, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/upgrade-cluster"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (targetVersion == null) {
            throw new RuntimeException("missing required params targetVersion")
        }

        if (targetVersion != null) {
            queryParams.put("targetVersion", targetVersion)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateClusterWorkerCount ( Integer clusterId, Long workerCount, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/worker-count"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }
        // verify required params are set
        if (workerCount == null) {
            throw new RuntimeException("missing required params workerCount")
        }

        if (workerCount != null) {
            queryParams.put("workerCount", workerCount)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateWikiCluster ( Integer clusterId, InlineObject268 inlineObject268, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/clusters/${clusterId}/wiki"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (clusterId == null) {
            throw new RuntimeException("missing required params clusterId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject268


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
