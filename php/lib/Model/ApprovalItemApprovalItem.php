<?php
/**
 * ApprovalItemApprovalItem
 *
 * PHP version 7.2
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Morpheus API
 *
 * Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.
 *
 * The version of the OpenAPI document: 6.2.1
 * Contact: dev@morpheusdata.com
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.0.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace OpenAPI\Client\Model;

use \ArrayAccess;
use \OpenAPI\Client\ObjectSerializer;

/**
 * ApprovalItemApprovalItem Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ApprovalItemApprovalItem implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'approvalItem_approvalItem';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'name' => 'string',
        'external_id' => 'string',
        'external_name' => 'string',
        'internal_id' => 'string',
        'approved_by' => 'string',
        'denied_by' => 'string',
        'status' => 'string',
        'error_message' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'date_approved' => '\DateTime',
        'date_denied' => '\DateTime',
        'approval' => '\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites',
        'reference' => '\OpenAPI\Client\Model\ApprovalItemApprovalItemReference'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'id' => 'int64',
        'name' => null,
        'external_id' => null,
        'external_name' => null,
        'internal_id' => null,
        'approved_by' => null,
        'denied_by' => null,
        'status' => null,
        'error_message' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'date_approved' => 'date-time',
        'date_denied' => 'date-time',
        'approval' => null,
        'reference' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'id' => 'id',
        'name' => 'name',
        'external_id' => 'externalId',
        'external_name' => 'externalName',
        'internal_id' => 'internalId',
        'approved_by' => 'approvedBy',
        'denied_by' => 'deniedBy',
        'status' => 'status',
        'error_message' => 'errorMessage',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'date_approved' => 'dateApproved',
        'date_denied' => 'dateDenied',
        'approval' => 'approval',
        'reference' => 'reference'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'name' => 'setName',
        'external_id' => 'setExternalId',
        'external_name' => 'setExternalName',
        'internal_id' => 'setInternalId',
        'approved_by' => 'setApprovedBy',
        'denied_by' => 'setDeniedBy',
        'status' => 'setStatus',
        'error_message' => 'setErrorMessage',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'date_approved' => 'setDateApproved',
        'date_denied' => 'setDateDenied',
        'approval' => 'setApproval',
        'reference' => 'setReference'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'name' => 'getName',
        'external_id' => 'getExternalId',
        'external_name' => 'getExternalName',
        'internal_id' => 'getInternalId',
        'approved_by' => 'getApprovedBy',
        'denied_by' => 'getDeniedBy',
        'status' => 'getStatus',
        'error_message' => 'getErrorMessage',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'date_approved' => 'getDateApproved',
        'date_denied' => 'getDateDenied',
        'approval' => 'getApproval',
        'reference' => 'getReference'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }

    

    

    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['id'] = $data['id'] ?? null;
        $this->container['name'] = $data['name'] ?? null;
        $this->container['external_id'] = $data['external_id'] ?? null;
        $this->container['external_name'] = $data['external_name'] ?? null;
        $this->container['internal_id'] = $data['internal_id'] ?? null;
        $this->container['approved_by'] = $data['approved_by'] ?? null;
        $this->container['denied_by'] = $data['denied_by'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['error_message'] = $data['error_message'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['date_approved'] = $data['date_approved'] ?? null;
        $this->container['date_denied'] = $data['date_denied'] ?? null;
        $this->container['approval'] = $data['approval'] ?? null;
        $this->container['reference'] = $data['reference'] ?? null;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets id
     *
     * @return int|null
     */
    public function getId()
    {
        return $this->container['id'];
    }

    /**
     * Sets id
     *
     * @param int|null $id id
     *
     * @return self
     */
    public function setId($id)
    {
        $this->container['id'] = $id;

        return $this;
    }

    /**
     * Gets name
     *
     * @return string|null
     */
    public function getName()
    {
        return $this->container['name'];
    }

    /**
     * Sets name
     *
     * @param string|null $name name
     *
     * @return self
     */
    public function setName($name)
    {
        $this->container['name'] = $name;

        return $this;
    }

    /**
     * Gets external_id
     *
     * @return string|null
     */
    public function getExternalId()
    {
        return $this->container['external_id'];
    }

    /**
     * Sets external_id
     *
     * @param string|null $external_id external_id
     *
     * @return self
     */
    public function setExternalId($external_id)
    {
        $this->container['external_id'] = $external_id;

        return $this;
    }

    /**
     * Gets external_name
     *
     * @return string|null
     */
    public function getExternalName()
    {
        return $this->container['external_name'];
    }

    /**
     * Sets external_name
     *
     * @param string|null $external_name external_name
     *
     * @return self
     */
    public function setExternalName($external_name)
    {
        $this->container['external_name'] = $external_name;

        return $this;
    }

    /**
     * Gets internal_id
     *
     * @return string|null
     */
    public function getInternalId()
    {
        return $this->container['internal_id'];
    }

    /**
     * Sets internal_id
     *
     * @param string|null $internal_id internal_id
     *
     * @return self
     */
    public function setInternalId($internal_id)
    {
        $this->container['internal_id'] = $internal_id;

        return $this;
    }

    /**
     * Gets approved_by
     *
     * @return string|null
     */
    public function getApprovedBy()
    {
        return $this->container['approved_by'];
    }

    /**
     * Sets approved_by
     *
     * @param string|null $approved_by approved_by
     *
     * @return self
     */
    public function setApprovedBy($approved_by)
    {
        $this->container['approved_by'] = $approved_by;

        return $this;
    }

    /**
     * Gets denied_by
     *
     * @return string|null
     */
    public function getDeniedBy()
    {
        return $this->container['denied_by'];
    }

    /**
     * Sets denied_by
     *
     * @param string|null $denied_by denied_by
     *
     * @return self
     */
    public function setDeniedBy($denied_by)
    {
        $this->container['denied_by'] = $denied_by;

        return $this;
    }

    /**
     * Gets status
     *
     * @return string|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param string|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {
        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets error_message
     *
     * @return string|null
     */
    public function getErrorMessage()
    {
        return $this->container['error_message'];
    }

    /**
     * Sets error_message
     *
     * @param string|null $error_message error_message
     *
     * @return self
     */
    public function setErrorMessage($error_message)
    {
        $this->container['error_message'] = $error_message;

        return $this;
    }

    /**
     * Gets date_created
     *
     * @return \DateTime|null
     */
    public function getDateCreated()
    {
        return $this->container['date_created'];
    }

    /**
     * Sets date_created
     *
     * @param \DateTime|null $date_created date_created
     *
     * @return self
     */
    public function setDateCreated($date_created)
    {
        $this->container['date_created'] = $date_created;

        return $this;
    }

    /**
     * Gets last_updated
     *
     * @return \DateTime|null
     */
    public function getLastUpdated()
    {
        return $this->container['last_updated'];
    }

    /**
     * Sets last_updated
     *
     * @param \DateTime|null $last_updated last_updated
     *
     * @return self
     */
    public function setLastUpdated($last_updated)
    {
        $this->container['last_updated'] = $last_updated;

        return $this;
    }

    /**
     * Gets date_approved
     *
     * @return \DateTime|null
     */
    public function getDateApproved()
    {
        return $this->container['date_approved'];
    }

    /**
     * Sets date_approved
     *
     * @param \DateTime|null $date_approved date_approved
     *
     * @return self
     */
    public function setDateApproved($date_approved)
    {
        $this->container['date_approved'] = $date_approved;

        return $this;
    }

    /**
     * Gets date_denied
     *
     * @return \DateTime|null
     */
    public function getDateDenied()
    {
        return $this->container['date_denied'];
    }

    /**
     * Sets date_denied
     *
     * @param \DateTime|null $date_denied date_denied
     *
     * @return self
     */
    public function setDateDenied($date_denied)
    {
        $this->container['date_denied'] = $date_denied;

        return $this;
    }

    /**
     * Gets approval
     *
     * @return \OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites|null
     */
    public function getApproval()
    {
        return $this->container['approval'];
    }

    /**
     * Sets approval
     *
     * @param \OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites|null $approval approval
     *
     * @return self
     */
    public function setApproval($approval)
    {
        $this->container['approval'] = $approval;

        return $this;
    }

    /**
     * Gets reference
     *
     * @return \OpenAPI\Client\Model\ApprovalItemApprovalItemReference|null
     */
    public function getReference()
    {
        return $this->container['reference'];
    }

    /**
     * Sets reference
     *
     * @param \OpenAPI\Client\Model\ApprovalItemApprovalItemReference|null $reference reference
     *
     * @return self
     */
    public function setReference($reference)
    {
        $this->container['reference'] = $reference;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


