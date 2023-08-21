package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject127
import org.openapitools.model.InlineObject128
import org.openapitools.model.InlineObject129
import org.openapitools.model.InlineObject130
import org.openapitools.model.InlineObject131
import org.openapitools.model.InlineObject132
import org.openapitools.model.InlineObject133
import org.openapitools.model.InlineObject134
import org.openapitools.model.InlineObject135
import org.openapitools.model.InlineObject136
import org.openapitools.model.InlineObject137
import org.openapitools.model.InlineObject138
import org.openapitools.model.InlineResponse20077
import org.openapitools.model.InlineResponse20078
import org.openapitools.model.InlineResponse20079
import org.openapitools.model.InlineResponse20080
import org.openapitools.model.InlineResponse20081
import org.openapitools.model.InlineResponse20082
import org.openapitools.model.InlineResponse20083
import org.openapitools.model.Model200Success

class LoadBalancersApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def createLoadBalancer ( InlineObject127 inlineObject127, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject127


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse20078.class )

    }

    def createLoadBalancerMonitor ( BigDecimal loadBalancerId, InlineObject129 inlineObject129, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/monitors"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject129


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def createLoadBalancerPool ( BigDecimal loadBalancerId, InlineObject131 inlineObject131, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/pools"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject131


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def createLoadBalancerPoolNode ( BigDecimal loadBalancerPoolId, InlineObject137 inlineObject137, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-pools/${loadBalancerPoolId}/nodes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerPoolId == null) {
            throw new RuntimeException("missing required params loadBalancerPoolId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject137


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def createLoadBalancerProfile ( BigDecimal loadBalancerId, InlineObject133 inlineObject133, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/profiles"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject133


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def createLoadBalancerVirtualServer ( BigDecimal loadBalancerId, InlineObject135 inlineObject135, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/virtual-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject135


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse20082.class )

    }

    def deleteLoadBalancer ( BigDecimal loadBalancerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteLoadBalancerMonitor ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/monitors/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteLoadBalancerPool ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/pools/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteLoadBalancerPoolNode ( BigDecimal loadBalancerPoolId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-pools/${loadBalancerPoolId}/nodes/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerPoolId == null) {
            throw new RuntimeException("missing required params loadBalancerPoolId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteLoadBalancerProfile ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/profiles/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteLoadBalancerVirtualServer ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/virtual-servers/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def getLoadBalancer ( BigDecimal loadBalancerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20078.class )

    }

    def getLoadBalancerMonitor ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/monitors/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20079.class )

    }

    def getLoadBalancerPool ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/pools/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20080.class )

    }

    def getLoadBalancerPoolNode ( BigDecimal loadBalancerPoolId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-pools/${loadBalancerPoolId}/nodes/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerPoolId == null) {
            throw new RuntimeException("missing required params loadBalancerPoolId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20083.class )

    }

    def getLoadBalancerProfile ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/profiles/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20081.class )

    }

    def getLoadBalancerType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-types/${id}"

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
                    InlineResponse20077.class )

    }

    def getLoadBalancerVirtualServer ( BigDecimal loadBalancerId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/virtual-servers/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20082.class )

    }

    def listLoadBalancerMonitors ( BigDecimal loadBalancerId, Long max, Long offset, String sort, String direction, String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/monitors"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
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
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLoadBalancerPoolNodes ( BigDecimal loadBalancerPoolId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-pools/${loadBalancerPoolId}/nodes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerPoolId == null) {
            throw new RuntimeException("missing required params loadBalancerPoolId")
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
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLoadBalancerPools ( BigDecimal loadBalancerId, Long max, Long offset, String sort, String direction, String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/pools"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
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
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLoadBalancerProfiles ( BigDecimal loadBalancerId, Long max, Long offset, String sort, String direction, String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/profiles"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
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
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLoadBalancerTypes ( Long max, Long offset, String sort, String direction, Boolean optionTypes, String phrase, String name, String code, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-types"

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
        if (optionTypes != null) {
            queryParams.put("optionTypes", optionTypes)
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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLoadBalancerVirtualServers ( BigDecimal loadBalancerId, Long max, Long offset, String sort, String direction, String phrase, String vipName, String vipAddress, String vipHostname, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/virtual-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
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
        if (direction != null) {
            queryParams.put("direction", direction)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (vipName != null) {
            queryParams.put("vipName", vipName)
        }
        if (vipAddress != null) {
            queryParams.put("vipAddress", vipAddress)
        }
        if (vipHostname != null) {
            queryParams.put("vipHostname", vipHostname)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLoadBalancers ( Long max, Long offset, String sort, String direction, String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers"

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
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def refreshLoadBalancer ( BigDecimal loadBalancerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/refresh"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateLoadBalancer ( BigDecimal loadBalancerId, InlineObject128 inlineObject128, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject128


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20078.class )

    }

    def updateLoadBalancerMonitor ( BigDecimal loadBalancerId, Long id, InlineObject130 inlineObject130, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/monitors/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject130


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateLoadBalancerPool ( BigDecimal loadBalancerId, Long id, InlineObject132 inlineObject132, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/pools/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject132


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateLoadBalancerPoolNode ( BigDecimal loadBalancerPoolId, Long id, InlineObject138 inlineObject138, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancer-pools/${loadBalancerPoolId}/nodes/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerPoolId == null) {
            throw new RuntimeException("missing required params loadBalancerPoolId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject138


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateLoadBalancerProfile ( BigDecimal loadBalancerId, Long id, InlineObject134 inlineObject134, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/profiles/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject134


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateLoadBalancerVirtualServer ( BigDecimal loadBalancerId, Long id, InlineObject136 inlineObject136, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/load-balancers/${loadBalancerId}/virtual-servers/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (loadBalancerId == null) {
            throw new RuntimeException("missing required params loadBalancerId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }



        contentType = 'application/json';
        bodyParams = inlineObject136


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse20082.class )

    }

}
