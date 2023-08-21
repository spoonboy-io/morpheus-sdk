package org.openapitools.model;

import groovy.transform.Canonical
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.math.BigDecimal;
import java.util.ArrayList;
import java.util.List;
import org.openapitools.model.ApiVdiPoolsIdVdiPoolConfig;

@Canonical
class ApiVdiPoolsIdVdiPool {
    /* Virtual Desktop name */
    String name
    /* Virtual Desktop description */
    String description
    /* Owner (User) ID */
    Long owner
    /* Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.  */
    BigDecimal minIdle
    /* The initial size of the pool to be allocated on creation */
    BigDecimal initialPoolSize
    /* Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances  */
    BigDecimal maxIdle
    /* Max limit on number of allocations and instances within the pool.  */
    BigDecimal maxPoolSize
    /* Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.  */
    BigDecimal allocationTimeoutMinutes
    /* Persistent Desktop Pool */
    Boolean persistentUser = false
    /* Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) */
    Boolean recyclable = false
    /* Allow copy from desktop */
    Boolean allowCopy = false
    /* Allow local printers from Desktop */
    Boolean allowPrinter = false
    /* Allow File Share */
    Boolean allowFileshare = false
    /* Allow hypervisor console */
    Boolean allowHypervisorConsole = false
    /* Auto create local user upon reservation */
    Boolean autoCreateLocalUserOnReservation = false
    /* Can be used to enable or disable the VDI pool */
    Boolean enabled = true
    /* The relative location of an icon image */
    String iconPath
    /* Array of VDI App IDs */
    List<Long> apps = new ArrayList<Long>()
    /* VDI Gateway ID */
    Long gateway
    /* Instance Config JSON. Passing as a string will preserve property order.  See `config` object for required values. */
    String instanceConfig
    
    ApiVdiPoolsIdVdiPoolConfig config
    /* Guest Console Jump Host */
    String guestConsoleJumpHost
    /* Guest Console Jump Port */
    Long guestConsoleJumpPort
    /* Guest Console Jump Username */
    String guestConsoleJumpUsername
    /* Guest Console Jump Password */
    String guestConsoleJumpPassword
    /* Guest Console Jump Key Pair. see `Key Pair` */
    Long guestConsoleJumpKeypair
}
