package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject142
import org.openapitools.model.InlineObject143
import org.openapitools.model.InlineObject144
import org.openapitools.model.InlineObject145
import org.openapitools.model.InlineObject146
import org.openapitools.model.InlineObject147
import org.openapitools.model.InlineObject148
import org.openapitools.model.InlineObject149
import org.openapitools.model.InlineObject150
import org.openapitools.model.InlineObject151
import org.openapitools.model.InlineObject152
import org.openapitools.model.InlineObject153
import org.openapitools.model.InlineObject154
import org.openapitools.model.InlineObject155
import org.openapitools.model.InlineObject156
import org.openapitools.model.InlineObject157
import org.openapitools.model.InlineObject158
import org.openapitools.model.InlineObject159
import org.openapitools.model.InlineObject160
import org.openapitools.model.InlineObject161
import org.openapitools.model.InlineObject162
import org.openapitools.model.InlineObject163
import org.openapitools.model.InlineObject164
import org.openapitools.model.InlineObject165
import org.openapitools.model.InlineObject166
import org.openapitools.model.InlineObject167
import org.openapitools.model.InlineObject168
import org.openapitools.model.InlineObject169
import org.openapitools.model.InlineObject170
import org.openapitools.model.InlineObject171
import org.openapitools.model.InlineObject172
import org.openapitools.model.InlineObject173
import org.openapitools.model.InlineObject174
import org.openapitools.model.InlineObject175
import org.openapitools.model.InlineObject176
import org.openapitools.model.InlineObject177
import org.openapitools.model.InlineObject178
import org.openapitools.model.InlineObject179
import org.openapitools.model.InlineObject180
import org.openapitools.model.InlineObject181
import org.openapitools.model.InlineObject244
import org.openapitools.model.InlineObject245
import org.openapitools.model.InlineResponse200100
import org.openapitools.model.InlineResponse200101
import org.openapitools.model.InlineResponse200102
import org.openapitools.model.InlineResponse200103
import org.openapitools.model.InlineResponse200104
import org.openapitools.model.InlineResponse200105
import org.openapitools.model.InlineResponse200106
import org.openapitools.model.InlineResponse200107
import org.openapitools.model.InlineResponse200108
import org.openapitools.model.InlineResponse200109
import org.openapitools.model.InlineResponse200110
import org.openapitools.model.InlineResponse200111
import org.openapitools.model.InlineResponse200112
import org.openapitools.model.InlineResponse200113
import org.openapitools.model.InlineResponse200114
import org.openapitools.model.InlineResponse200115
import org.openapitools.model.InlineResponse200116
import org.openapitools.model.InlineResponse200117
import org.openapitools.model.InlineResponse200118
import org.openapitools.model.InlineResponse200119
import org.openapitools.model.InlineResponse200120
import org.openapitools.model.InlineResponse200121
import org.openapitools.model.InlineResponse200154
import org.openapitools.model.InlineResponse20086
import org.openapitools.model.InlineResponse20087
import org.openapitools.model.InlineResponse20088
import org.openapitools.model.InlineResponse20089
import org.openapitools.model.InlineResponse20090
import org.openapitools.model.InlineResponse20091
import org.openapitools.model.InlineResponse20092
import org.openapitools.model.InlineResponse20093
import org.openapitools.model.InlineResponse20094
import org.openapitools.model.InlineResponse20095
import org.openapitools.model.InlineResponse20096
import org.openapitools.model.InlineResponse20097
import org.openapitools.model.InlineResponse20098
import org.openapitools.model.InlineResponse20099
import org.openapitools.model.Model200Success
import org.openapitools.model.Model400Error
import org.openapitools.model.Model401Error
import org.openapitools.model.Model403Error
import org.openapitools.model.Model404Error
import org.openapitools.model.Model405Error
import org.openapitools.model.Model406Error
import org.openapitools.model.Model410Error
import org.openapitools.model.Model429Error
import org.openapitools.model.Model500Error
import org.openapitools.model.Model503Error
import org.openapitools.model.SuccessId

class NetworksApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def createNetworkDhcpRelay ( BigDecimal serverId, InlineObject167 inlineObject167, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-relays"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject167


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkDhcpServer ( BigDecimal serverId, InlineObject169 inlineObject169, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject169


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkDomain ( InlineObject163 inlineObject163, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/domains"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject163


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200109.class )

    }

    def createNetworkFirewallRule ( BigDecimal serverId, InlineObject172 inlineObject172, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject172


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkFirewallRuleGroup ( BigDecimal serverId, InlineObject174 inlineObject174, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rule-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject174


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkGroup ( InlineObject146 inlineObject146, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject146


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkPool ( InlineObject160 inlineObject160, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject160


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200106.class )

    }

    def createNetworkPoolIp ( Long id, InlineObject162 inlineObject162, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${id}/ips"

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
        bodyParams = inlineObject162


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200107.class )

    }

    def createNetworkPoolServer ( InlineObject180 inlineObject180, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject180


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def createNetworkProxy ( InlineObject165 inlineObject165, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/proxies"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject165


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200110.class )

    }

    def createNetworkRouter ( InlineObject148 inlineObject148, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject148


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkRouterBgpNeighbor ( BigDecimal routerId, InlineObject150 inlineObject150, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/bgp-neighbors"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject150


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkRouterFirewallRule ( BigDecimal routerId, InlineObject152 inlineObject152, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject152


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkRouterFirewallRuleGroup ( BigDecimal routerId, InlineObject154 inlineObject154, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rule-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject154


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkRouterNat ( BigDecimal routerId, InlineObject156 inlineObject156, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/nats"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject156


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkRouterRoute ( BigDecimal routerId, InlineObject159 inlineObject159, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/routes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject159


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkServerGroup ( BigDecimal serverId, InlineObject176 inlineObject176, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject176


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworkTransportZone ( BigDecimal serverId, InlineObject178 inlineObject178, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/scopes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject178


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def createNetworks ( InlineObject142 inlineObject142, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject142


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def createStaticRoute ( Long id, InlineObject144 inlineObject144, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}/routes"

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
        bodyParams = inlineObject144


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessId.class )

    }

    def createSubnet ( InlineObject244 inlineObject244, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/subnets"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject244


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    InlineResponse200154.class )

    }

    def deleteNetwork ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}"

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

    def deleteNetworkDhcpRelay ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-relays/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkDhcpServer ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-servers/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkDomain ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/domains/${id}"

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

    def deleteNetworkFirewallRule ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rules/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkFirewallRuleGroup ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rule-groups/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkGroup ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/groups/${id}"

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

    def deleteNetworkPool ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${id}"

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

    def deleteNetworkPoolIp ( Long networkPoolId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${networkPoolId}/ips/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (networkPoolId == null) {
            throw new RuntimeException("missing required params networkPoolId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkPoolServer ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-servers/${id}"

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

    def deleteNetworkProxy ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/proxies/${id}"

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

    def deleteNetworkRouter ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${id}"

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

    def deleteNetworkRouterBgpNeighbor ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/bgp-neighbors/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkRouterFirewallRule ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rules/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkRouterFirewallRuleGroup ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rule-groups/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkRouterNat ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/nats/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkRouterRoute ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/routes/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkServerGroup ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/groups/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteNetworkTransportZone ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/scopes/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteStaticRoute ( Long id, BigDecimal routeId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}/routes/${routeId}"

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
        if (routeId == null) {
            throw new RuntimeException("missing required params routeId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteSubnet ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/subnets/${id}"

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

    def getAllNetworkFloatingIps ( String phrase, String ipAddress, String ipStatus, Long zoneId, Long serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/floating-ips"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (ipAddress != null) {
            queryParams.put("ipAddress", ipAddress)
        }
        if (ipStatus != null) {
            queryParams.put("ipStatus", ipStatus)
        }
        if (zoneId != null) {
            queryParams.put("zoneId", zoneId)
        }
        if (serverId != null) {
            queryParams.put("serverId", serverId)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getNetwork ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}"

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
                    InlineResponse20087.class )

    }

    def getNetworkDhcpRelay ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-relays/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200112.class )

    }

    def getNetworkDhcpRelays ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-relays"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def getNetworkDhcpServer ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-servers/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200113.class )

    }

    def getNetworkDhcpServers ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-servers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def getNetworkDomain ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/domains/${id}"

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
                    InlineResponse200109.class )

    }

    def getNetworkDomains ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/domains"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getNetworkEdgeCluster ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/edge-clusters/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200114.class )

    }

    def getNetworkEdgeClusters ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/edge-clusters"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def getNetworkFirewallRule ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rules/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200115.class )

    }

    def getNetworkFirewallRuleGroup ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rule-groups/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200116.class )

    }

    def getNetworkFirewallRuleGroups ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rule-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def getNetworkFirewallRules ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def getNetworkFloatingIp ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/floating-ips/${id}"

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
                    InlineResponse200108.class )

    }

    def getNetworkGroup ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/groups/${id}"

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
                    InlineResponse20091.class )

    }

    def getNetworkGroups ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20090.class )

    }

    def getNetworkPool ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${id}"

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
                    InlineResponse200106.class )

    }

    def getNetworkPoolIp ( Long networkPoolId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${networkPoolId}/ips/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (networkPoolId == null) {
            throw new RuntimeException("missing required params networkPoolId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getNetworkPoolIps ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${id}/ips"

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

    def getNetworkPoolServer ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-servers/${id}"

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
                    InlineResponse200120.class )

    }

    def getNetworkPoolServerType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-server-types/${id}"

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
                    InlineResponse200121.class )

    }

    def getNetworkPools ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getNetworkProxies ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/proxies"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def getNetworkProxy ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/proxies/${id}"

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
                    InlineResponse200111.class )

    }

    def getNetworkRouter ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${id}"

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
                    InlineResponse20095.class )

    }

    def getNetworkRouterBgpNeighbor ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/bgp-neighbors/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20097.class )

    }

    def getNetworkRouterFirewallRule ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rules/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20099.class )

    }

    def getNetworkRouterFirewallRuleGroup ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rule-groups/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200101.class )

    }

    def getNetworkRouterFirewallRuleGroups ( BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rule-groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200100.class )

    }

    def getNetworkRouterNat ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/nats/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200103.class )

    }

    def getNetworkRouterRoute ( Long id, BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/routes/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200105.class )

    }

    def getNetworkRouterType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/network-router-types/${id}"

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
                    InlineResponse20093.class )

    }

    def getNetworkRouters ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20094.class )

    }

    def getNetworkRoutersBgpNeighbors ( BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/bgp-neighbors"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20096.class )

    }

    def getNetworkRoutersFirewallRules ( BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rules"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20098.class )

    }

    def getNetworkRoutersNats ( BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/nats"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200102.class )

    }

    def getNetworkRoutersRoutes ( BigDecimal routerId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/routes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200104.class )

    }

    def getNetworkServerGroup ( BigDecimal serverId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/groups/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200117.class )

    }

    def getNetworkSubnets ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}/subnets"

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

    def getNetworkTransportZone ( Long id, BigDecimal serverId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/scopes/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200118.class )

    }

    def getNetworkTransportZones ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/scopes"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def getNetworkType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/network-types/${id}"

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
                    InlineResponse20086.class )

    }

    def getStaticRoute ( Long id, BigDecimal routeId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}/routes/${routeId}"

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
        if (routeId == null) {
            throw new RuntimeException("missing required params routeId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20089.class )

    }

    def getStaticRoutes ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}/routes"

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
                    InlineResponse20088.class )

    }

    def getSubnet ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/subnets/${id}"

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
                    InlineResponse200154.class )

    }

    def listNetworkPoolServerTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-server-types"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listNetworkPoolServers ( Long max, Long offset, String sort, String direction, String phrase, String name, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-servers"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listNetworkRouterTypes ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/network-router-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20092.class )

    }

    def listNetworkServerGroups ( BigDecimal serverId, Long max, Long offset, String sort, String direction, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/groups"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
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

    def listNetworkServices ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/services"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200119.class )

    }

    def listNetworkTypes ( String name, String code, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/network-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType


        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listNetworks ( String name, String phrase, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks"

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

    def listSubnetTypes ( String name, String phrase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/subnet-types"

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




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listSubnets ( String name, String phrase, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/subnets"

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

    def refreshNetworkServer ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${id}/refresh"

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
                    "POST", "",
                    SuccessId.class )

    }

    def releaseNetworkFloatingIp ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/floating-ips/${id}/release"

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

    def updateNetwork ( Long id, InlineObject143 inlineObject143, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}"

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
        bodyParams = inlineObject143


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

    def updateNetworkDhcpRelay ( Long id, BigDecimal serverId, InlineObject168 inlineObject168, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-relays/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject168


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkDhcpServer ( Long id, BigDecimal serverId, InlineObject170 inlineObject170, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/dhcp-servers/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject170


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkDomain ( Long id, InlineObject164 inlineObject164, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/domains/${id}"

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
        bodyParams = inlineObject164


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200109.class )

    }

    def updateNetworkEdgeCluster ( Long id, BigDecimal serverId, InlineObject171 inlineObject171, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/edge-clusters/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject171


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkFirewallRule ( Long id, BigDecimal serverId, InlineObject173 inlineObject173, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rules/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject173


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkFirewallRuleGroup ( Long id, BigDecimal serverId, InlineObject175 inlineObject175, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/firewall-rule-groups/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject175


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkGroup ( Long id, InlineObject147 inlineObject147, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/groups/${id}"

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
        bodyParams = inlineObject147


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkPool ( Long id, InlineObject161 inlineObject161, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pools/${id}"

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
        bodyParams = inlineObject161


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200106.class )

    }

    def updateNetworkPoolServer ( Long id, InlineObject181 inlineObject181, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/pool-servers/${id}"

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
        bodyParams = inlineObject181


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkProxy ( Long id, InlineObject166 inlineObject166, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/proxies/${id}"

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
        bodyParams = inlineObject166


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkRouter ( Long id, InlineObject149 inlineObject149, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${id}"

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
        bodyParams = inlineObject149


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkRouterBgpNeighbor ( Long id, BigDecimal routerId, InlineObject151 inlineObject151, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/bgp-neighbors/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject151


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkRouterFirewallRule ( Long id, BigDecimal routerId, InlineObject153 inlineObject153, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rules/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject153


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkRouterFirewallRuleGroup ( Long id, BigDecimal routerId, InlineObject155 inlineObject155, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/firewall-rule-groups/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject155


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkRouterNat ( Long id, BigDecimal routerId, InlineObject157 inlineObject157, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/nats/${id}"

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
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject157


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkRouterPermissions ( BigDecimal routerId, InlineObject158 inlineObject158, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/routers/${routerId}/permissions"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (routerId == null) {
            throw new RuntimeException("missing required params routerId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject158


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessId.class )

    }

    def updateNetworkServerGroup ( Long id, BigDecimal serverId, InlineObject177 inlineObject177, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/groups/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject177


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNetworkTransportZone ( Long id, BigDecimal serverId, InlineObject179 inlineObject179, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/servers/${serverId}/scopes/${id}"

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
        if (serverId == null) {
            throw new RuntimeException("missing required params serverId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject179


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateStaticRoute ( Long id, BigDecimal routeId, InlineObject145 inlineObject145, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/networks/${id}/routes/${routeId}"

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
        if (routeId == null) {
            throw new RuntimeException("missing required params routeId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject145


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    SuccessId.class )

    }

    def updateSubnet ( Long id, InlineObject245 inlineObject245, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/subnets/${id}"

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
        bodyParams = inlineObject245


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    InlineResponse200154.class )

    }

}
