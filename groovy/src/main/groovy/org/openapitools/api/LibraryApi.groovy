package org.openapitools.api;

import org.openapitools.api.ApiUtils
import java.math.BigDecimal
import org.openapitools.model.DefaultError
import org.openapitools.model.InlineObject107
import org.openapitools.model.InlineObject108
import org.openapitools.model.InlineObject109
import org.openapitools.model.InlineObject110
import org.openapitools.model.InlineObject111
import org.openapitools.model.InlineObject112
import org.openapitools.model.InlineObject113
import org.openapitools.model.InlineObject114
import org.openapitools.model.InlineObject115
import org.openapitools.model.InlineObject117
import org.openapitools.model.InlineObject118
import org.openapitools.model.InlineObject119
import org.openapitools.model.InlineObject120
import org.openapitools.model.InlineObject121
import org.openapitools.model.InlineObject122
import org.openapitools.model.InlineObject123
import org.openapitools.model.InlineObject124
import org.openapitools.model.InlineObject263
import org.openapitools.model.InlineObject264
import org.openapitools.model.InlineResponse200136
import org.openapitools.model.InlineResponse200165
import org.openapitools.model.InlineResponse20068
import org.openapitools.model.InlineResponse20069
import org.openapitools.model.InlineResponse20070
import org.openapitools.model.InlineResponse20071
import org.openapitools.model.InlineResponse20072
import org.openapitools.model.InlineResponse20073
import org.openapitools.model.InlineResponse20074
import org.openapitools.model.InlineResponse20075
import org.openapitools.model.InlineResponse20076
import org.openapitools.model.Model200Success
import org.openapitools.model.ScriptSuccessId
import org.openapitools.model.SuccessId

class LibraryApi {
    String basePath = "https://CHANGEME"
    String versionPath = ""
    ApiUtils apiUtils = new ApiUtils();

