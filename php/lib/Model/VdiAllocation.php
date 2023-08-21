<?php
/**
 * VdiAllocation
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
 * VdiAllocation Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class VdiAllocation implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'vdi_allocation';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'pool_id' => 'int',
        'pool' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'instance' => '\OpenAPI\Client\Model\VdiAllocationInstance',
        'user' => '\OpenAPI\Client\Model\VdiPoolOwner',
        'local_user_created' => 'bool',
        'persistent' => 'bool',
        'recyclable' => 'bool',
        'status' => 'string',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'last_reserved' => '\DateTime',
        'release_date' => '\DateTime'
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
        'pool_id' => 'int64',
        'pool' => null,
        'instance' => null,
        'user' => null,
        'local_user_created' => null,
        'persistent' => null,
        'recyclable' => null,
        'status' => null,
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'last_reserved' => 'date-time',
        'release_date' => 'date-time'
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
        'pool_id' => 'poolId',
        'pool' => 'pool',
        'instance' => 'instance',
        'user' => 'user',
        'local_user_created' => 'localUserCreated',
        'persistent' => 'persistent',
        'recyclable' => 'recyclable',
        'status' => 'status',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'last_reserved' => 'lastReserved',
        'release_date' => 'releaseDate'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'pool_id' => 'setPoolId',
        'pool' => 'setPool',
        'instance' => 'setInstance',
        'user' => 'setUser',
        'local_user_created' => 'setLocalUserCreated',
        'persistent' => 'setPersistent',
        'recyclable' => 'setRecyclable',
        'status' => 'setStatus',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'last_reserved' => 'setLastReserved',
        'release_date' => 'setReleaseDate'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'pool_id' => 'getPoolId',
        'pool' => 'getPool',
        'instance' => 'getInstance',
        'user' => 'getUser',
        'local_user_created' => 'getLocalUserCreated',
        'persistent' => 'getPersistent',
        'recyclable' => 'getRecyclable',
        'status' => 'getStatus',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'last_reserved' => 'getLastReserved',
        'release_date' => 'getReleaseDate'
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
        $this->container['pool_id'] = $data['pool_id'] ?? null;
        $this->container['pool'] = $data['pool'] ?? null;
        $this->container['instance'] = $data['instance'] ?? null;
        $this->container['user'] = $data['user'] ?? null;
        $this->container['local_user_created'] = $data['local_user_created'] ?? null;
        $this->container['persistent'] = $data['persistent'] ?? null;
        $this->container['recyclable'] = $data['recyclable'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['last_reserved'] = $data['last_reserved'] ?? null;
        $this->container['release_date'] = $data['release_date'] ?? null;
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
     * Gets pool_id
     *
     * @return int|null
     */
    public function getPoolId()
    {
        return $this->container['pool_id'];
    }

    /**
     * Sets pool_id
     *
     * @param int|null $pool_id pool_id
     *
     * @return self
     */
    public function setPoolId($pool_id)
    {
        $this->container['pool_id'] = $pool_id;

        return $this;
    }

    /**
     * Gets pool
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getPool()
    {
        return $this->container['pool'];
    }

    /**
     * Sets pool
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $pool pool
     *
     * @return self
     */
    public function setPool($pool)
    {
        $this->container['pool'] = $pool;

        return $this;
    }

    /**
     * Gets instance
     *
     * @return \OpenAPI\Client\Model\VdiAllocationInstance|null
     */
    public function getInstance()
    {
        return $this->container['instance'];
    }

    /**
     * Sets instance
     *
     * @param \OpenAPI\Client\Model\VdiAllocationInstance|null $instance instance
     *
     * @return self
     */
    public function setInstance($instance)
    {
        $this->container['instance'] = $instance;

        return $this;
    }

    /**
     * Gets user
     *
     * @return \OpenAPI\Client\Model\VdiPoolOwner|null
     */
    public function getUser()
    {
        return $this->container['user'];
    }

    /**
     * Sets user
     *
     * @param \OpenAPI\Client\Model\VdiPoolOwner|null $user user
     *
     * @return self
     */
    public function setUser($user)
    {
        $this->container['user'] = $user;

        return $this;
    }

    /**
     * Gets local_user_created
     *
     * @return bool|null
     */
    public function getLocalUserCreated()
    {
        return $this->container['local_user_created'];
    }

    /**
     * Sets local_user_created
     *
     * @param bool|null $local_user_created local_user_created
     *
     * @return self
     */
    public function setLocalUserCreated($local_user_created)
    {
        $this->container['local_user_created'] = $local_user_created;

        return $this;
    }

    /**
     * Gets persistent
     *
     * @return bool|null
     */
    public function getPersistent()
    {
        return $this->container['persistent'];
    }

    /**
     * Sets persistent
     *
     * @param bool|null $persistent persistent
     *
     * @return self
     */
    public function setPersistent($persistent)
    {
        $this->container['persistent'] = $persistent;

        return $this;
    }

    /**
     * Gets recyclable
     *
     * @return bool|null
     */
    public function getRecyclable()
    {
        return $this->container['recyclable'];
    }

    /**
     * Sets recyclable
     *
     * @param bool|null $recyclable recyclable
     *
     * @return self
     */
    public function setRecyclable($recyclable)
    {
        $this->container['recyclable'] = $recyclable;

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
     * Gets last_reserved
     *
     * @return \DateTime|null
     */
    public function getLastReserved()
    {
        return $this->container['last_reserved'];
    }

    /**
     * Sets last_reserved
     *
     * @param \DateTime|null $last_reserved last_reserved
     *
     * @return self
     */
    public function setLastReserved($last_reserved)
    {
        $this->container['last_reserved'] = $last_reserved;

        return $this;
    }

    /**
     * Gets release_date
     *
     * @return \DateTime|null
     */
    public function getReleaseDate()
    {
        return $this->container['release_date'];
    }

    /**
     * Sets release_date
     *
     * @param \DateTime|null $release_date release_date
     *
     * @return self
     */
    public function setReleaseDate($release_date)
    {
        $this->container['release_date'] = $release_date;

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

