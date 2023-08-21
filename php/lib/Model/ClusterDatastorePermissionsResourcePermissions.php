<?php
/**
 * ClusterDatastorePermissionsResourcePermissions
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
 * ClusterDatastorePermissionsResourcePermissions Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class ClusterDatastorePermissionsResourcePermissions implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'clusterDatastore_permissions_resourcePermissions';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'all_groups' => 'bool',
        'default_store' => 'bool',
        'all_plans' => 'bool',
        'default_target' => 'bool',
        'morpheus_resource_type' => 'string',
        'morpheus_resource_id' => 'int',
        'can_manage' => 'bool',
        'all' => 'bool',
        'account' => '\OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites',
        'sites' => 'object[]',
        'plans' => 'object[]'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'all_groups' => null,
        'default_store' => null,
        'all_plans' => null,
        'default_target' => null,
        'morpheus_resource_type' => null,
        'morpheus_resource_id' => 'int64',
        'can_manage' => null,
        'all' => null,
        'account' => null,
        'sites' => null,
        'plans' => null
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
        'all_groups' => 'allGroups',
        'default_store' => 'defaultStore',
        'all_plans' => 'allPlans',
        'default_target' => 'defaultTarget',
        'morpheus_resource_type' => 'morpheusResourceType',
        'morpheus_resource_id' => 'morpheusResourceId',
        'can_manage' => 'canManage',
        'all' => 'all',
        'account' => 'account',
        'sites' => 'sites',
        'plans' => 'plans'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'all_groups' => 'setAllGroups',
        'default_store' => 'setDefaultStore',
        'all_plans' => 'setAllPlans',
        'default_target' => 'setDefaultTarget',
        'morpheus_resource_type' => 'setMorpheusResourceType',
        'morpheus_resource_id' => 'setMorpheusResourceId',
        'can_manage' => 'setCanManage',
        'all' => 'setAll',
        'account' => 'setAccount',
        'sites' => 'setSites',
        'plans' => 'setPlans'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'all_groups' => 'getAllGroups',
        'default_store' => 'getDefaultStore',
        'all_plans' => 'getAllPlans',
        'default_target' => 'getDefaultTarget',
        'morpheus_resource_type' => 'getMorpheusResourceType',
        'morpheus_resource_id' => 'getMorpheusResourceId',
        'can_manage' => 'getCanManage',
        'all' => 'getAll',
        'account' => 'getAccount',
        'sites' => 'getSites',
        'plans' => 'getPlans'
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
        $this->container['all_groups'] = $data['all_groups'] ?? null;
        $this->container['default_store'] = $data['default_store'] ?? null;
        $this->container['all_plans'] = $data['all_plans'] ?? null;
        $this->container['default_target'] = $data['default_target'] ?? null;
        $this->container['morpheus_resource_type'] = $data['morpheus_resource_type'] ?? null;
        $this->container['morpheus_resource_id'] = $data['morpheus_resource_id'] ?? null;
        $this->container['can_manage'] = $data['can_manage'] ?? null;
        $this->container['all'] = $data['all'] ?? null;
        $this->container['account'] = $data['account'] ?? null;
        $this->container['sites'] = $data['sites'] ?? null;
        $this->container['plans'] = $data['plans'] ?? null;
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
     * Gets all_groups
     *
     * @return bool|null
     */
    public function getAllGroups()
    {
        return $this->container['all_groups'];
    }

    /**
     * Sets all_groups
     *
     * @param bool|null $all_groups all_groups
     *
     * @return self
     */
    public function setAllGroups($all_groups)
    {
        $this->container['all_groups'] = $all_groups;

        return $this;
    }

    /**
     * Gets default_store
     *
     * @return bool|null
     */
    public function getDefaultStore()
    {
        return $this->container['default_store'];
    }

    /**
     * Sets default_store
     *
     * @param bool|null $default_store default_store
     *
     * @return self
     */
    public function setDefaultStore($default_store)
    {
        $this->container['default_store'] = $default_store;

        return $this;
    }

    /**
     * Gets all_plans
     *
     * @return bool|null
     */
    public function getAllPlans()
    {
        return $this->container['all_plans'];
    }

    /**
     * Sets all_plans
     *
     * @param bool|null $all_plans all_plans
     *
     * @return self
     */
    public function setAllPlans($all_plans)
    {
        $this->container['all_plans'] = $all_plans;

        return $this;
    }

    /**
     * Gets default_target
     *
     * @return bool|null
     */
    public function getDefaultTarget()
    {
        return $this->container['default_target'];
    }

    /**
     * Sets default_target
     *
     * @param bool|null $default_target default_target
     *
     * @return self
     */
    public function setDefaultTarget($default_target)
    {
        $this->container['default_target'] = $default_target;

        return $this;
    }

    /**
     * Gets morpheus_resource_type
     *
     * @return string|null
     */
    public function getMorpheusResourceType()
    {
        return $this->container['morpheus_resource_type'];
    }

    /**
     * Sets morpheus_resource_type
     *
     * @param string|null $morpheus_resource_type morpheus_resource_type
     *
     * @return self
     */
    public function setMorpheusResourceType($morpheus_resource_type)
    {
        $this->container['morpheus_resource_type'] = $morpheus_resource_type;

        return $this;
    }

    /**
     * Gets morpheus_resource_id
     *
     * @return int|null
     */
    public function getMorpheusResourceId()
    {
        return $this->container['morpheus_resource_id'];
    }

    /**
     * Sets morpheus_resource_id
     *
     * @param int|null $morpheus_resource_id morpheus_resource_id
     *
     * @return self
     */
    public function setMorpheusResourceId($morpheus_resource_id)
    {
        $this->container['morpheus_resource_id'] = $morpheus_resource_id;

        return $this;
    }

    /**
     * Gets can_manage
     *
     * @return bool|null
     */
    public function getCanManage()
    {
        return $this->container['can_manage'];
    }

    /**
     * Sets can_manage
     *
     * @param bool|null $can_manage can_manage
     *
     * @return self
     */
    public function setCanManage($can_manage)
    {
        $this->container['can_manage'] = $can_manage;

        return $this;
    }

    /**
     * Gets all
     *
     * @return bool|null
     */
    public function getAll()
    {
        return $this->container['all'];
    }

    /**
     * Sets all
     *
     * @param bool|null $all all
     *
     * @return self
     */
    public function setAll($all)
    {
        $this->container['all'] = $all;

        return $this;
    }

    /**
     * Gets account
     *
     * @return \OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites|null
     */
    public function getAccount()
    {
        return $this->container['account'];
    }

    /**
     * Sets account
     *
     * @param \OpenAPI\Client\Model\ApiBlueprintsIdUpdatePermissionsResourcePermissionSites|null $account account
     *
     * @return self
     */
    public function setAccount($account)
    {
        $this->container['account'] = $account;

        return $this;
    }

    /**
     * Gets sites
     *
     * @return object[]|null
     */
    public function getSites()
    {
        return $this->container['sites'];
    }

    /**
     * Sets sites
     *
     * @param object[]|null $sites sites
     *
     * @return self
     */
    public function setSites($sites)
    {
        $this->container['sites'] = $sites;

        return $this;
    }

    /**
     * Gets plans
     *
     * @return object[]|null
     */
    public function getPlans()
    {
        return $this->container['plans'];
    }

    /**
     * Sets plans
     *
     * @param object[]|null $plans plans
     *
     * @return self
     */
    public function setPlans($plans)
    {
        $this->container['plans'] = $plans;

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


