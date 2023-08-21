/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import OptionTypeListCreateConfig from './OptionTypeListCreateConfig';
import OptionTypeListCreateCredential from './OptionTypeListCreateCredential';

/**
 * The OptionTypeListCreate model module.
 * @module model/OptionTypeListCreate
 * @version 6.2.1
 */
class OptionTypeListCreate {
    /**
     * Constructs a new <code>OptionTypeListCreate</code>.
     * @alias module:model/OptionTypeListCreate
     * @param name {String} Name
     */
    constructor(name) { 
        
        OptionTypeListCreate.initialize(this, name);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj, name) { 
        obj['name'] = name;
    }

    /**
     * Constructs a <code>OptionTypeListCreate</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/OptionTypeListCreate} obj Optional instance to populate.
     * @return {module:model/OptionTypeListCreate} The populated <code>OptionTypeListCreate</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new OptionTypeListCreate();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('labels')) {
                obj['labels'] = ApiClient.convertToType(data['labels'], ['String']);
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('sourceUrl')) {
                obj['sourceUrl'] = ApiClient.convertToType(data['sourceUrl'], 'String');
            }
            if (data.hasOwnProperty('visibility')) {
                obj['visibility'] = ApiClient.convertToType(data['visibility'], 'String');
            }
            if (data.hasOwnProperty('sourceMethod')) {
                obj['sourceMethod'] = ApiClient.convertToType(data['sourceMethod'], 'String');
            }
            if (data.hasOwnProperty('apiType')) {
                obj['apiType'] = ApiClient.convertToType(data['apiType'], 'String');
            }
            if (data.hasOwnProperty('ignoreSSLErrors')) {
                obj['ignoreSSLErrors'] = ApiClient.convertToType(data['ignoreSSLErrors'], 'Boolean');
            }
            if (data.hasOwnProperty('realTime')) {
                obj['realTime'] = ApiClient.convertToType(data['realTime'], 'Boolean');
            }
            if (data.hasOwnProperty('credential')) {
                obj['credential'] = OptionTypeListCreateCredential.constructFromObject(data['credential']);
            }
            if (data.hasOwnProperty('serviceUsername')) {
                obj['serviceUsername'] = ApiClient.convertToType(data['serviceUsername'], 'String');
            }
            if (data.hasOwnProperty('servicePassword')) {
                obj['servicePassword'] = ApiClient.convertToType(data['servicePassword'], 'String');
            }
            if (data.hasOwnProperty('initialDataset')) {
                obj['initialDataset'] = ApiClient.convertToType(data['initialDataset'], 'String');
            }
            if (data.hasOwnProperty('translationScript')) {
                obj['translationScript'] = ApiClient.convertToType(data['translationScript'], 'String');
            }
            if (data.hasOwnProperty('requestScript')) {
                obj['requestScript'] = ApiClient.convertToType(data['requestScript'], 'String');
            }
            if (data.hasOwnProperty('config')) {
                obj['config'] = OptionTypeListCreateConfig.constructFromObject(data['config']);
            }
        }
        return obj;
    }


}

/**
 * Name
 * @member {String} name
 */
OptionTypeListCreate.prototype['name'] = undefined;

/**
 * Description
 * @member {String} description
 */
OptionTypeListCreate.prototype['description'] = undefined;

/**
 * Array of label strings, can be used for filtering.
 * @member {Array.<String>} labels
 */
OptionTypeListCreate.prototype['labels'] = undefined;

/**
 * Option List Type eg. `rest`, `api`, `ldap` or `manual`.
 * @member {module:model/OptionTypeListCreate.TypeEnum} type
 * @default 'rest'
 */
OptionTypeListCreate.prototype['type'] = 'rest';

/**
 * Source URL, the http(s) URL to request data from. Required when type is rest.
 * @member {String} sourceUrl
 */
OptionTypeListCreate.prototype['sourceUrl'] = undefined;

/**
 * Visibility
 * @member {module:model/OptionTypeListCreate.VisibilityEnum} visibility
 * @default 'private'
 */
OptionTypeListCreate.prototype['visibility'] = 'private';

/**
 * Source Method, the HTTP method.
 * @member {module:model/OptionTypeListCreate.SourceMethodEnum} sourceMethod
 * @default 'GET'
 */
OptionTypeListCreate.prototype['sourceMethod'] = 'GET';

/**
 * Api Type, The code of the api option list to use, eg. clouds, environments, groups, instances, instance-wiki, networks, servicePlans, resourcePools, securityGroups, servers, server-wiki. Required when type is api.
 * @member {String} apiType
 */
OptionTypeListCreate.prototype['apiType'] = undefined;

/**
 * Ignore SSL Errors.
 * @member {Boolean} ignoreSSLErrors
 * @default false
 */
OptionTypeListCreate.prototype['ignoreSSLErrors'] = false;

/**
 * Real Time.
 * @member {Boolean} realTime
 * @default false
 */
OptionTypeListCreate.prototype['realTime'] = false;

/**
 * @member {module:model/OptionTypeListCreateCredential} credential
 */
OptionTypeListCreate.prototype['credential'] = undefined;

/**
 * Username for authenticating with Basic Auth when type is rest or ldap username.
 * @member {String} serviceUsername
 */
OptionTypeListCreate.prototype['serviceUsername'] = undefined;

/**
 * Password for authenticating with Basic Auth when type is rest or ldap password.
 * @member {String} servicePassword
 */
OptionTypeListCreate.prototype['servicePassword'] = undefined;

/**
 * Initial Dataset. Create an initial JSON or CSV dataset to be used as the collection for this option list. It should be a list containing objects with properties 'name', and 'value'. Required when type is manual.
 * @member {String} initialDataset
 */
OptionTypeListCreate.prototype['initialDataset'] = undefined;

/**
 * Translation Script. Create a js script to translate the result data object into an Array containing objects with properties 'name' and 'value'. The input data is provided as data and the result should be put on the global variable results.
 * @member {String} translationScript
 */
OptionTypeListCreate.prototype['translationScript'] = undefined;

/**
 * Request Script. Create a js script to prepare the request. Return a data object as the body for a post, and return an array containing properties 'name' and 'value' for a get. The input data is provided as data and the result should be put on the global variable results.
 * @member {String} requestScript
 */
OptionTypeListCreate.prototype['requestScript'] = undefined;

/**
 * @member {module:model/OptionTypeListCreateConfig} config
 */
OptionTypeListCreate.prototype['config'] = undefined;





/**
 * Allowed values for the <code>type</code> property.
 * @enum {String}
 * @readonly
 */
OptionTypeListCreate['TypeEnum'] = {

    /**
     * value: "rest"
     * @const
     */
    "rest": "rest",

    /**
     * value: "api"
     * @const
     */
    "api": "api",

    /**
     * value: "ldap"
     * @const
     */
    "ldap": "ldap",

    /**
     * value: "manual"
     * @const
     */
    "manual": "manual"
};


/**
 * Allowed values for the <code>visibility</code> property.
 * @enum {String}
 * @readonly
 */
OptionTypeListCreate['VisibilityEnum'] = {

    /**
     * value: "private"
     * @const
     */
    "private": "private",

    /**
     * value: "public"
     * @const
     */
    "public": "public"
};


/**
 * Allowed values for the <code>sourceMethod</code> property.
 * @enum {String}
 * @readonly
 */
OptionTypeListCreate['SourceMethodEnum'] = {

    /**
     * value: "GET"
     * @const
     */
    "GET": "GET",

    /**
     * value: "POST"
     * @const
     */
    "POST": "POST"
};



export default OptionTypeListCreate;
