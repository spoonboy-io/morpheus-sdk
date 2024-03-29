<?php
/**
 * InlineResponse20040AppDeploy
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
 * InlineResponse20040AppDeploy Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class InlineResponse20040AppDeploy implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'inline_response_200_40_appDeploy';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'instance_id' => 'int',
        'instance' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployInstance',
        'deployment' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployDeployment',
        'deployment_version_id' => 'int',
        'deployment_version' => '\OpenAPI\Client\Model\InlineResponse20040AppDeployDeploymentVersion',
        'config' => 'object',
        'status' => 'string',
        'deploy_date' => '\DateTime',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime'
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
        'instance_id' => 'int64',
        'instance' => null,
        'deployment' => null,
        'deployment_version_id' => 'int64',
        'deployment_version' => null,
        'config' => null,
        'status' => null,
        'deploy_date' => 'date-time',
        'date_created' => 'date-time',
        'last_updated' => 'date-time'
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
        'instance_id' => 'instanceId',
        'instance' => 'instance',
        'deployment' => 'deployment',
        'deployment_version_id' => 'deploymentVersionId',
        'deployment_version' => 'deploymentVersion',
        'config' => 'config',
        'status' => 'status',
        'deploy_date' => 'deployDate',
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'instance_id' => 'setInstanceId',
        'instance' => 'setInstance',
        'deployment' => 'setDeployment',
        'deployment_version_id' => 'setDeploymentVersionId',
        'deployment_version' => 'setDeploymentVersion',
        'config' => 'setConfig',
        'status' => 'setStatus',
        'deploy_date' => 'setDeployDate',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'instance_id' => 'getInstanceId',
        'instance' => 'getInstance',
        'deployment' => 'getDeployment',
        'deployment_version_id' => 'getDeploymentVersionId',
        'deployment_version' => 'getDeploymentVersion',
        'config' => 'getConfig',
        'status' => 'getStatus',
        'deploy_date' => 'getDeployDate',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated'
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
        $this->container['instance_id'] = $data['instance_id'] ?? null;
        $this->container['instance'] = $data['instance'] ?? null;
        $this->container['deployment'] = $data['deployment'] ?? null;
        $this->container['deployment_version_id'] = $data['deployment_version_id'] ?? null;
        $this->container['deployment_version'] = $data['deployment_version'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
        $this->container['status'] = $data['status'] ?? null;
        $this->container['deploy_date'] = $data['deploy_date'] ?? null;
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
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
     * Gets instance_id
     *
     * @return int|null
     */
    public function getInstanceId()
    {
        return $this->container['instance_id'];
    }

    /**
     * Sets instance_id
     *
     * @param int|null $instance_id instance_id
     *
     * @return self
     */
    public function setInstanceId($instance_id)
    {
        $this->container['instance_id'] = $instance_id;

        return $this;
    }

    /**
     * Gets instance
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null
     */
    public function getInstance()
    {
        return $this->container['instance'];
    }

    /**
     * Sets instance
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployInstance|null $instance instance
     *
     * @return self
     */
    public function setInstance($instance)
    {
        $this->container['instance'] = $instance;

        return $this;
    }

    /**
     * Gets deployment
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployDeployment|null
     */
    public function getDeployment()
    {
        return $this->container['deployment'];
    }

    /**
     * Sets deployment
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployDeployment|null $deployment deployment
     *
     * @return self
     */
    public function setDeployment($deployment)
    {
        $this->container['deployment'] = $deployment;

        return $this;
    }

    /**
     * Gets deployment_version_id
     *
     * @return int|null
     */
    public function getDeploymentVersionId()
    {
        return $this->container['deployment_version_id'];
    }

    /**
     * Sets deployment_version_id
     *
     * @param int|null $deployment_version_id deployment_version_id
     *
     * @return self
     */
    public function setDeploymentVersionId($deployment_version_id)
    {
        $this->container['deployment_version_id'] = $deployment_version_id;

        return $this;
    }

    /**
     * Gets deployment_version
     *
     * @return \OpenAPI\Client\Model\InlineResponse20040AppDeployDeploymentVersion|null
     */
    public function getDeploymentVersion()
    {
        return $this->container['deployment_version'];
    }

    /**
     * Sets deployment_version
     *
     * @param \OpenAPI\Client\Model\InlineResponse20040AppDeployDeploymentVersion|null $deployment_version deployment_version
     *
     * @return self
     */
    public function setDeploymentVersion($deployment_version)
    {
        $this->container['deployment_version'] = $deployment_version;

        return $this;
    }

    /**
     * Gets config
     *
     * @return object|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param object|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

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
     * Gets deploy_date
     *
     * @return \DateTime|null
     */
    public function getDeployDate()
    {
        return $this->container['deploy_date'];
    }

    /**
     * Sets deploy_date
     *
     * @param \DateTime|null $deploy_date deploy_date
     *
     * @return self
     */
    public function setDeployDate($deploy_date)
    {
        $this->container['deploy_date'] = $deploy_date;

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


