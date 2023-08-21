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
import MetaObject from './MetaObject';
import SearchHits from './SearchHits';

/**
 * The Search model module.
 * @module model/Search
 * @version 6.2.1
 */
class Search {
    /**
     * Constructs a new <code>Search</code>.
     * @alias module:model/Search
     */
    constructor() { 
        
        Search.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>Search</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/Search} obj Optional instance to populate.
     * @return {module:model/Search} The populated <code>Search</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new Search();

            if (data.hasOwnProperty('hits')) {
                obj['hits'] = ApiClient.convertToType(data['hits'], [SearchHits]);
            }
            if (data.hasOwnProperty('query')) {
                obj['query'] = ApiClient.convertToType(data['query'], 'String');
            }
            if (data.hasOwnProperty('took')) {
                obj['took'] = ApiClient.convertToType(data['took'], 'Number');
            }
            if (data.hasOwnProperty('meta')) {
                obj['meta'] = MetaObject.constructFromObject(data['meta']);
            }
        }
        return obj;
    }


}

/**
 * @member {Array.<module:model/SearchHits>} hits
 */
Search.prototype['hits'] = undefined;

/**
 * @member {String} query
 */
Search.prototype['query'] = undefined;

/**
 * @member {Number} took
 */
Search.prototype['took'] = undefined;

/**
 * @member {module:model/MetaObject} meta
 */
Search.prototype['meta'] = undefined;






export default Search;

