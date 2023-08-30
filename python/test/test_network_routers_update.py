"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import openapi_client
from openapi_client.model.network_routers_update_network_server import NetworkRoutersUpdateNetworkServer
from openapi_client.model.network_routers_update_site import NetworkRoutersUpdateSite
from openapi_client.model.network_routers_update_type import NetworkRoutersUpdateType
from openapi_client.model.network_routers_update_zone import NetworkRoutersUpdateZone
globals()['NetworkRoutersUpdateNetworkServer'] = NetworkRoutersUpdateNetworkServer
globals()['NetworkRoutersUpdateSite'] = NetworkRoutersUpdateSite
globals()['NetworkRoutersUpdateType'] = NetworkRoutersUpdateType
globals()['NetworkRoutersUpdateZone'] = NetworkRoutersUpdateZone
from openapi_client.model.network_routers_update import NetworkRoutersUpdate


class TestNetworkRoutersUpdate(unittest.TestCase):
    """NetworkRoutersUpdate unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testNetworkRoutersUpdate(self):
        """Test NetworkRoutersUpdate"""
        # FIXME: construct object with mandatory attributes with example values
        # model = NetworkRoutersUpdate()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()