    def addFileTemplate ( InlineObject111 inlineObject111, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-templates"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject111


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def addInstanceType ( InlineObject113 inlineObject113, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject113


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def addLayout ( BigDecimal instanceTypeId, InlineObject115 inlineObject115, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}/layouts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject115


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def addNodeType ( InlineObject109 inlineObject109, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject109


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addOptionList ( InlineObject119 inlineObject119, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-type-lists"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject119


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def addOptionType ( InlineObject121 inlineObject121, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject121


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def addScript ( InlineObject107 inlineObject107, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-scripts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject107


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    ScriptSuccessId.class )

    }

    def addSpecTemplate ( InlineObject123 inlineObject123, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/spec-templates"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject123


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    SuccessId.class )

    }

    def addVirtualImage ( InlineObject263 inlineObject263, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType




        contentType = 'application/json';
        bodyParams = inlineObject263


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Object.class )

    }

    def addVirtualImageFile ( BigDecimal virtualImageId, String filename, String url, File body, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}/upload"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }

        if (filename != null) {
            queryParams.put("filename", filename)
        }
        if (url != null) {
            queryParams.put("url", url)
        }


        contentType = 'application/octet-stream';
        bodyParams = body


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def deleteFileTemplate ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-templates/${id}"

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

    def deleteInstanceType ( BigDecimal instanceTypeId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def deleteLayout ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/layouts/${id}"

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

    def deleteNodeType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-types/${id}"

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

    def deleteOptionList ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-type-lists/${id}"

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

    def deleteOptionType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-types/${id}"

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

    def deleteScript ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-scripts/${id}"

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

    def deleteSpecTemplate ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/spec-templates/${id}"

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

    def getFileTemplate ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-templates/${id}"

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
                    InlineResponse20070.class )

    }

    def getInput ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-types/${id}"

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
                    InlineResponse20075.class )

    }

    def getInstanceType ( BigDecimal instanceTypeId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse20071.class )

    }

    def getLayout ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/layouts/${id}"

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
                    InlineResponse20072.class )

    }

    def getNodeType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-types/${id}"

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
                    InlineResponse20069.class )

    }

    def getOptionList ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-type-lists/${id}"

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
                    InlineResponse20073.class )

    }

    def getOptionListItems ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-type-lists/${id}/items"

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
                    InlineResponse20074.class )

    }

    def getScript ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-scripts/${id}"

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
                    InlineResponse20068.class )

    }

    def getSecurityPackageType ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-package-types/${id}"

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
                    InlineResponse200136.class )

    }

    def getSpecTemplate ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/spec-templates/${id}"

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
                    InlineResponse20076.class )

    }

    def getVirtualImage ( BigDecimal virtualImageId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    InlineResponse200165.class )

    }

    def listFileTemplates ( Long max, Long offset, String sort, String direction, String phrase, String name, String labels, String allLabels, String fileName, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-templates"

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
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (fileName != null) {
            queryParams.put("fileName", fileName)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listInputs ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, String labels, String allLabels, String fieldName, String fieldContext, String fieldLabel, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-types"

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
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (fieldName != null) {
            queryParams.put("fieldName", fieldName)
        }
        if (fieldContext != null) {
            queryParams.put("fieldContext", fieldContext)
        }
        if (fieldLabel != null) {
            queryParams.put("fieldLabel", fieldLabel)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listInstanceTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, Boolean featured, String labels, String allLabels, Boolean details, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types"

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
        if (featured != null) {
            queryParams.put("featured", featured)
        }
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (details != null) {
            queryParams.put("details", details)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listLayouts ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, String provisionType, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/layouts"

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
        if (provisionType != null) {
            queryParams.put("provisionType", provisionType)
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

    def listLayoutsForInstanceType ( BigDecimal instanceTypeId, Long max, Long offset, String sort, String direction, String phrase, String name, String code, String provisionType, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}/layouts"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (code != null) {
            queryParams.put("code", code)
        }
        if (provisionType != null) {
            queryParams.put("provisionType", provisionType)
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

    def listNodeTypes ( Long max, Long offset, String sort, String direction, String phrase, String name, String code, String provisionType, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-types"

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
        if (provisionType != null) {
            queryParams.put("provisionType", provisionType)
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

    def listOptionLists ( Long max, Long offset, String sort, String direction, String phrase, String name, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-type-lists"

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

    def listScripts ( Long max, Long offset, String sort, String direction, String phrase, String labels, String allLabels, String scriptType, String scriptPhase, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-scripts"

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
        if (labels != null) {
            queryParams.put("labels", labels)
        }
        if (allLabels != null) {
            queryParams.put("allLabels", allLabels)
        }
        if (scriptType != null) {
            queryParams.put("scriptType", scriptType)
        }
        if (scriptPhase != null) {
            queryParams.put("scriptPhase", scriptPhase)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listSecurityPackageTypes ( Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-package-types"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType






        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listSpecTemplates ( Long max, Long offset, String sort, String direction, String phrase, String name, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/spec-templates"

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

    def listVirtualImageLocations ( BigDecimal virtualImageId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}/locations"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "GET", "",
                    Object.class )

    }

    def listVirtualImages ( Long max, Long offset, String name, String phrase, Date lastUpdated, String filterType, String imageType, String tags, String labels, String allLabels, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images"

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
        if (name != null) {
            queryParams.put("name", name)
        }
        if (phrase != null) {
            queryParams.put("phrase", phrase)
        }
        if (lastUpdated != null) {
            queryParams.put("lastUpdated", lastUpdated)
        }
        if (filterType != null) {
            queryParams.put("filterType", filterType)
        }
        if (imageType != null) {
            queryParams.put("imageType", imageType)
        }
        if (tags != null) {
            queryParams.put("tags", tags)
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

    def removeSecurityScans ( Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/security-scans/${id}"

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

    def removeVirtualImage ( BigDecimal virtualImageId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeVirtualImageFile ( BigDecimal virtualImageId, String filename, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}/files"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }

        if (filename != null) {
            queryParams.put("filename", filename)
        }




        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def removeVirtualImageLocation ( BigDecimal virtualImageId, Long id, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}/locations/${id}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }
        // verify required params are set
        if (id == null) {
            throw new RuntimeException("missing required params id")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "DELETE", "",
                    Model200Success.class )

    }

    def setInstanceTypeFeatured ( BigDecimal instanceTypeId, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}/toggle-featured"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
        }





        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateFileTemplate ( Long id, InlineObject112 inlineObject112, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-templates/${id}"

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
        bodyParams = inlineObject112


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateInstanceType ( BigDecimal instanceTypeId, InlineObject114 inlineObject114, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject114


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateInstanceTypeLogo ( BigDecimal instanceTypeId, File logo, File darkLogo, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/instance-types/${instanceTypeId}/update-logo"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (instanceTypeId == null) {
            throw new RuntimeException("missing required params instanceTypeId")
        }




        contentType = 'multipart/form-data';
        bodyParams = [:]
        bodyParams.put("logo", logo)
        bodyParams.put("darkLogo", darkLogo)

        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "POST", "",
                    Model200Success.class )

    }

    def updateLayout ( Long id, InlineObject117 inlineObject117, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/layouts/${id}"

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
        bodyParams = inlineObject117


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateLayoutPermissions ( Long id, InlineObject118 inlineObject118, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/layouts/${id}/permissions"

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
        bodyParams = inlineObject118


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateNodeType ( Long id, InlineObject110 inlineObject110, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-types/${id}"

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
        bodyParams = inlineObject110


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateOptionList ( Long id, InlineObject120 inlineObject120, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-type-lists/${id}"

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
        bodyParams = inlineObject120


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateOptionType ( Long id, InlineObject122 inlineObject122, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/option-types/${id}"

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
        bodyParams = inlineObject122


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateScript ( Long id, InlineObject108 inlineObject108, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/container-scripts/${id}"

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
        bodyParams = inlineObject108


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    ScriptSuccessId.class )

    }

    def updateSpecTemplate ( Long id, InlineObject124 inlineObject124, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/library/spec-templates/${id}"

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
        bodyParams = inlineObject124


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Model200Success.class )

    }

    def updateVirtualImage ( BigDecimal virtualImageId, InlineObject264 inlineObject264, Closure onSuccess, Closure onFailure)  {
        String resourcePath = "/api/virtual-images/${virtualImageId}"

        // params
        def queryParams = [:]
        def headerParams = [:]
        def bodyParams
        def contentType

        // verify required params are set
        if (virtualImageId == null) {
            throw new RuntimeException("missing required params virtualImageId")
        }



        contentType = 'application/json';
        bodyParams = inlineObject264


        apiUtils.invokeApi(onSuccess, onFailure, basePath, versionPath, resourcePath, queryParams, headerParams, bodyParams, contentType,
                    "PUT", "",
                    Object.class )

    }

}
