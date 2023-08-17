/**
 * Morpheus API
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.1.1
 * Contact: dev@morpheusdata.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The AppCreateBlueprintId model module.
 * @module model/AppCreateBlueprintId
 * @version 6.1.1
 */
class AppCreateBlueprintId {
    /**
     * Constructs a new <code>AppCreateBlueprintId</code>.
     * The ID of the Blueprint. Use \&quot;existing\&quot; to create a blank app.
     * @alias module:model/AppCreateBlueprintId
     * @param {(module:model/Number|module:model/String)} instance The actual instance to initialize AppCreateBlueprintId.
     */
    constructor(instance = null) {
        if (instance === null) {
            this.actualInstance = null;
            return;
        }
        var match = 0;
        var errorMessages = [];
        // Blueprint ID
        try {
            this.actualInstance = instance;
            match++;
        } catch(err) {
            // json data failed to deserialize into Number
            errorMessages.push("Failed to construct Number: " + err)
        }

        // String for `existing` selection
        try {
            // validate string
            if (!(typeof instance === 'string')) {
                throw new Error("Invalid value. Must be string. Input: " + JSON.stringify(instance));
            }
            this.actualInstance = instance;
            match++;
        } catch(err) {
            // json data failed to deserialize into String
            errorMessages.push("Failed to construct String: " + err)
        }

        if (match > 1) {
            throw new Error("Multiple matches found constructing `AppCreateBlueprintId` with oneOf schemas Number, String. Input: " + JSON.stringify(instance));
        } else if (match === 0) {
            this.actualInstance = null; // clear the actual instance in case there are multiple matches
            throw new Error("No match found constructing `AppCreateBlueprintId` with oneOf schemas Number, String. Details: " +
                            errorMessages.join(", "));
        } else { // only 1 match
            // the input is valid
        }
    }

    /**
     * Constructs a <code>AppCreateBlueprintId</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/AppCreateBlueprintId} obj Optional instance to populate.
     * @return {module:model/AppCreateBlueprintId} The populated <code>AppCreateBlueprintId</code> instance.
     */
    static constructFromObject(data, obj) {
        return new AppCreateBlueprintId(data);
    }

    /**
     * Gets the actual instance, which can be <code>Number</code>, <code>String</code>.
     * @return {(module:model/Number|module:model/String)} The actual instance.
     */
    getActualInstance() {
        return this.actualInstance;
    }

    /**
     * Sets the actual instance, which can be <code>Number</code>, <code>String</code>.
     * @param {(module:model/Number|module:model/String)} obj The actual instance.
     */
    setActualInstance(obj) {
       this.actualInstance = AppCreateBlueprintId.constructFromObject(obj).getActualInstance();
    }

    /**
     * Returns the JSON representation of the actual instance.
     * @return {string}
     */
    toJSON = function(){
        return this.getActualInstance();
    }

    /**
     * Create an instance of AppCreateBlueprintId from a JSON string.
     * @param {string} json_string JSON string.
     * @return {module:model/AppCreateBlueprintId} An instance of AppCreateBlueprintId.
     */
    static fromJSON = function(json_string){
        return AppCreateBlueprintId.constructFromObject(JSON.parse(json_string));
    }
}


AppCreateBlueprintId.OneOf = ["Number", "String"];

export default AppCreateBlueprintId;

