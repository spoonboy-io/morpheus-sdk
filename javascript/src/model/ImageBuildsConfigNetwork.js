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
import AppPrepareApplyDataGroup from './AppPrepareApplyDataGroup';

/**
 * The ImageBuildsConfigNetwork model module.
 * @module model/ImageBuildsConfigNetwork
 * @version 6.2.1
 */
class ImageBuildsConfigNetwork {
    /**
     * Constructs a new <code>ImageBuildsConfigNetwork</code>.
     * @alias module:model/ImageBuildsConfigNetwork
     */
    constructor() { 
        
        ImageBuildsConfigNetwork.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>ImageBuildsConfigNetwork</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/ImageBuildsConfigNetwork} obj Optional instance to populate.
     * @return {module:model/ImageBuildsConfigNetwork} The populated <code>ImageBuildsConfigNetwork</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new ImageBuildsConfigNetwork();

            if (data.hasOwnProperty('idName')) {
                obj['idName'] = ApiClient.convertToType(data['idName'], 'String');
            }
            if (data.hasOwnProperty('pool')) {
                obj['pool'] = AppPrepareApplyDataGroup.constructFromObject(data['pool']);
            }
            if (data.hasOwnProperty('id')) {
                obj['id'] = ApiClient.convertToType(data['id'], 'String');
            }
            if (data.hasOwnProperty('hasPool')) {
                obj['hasPool'] = ApiClient.convertToType(data['hasPool'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {String} idName
 */
ImageBuildsConfigNetwork.prototype['idName'] = undefined;

/**
 * @member {module:model/AppPrepareApplyDataGroup} pool
 */
ImageBuildsConfigNetwork.prototype['pool'] = undefined;

/**
 * @member {String} id
 */
ImageBuildsConfigNetwork.prototype['id'] = undefined;

/**
 * @member {Boolean} hasPool
 */
ImageBuildsConfigNetwork.prototype['hasPool'] = undefined;






export default ImageBuildsConfigNetwork;

