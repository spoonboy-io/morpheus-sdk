<?php
/**
 * GuidanceVmwareSizing
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
 * GuidanceVmwareSizing Class Doc Comment
 *
 * @category Class
 * @package  OpenAPI\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null  
 */
class GuidanceVmwareSizing implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'guidanceVmwareSizing';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'id' => 'int',
        'date_created' => '\DateTime',
        'last_updated' => '\DateTime',
        'action_category' => 'string',
        'action_message' => 'string',
        'action_title' => 'string',
        'action_type' => 'string',
        'action_value' => 'string',
        'action_value_type' => 'string',
        'action_plan_id' => 'int',
        'status_message' => 'string',
        'account_id' => 'int',
        'user_id' => 'string',
        'site_id' => 'int',
        'zone' => '\OpenAPI\Client\Model\GuidanceVmwareSizingZone',
        'state' => 'string',
        'state_message' => 'string',
        'severity' => 'string',
        'resolved' => 'bool',
        'resolved_message' => 'string',
        'ref_type' => 'string',
        'ref_id' => 'int',
        'ref_name' => 'string',
        'type' => '\OpenAPI\Client\Model\GuidanceVmwareSizingType',
        'savings' => '\OpenAPI\Client\Model\GuidanceVmwareSizingSavings',
        'resource' => '\OpenAPI\Client\Model\GuidanceVmwareSizingResource',
        'plan_before_action' => '\OpenAPI\Client\Model\GuidanceVmwareSizingPlanBeforeAction',
        'plan_after_action' => '\OpenAPI\Client\Model\GuidanceVmwareSizingPlanAfterAction',
        'config' => '\OpenAPI\Client\Model\GuidanceVmwareSizingConfig'
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
        'date_created' => 'date-time',
        'last_updated' => 'date-time',
        'action_category' => null,
        'action_message' => null,
        'action_title' => null,
        'action_type' => null,
        'action_value' => null,
        'action_value_type' => null,
        'action_plan_id' => 'int64',
        'status_message' => null,
        'account_id' => 'int64',
        'user_id' => null,
        'site_id' => 'int64',
        'zone' => null,
        'state' => null,
        'state_message' => null,
        'severity' => null,
        'resolved' => null,
        'resolved_message' => null,
        'ref_type' => null,
        'ref_id' => 'int64',
        'ref_name' => null,
        'type' => null,
        'savings' => null,
        'resource' => null,
        'plan_before_action' => null,
        'plan_after_action' => null,
        'config' => null
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
        'date_created' => 'dateCreated',
        'last_updated' => 'lastUpdated',
        'action_category' => 'actionCategory',
        'action_message' => 'actionMessage',
        'action_title' => 'actionTitle',
        'action_type' => 'actionType',
        'action_value' => 'actionValue',
        'action_value_type' => 'actionValueType',
        'action_plan_id' => 'actionPlanId',
        'status_message' => 'statusMessage',
        'account_id' => 'accountId',
        'user_id' => 'userId',
        'site_id' => 'siteId',
        'zone' => 'zone',
        'state' => 'state',
        'state_message' => 'stateMessage',
        'severity' => 'severity',
        'resolved' => 'resolved',
        'resolved_message' => 'resolvedMessage',
        'ref_type' => 'refType',
        'ref_id' => 'refId',
        'ref_name' => 'refName',
        'type' => 'type',
        'savings' => 'savings',
        'resource' => 'resource',
        'plan_before_action' => 'planBeforeAction',
        'plan_after_action' => 'planAfterAction',
        'config' => 'config'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'id' => 'setId',
        'date_created' => 'setDateCreated',
        'last_updated' => 'setLastUpdated',
        'action_category' => 'setActionCategory',
        'action_message' => 'setActionMessage',
        'action_title' => 'setActionTitle',
        'action_type' => 'setActionType',
        'action_value' => 'setActionValue',
        'action_value_type' => 'setActionValueType',
        'action_plan_id' => 'setActionPlanId',
        'status_message' => 'setStatusMessage',
        'account_id' => 'setAccountId',
        'user_id' => 'setUserId',
        'site_id' => 'setSiteId',
        'zone' => 'setZone',
        'state' => 'setState',
        'state_message' => 'setStateMessage',
        'severity' => 'setSeverity',
        'resolved' => 'setResolved',
        'resolved_message' => 'setResolvedMessage',
        'ref_type' => 'setRefType',
        'ref_id' => 'setRefId',
        'ref_name' => 'setRefName',
        'type' => 'setType',
        'savings' => 'setSavings',
        'resource' => 'setResource',
        'plan_before_action' => 'setPlanBeforeAction',
        'plan_after_action' => 'setPlanAfterAction',
        'config' => 'setConfig'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'id' => 'getId',
        'date_created' => 'getDateCreated',
        'last_updated' => 'getLastUpdated',
        'action_category' => 'getActionCategory',
        'action_message' => 'getActionMessage',
        'action_title' => 'getActionTitle',
        'action_type' => 'getActionType',
        'action_value' => 'getActionValue',
        'action_value_type' => 'getActionValueType',
        'action_plan_id' => 'getActionPlanId',
        'status_message' => 'getStatusMessage',
        'account_id' => 'getAccountId',
        'user_id' => 'getUserId',
        'site_id' => 'getSiteId',
        'zone' => 'getZone',
        'state' => 'getState',
        'state_message' => 'getStateMessage',
        'severity' => 'getSeverity',
        'resolved' => 'getResolved',
        'resolved_message' => 'getResolvedMessage',
        'ref_type' => 'getRefType',
        'ref_id' => 'getRefId',
        'ref_name' => 'getRefName',
        'type' => 'getType',
        'savings' => 'getSavings',
        'resource' => 'getResource',
        'plan_before_action' => 'getPlanBeforeAction',
        'plan_after_action' => 'getPlanAfterAction',
        'config' => 'getConfig'
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
        $this->container['date_created'] = $data['date_created'] ?? null;
        $this->container['last_updated'] = $data['last_updated'] ?? null;
        $this->container['action_category'] = $data['action_category'] ?? null;
        $this->container['action_message'] = $data['action_message'] ?? null;
        $this->container['action_title'] = $data['action_title'] ?? null;
        $this->container['action_type'] = $data['action_type'] ?? null;
        $this->container['action_value'] = $data['action_value'] ?? null;
        $this->container['action_value_type'] = $data['action_value_type'] ?? null;
        $this->container['action_plan_id'] = $data['action_plan_id'] ?? null;
        $this->container['status_message'] = $data['status_message'] ?? null;
        $this->container['account_id'] = $data['account_id'] ?? null;
        $this->container['user_id'] = $data['user_id'] ?? null;
        $this->container['site_id'] = $data['site_id'] ?? null;
        $this->container['zone'] = $data['zone'] ?? null;
        $this->container['state'] = $data['state'] ?? null;
        $this->container['state_message'] = $data['state_message'] ?? null;
        $this->container['severity'] = $data['severity'] ?? null;
        $this->container['resolved'] = $data['resolved'] ?? null;
        $this->container['resolved_message'] = $data['resolved_message'] ?? null;
        $this->container['ref_type'] = $data['ref_type'] ?? null;
        $this->container['ref_id'] = $data['ref_id'] ?? null;
        $this->container['ref_name'] = $data['ref_name'] ?? null;
        $this->container['type'] = $data['type'] ?? null;
        $this->container['savings'] = $data['savings'] ?? null;
        $this->container['resource'] = $data['resource'] ?? null;
        $this->container['plan_before_action'] = $data['plan_before_action'] ?? null;
        $this->container['plan_after_action'] = $data['plan_after_action'] ?? null;
        $this->container['config'] = $data['config'] ?? null;
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
     * Gets action_category
     *
     * @return string|null
     */
    public function getActionCategory()
    {
        return $this->container['action_category'];
    }

    /**
     * Sets action_category
     *
     * @param string|null $action_category action_category
     *
     * @return self
     */
    public function setActionCategory($action_category)
    {
        $this->container['action_category'] = $action_category;

        return $this;
    }

    /**
     * Gets action_message
     *
     * @return string|null
     */
    public function getActionMessage()
    {
        return $this->container['action_message'];
    }

    /**
     * Sets action_message
     *
     * @param string|null $action_message action_message
     *
     * @return self
     */
    public function setActionMessage($action_message)
    {
        $this->container['action_message'] = $action_message;

        return $this;
    }

    /**
     * Gets action_title
     *
     * @return string|null
     */
    public function getActionTitle()
    {
        return $this->container['action_title'];
    }

    /**
     * Sets action_title
     *
     * @param string|null $action_title action_title
     *
     * @return self
     */
    public function setActionTitle($action_title)
    {
        $this->container['action_title'] = $action_title;

        return $this;
    }

    /**
     * Gets action_type
     *
     * @return string|null
     */
    public function getActionType()
    {
        return $this->container['action_type'];
    }

    /**
     * Sets action_type
     *
     * @param string|null $action_type action_type
     *
     * @return self
     */
    public function setActionType($action_type)
    {
        $this->container['action_type'] = $action_type;

        return $this;
    }

    /**
     * Gets action_value
     *
     * @return string|null
     */
    public function getActionValue()
    {
        return $this->container['action_value'];
    }

    /**
     * Sets action_value
     *
     * @param string|null $action_value action_value
     *
     * @return self
     */
    public function setActionValue($action_value)
    {
        $this->container['action_value'] = $action_value;

        return $this;
    }

    /**
     * Gets action_value_type
     *
     * @return string|null
     */
    public function getActionValueType()
    {
        return $this->container['action_value_type'];
    }

    /**
     * Sets action_value_type
     *
     * @param string|null $action_value_type action_value_type
     *
     * @return self
     */
    public function setActionValueType($action_value_type)
    {
        $this->container['action_value_type'] = $action_value_type;

        return $this;
    }

    /**
     * Gets action_plan_id
     *
     * @return int|null
     */
    public function getActionPlanId()
    {
        return $this->container['action_plan_id'];
    }

    /**
     * Sets action_plan_id
     *
     * @param int|null $action_plan_id action_plan_id
     *
     * @return self
     */
    public function setActionPlanId($action_plan_id)
    {
        $this->container['action_plan_id'] = $action_plan_id;

        return $this;
    }

    /**
     * Gets status_message
     *
     * @return string|null
     */
    public function getStatusMessage()
    {
        return $this->container['status_message'];
    }

    /**
     * Sets status_message
     *
     * @param string|null $status_message status_message
     *
     * @return self
     */
    public function setStatusMessage($status_message)
    {
        $this->container['status_message'] = $status_message;

        return $this;
    }

    /**
     * Gets account_id
     *
     * @return int|null
     */
    public function getAccountId()
    {
        return $this->container['account_id'];
    }

    /**
     * Sets account_id
     *
     * @param int|null $account_id account_id
     *
     * @return self
     */
    public function setAccountId($account_id)
    {
        $this->container['account_id'] = $account_id;

        return $this;
    }

    /**
     * Gets user_id
     *
     * @return string|null
     */
    public function getUserId()
    {
        return $this->container['user_id'];
    }

    /**
     * Sets user_id
     *
     * @param string|null $user_id user_id
     *
     * @return self
     */
    public function setUserId($user_id)
    {
        $this->container['user_id'] = $user_id;

        return $this;
    }

    /**
     * Gets site_id
     *
     * @return int|null
     */
    public function getSiteId()
    {
        return $this->container['site_id'];
    }

    /**
     * Sets site_id
     *
     * @param int|null $site_id site_id
     *
     * @return self
     */
    public function setSiteId($site_id)
    {
        $this->container['site_id'] = $site_id;

        return $this;
    }

    /**
     * Gets zone
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingZone|null
     */
    public function getZone()
    {
        return $this->container['zone'];
    }

    /**
     * Sets zone
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingZone|null $zone zone
     *
     * @return self
     */
    public function setZone($zone)
    {
        $this->container['zone'] = $zone;

        return $this;
    }

    /**
     * Gets state
     *
     * @return string|null
     */
    public function getState()
    {
        return $this->container['state'];
    }

    /**
     * Sets state
     *
     * @param string|null $state state
     *
     * @return self
     */
    public function setState($state)
    {
        $this->container['state'] = $state;

        return $this;
    }

    /**
     * Gets state_message
     *
     * @return string|null
     */
    public function getStateMessage()
    {
        return $this->container['state_message'];
    }

    /**
     * Sets state_message
     *
     * @param string|null $state_message state_message
     *
     * @return self
     */
    public function setStateMessage($state_message)
    {
        $this->container['state_message'] = $state_message;

        return $this;
    }

    /**
     * Gets severity
     *
     * @return string|null
     */
    public function getSeverity()
    {
        return $this->container['severity'];
    }

    /**
     * Sets severity
     *
     * @param string|null $severity severity
     *
     * @return self
     */
    public function setSeverity($severity)
    {
        $this->container['severity'] = $severity;

        return $this;
    }

    /**
     * Gets resolved
     *
     * @return bool|null
     */
    public function getResolved()
    {
        return $this->container['resolved'];
    }

    /**
     * Sets resolved
     *
     * @param bool|null $resolved resolved
     *
     * @return self
     */
    public function setResolved($resolved)
    {
        $this->container['resolved'] = $resolved;

        return $this;
    }

    /**
     * Gets resolved_message
     *
     * @return string|null
     */
    public function getResolvedMessage()
    {
        return $this->container['resolved_message'];
    }

    /**
     * Sets resolved_message
     *
     * @param string|null $resolved_message resolved_message
     *
     * @return self
     */
    public function setResolvedMessage($resolved_message)
    {
        $this->container['resolved_message'] = $resolved_message;

        return $this;
    }

    /**
     * Gets ref_type
     *
     * @return string|null
     */
    public function getRefType()
    {
        return $this->container['ref_type'];
    }

    /**
     * Sets ref_type
     *
     * @param string|null $ref_type ref_type
     *
     * @return self
     */
    public function setRefType($ref_type)
    {
        $this->container['ref_type'] = $ref_type;

        return $this;
    }

    /**
     * Gets ref_id
     *
     * @return int|null
     */
    public function getRefId()
    {
        return $this->container['ref_id'];
    }

    /**
     * Sets ref_id
     *
     * @param int|null $ref_id ref_id
     *
     * @return self
     */
    public function setRefId($ref_id)
    {
        $this->container['ref_id'] = $ref_id;

        return $this;
    }

    /**
     * Gets ref_name
     *
     * @return string|null
     */
    public function getRefName()
    {
        return $this->container['ref_name'];
    }

    /**
     * Sets ref_name
     *
     * @param string|null $ref_name ref_name
     *
     * @return self
     */
    public function setRefName($ref_name)
    {
        $this->container['ref_name'] = $ref_name;

        return $this;
    }

    /**
     * Gets type
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingType|null
     */
    public function getType()
    {
        return $this->container['type'];
    }

    /**
     * Sets type
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingType|null $type type
     *
     * @return self
     */
    public function setType($type)
    {
        $this->container['type'] = $type;

        return $this;
    }

    /**
     * Gets savings
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingSavings|null
     */
    public function getSavings()
    {
        return $this->container['savings'];
    }

    /**
     * Sets savings
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingSavings|null $savings savings
     *
     * @return self
     */
    public function setSavings($savings)
    {
        $this->container['savings'] = $savings;

        return $this;
    }

    /**
     * Gets resource
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingResource|null
     */
    public function getResource()
    {
        return $this->container['resource'];
    }

    /**
     * Sets resource
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingResource|null $resource resource
     *
     * @return self
     */
    public function setResource($resource)
    {
        $this->container['resource'] = $resource;

        return $this;
    }

    /**
     * Gets plan_before_action
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingPlanBeforeAction|null
     */
    public function getPlanBeforeAction()
    {
        return $this->container['plan_before_action'];
    }

    /**
     * Sets plan_before_action
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingPlanBeforeAction|null $plan_before_action plan_before_action
     *
     * @return self
     */
    public function setPlanBeforeAction($plan_before_action)
    {
        $this->container['plan_before_action'] = $plan_before_action;

        return $this;
    }

    /**
     * Gets plan_after_action
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingPlanAfterAction|null
     */
    public function getPlanAfterAction()
    {
        return $this->container['plan_after_action'];
    }

    /**
     * Sets plan_after_action
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingPlanAfterAction|null $plan_after_action plan_after_action
     *
     * @return self
     */
    public function setPlanAfterAction($plan_after_action)
    {
        $this->container['plan_after_action'] = $plan_after_action;

        return $this;
    }

    /**
     * Gets config
     *
     * @return \OpenAPI\Client\Model\GuidanceVmwareSizingConfig|null
     */
    public function getConfig()
    {
        return $this->container['config'];
    }

    /**
     * Sets config
     *
     * @param \OpenAPI\Client\Model\GuidanceVmwareSizingConfig|null $config config
     *
     * @return self
     */
    public function setConfig($config)
    {
        $this->container['config'] = $config;

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

