"""
    Morpheus API

    Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.  # noqa: E501

    The version of the OpenAPI document: 6.2.1
    Contact: dev@morpheusdata.com
    Generated by: https://openapi-generator.tech
"""


import unittest

import openapi_client
from openapi_client.api.logs_api import LogsApi  # noqa: E501


class TestLogsApi(unittest.TestCase):
    """LogsApi unit test stubs"""

    def setUp(self):
        self.api = LogsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_list_logs(self):
        """Test case for list_logs

        Retrieves Logs  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()