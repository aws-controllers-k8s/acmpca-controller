# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Helper functions for acmpca tests
"""

from typing import Union, Dict


class ACMPCAValidator:
    def __init__(self, acmpca_client):
        self.acmpca_client = acmpca_client

    def assert_certificate_authority(self, ca_arn: str, status: str, exists=True):
        try:
            aws_res = self.acmpca_client.describe_certificate_authority(CertificateAuthorityArn=ca_arn)
            assert aws_res["Status"] == status
            assert aws_res["Type"] == "ROOT"
            assert aws_res["CertificateAuthorityConfiguration"]["KeyAlgorithm"] == "RSA_2048"
            assert aws_res["CertificateAuthorityConfiguration"]["SigningAlgorithm"] == "SHA256WITHRSA"
            assert aws_res["CertificateAuthorityConfiguration"]["Subject"]["Organization"] == "Example Organization"
            assert aws_res["CertificateAuthorityConfiguration"]["Subject"]["OrganizationalUnit"] == "Example"
            assert aws_res["CertificateAuthorityConfiguration"]["Subject"]["Country"] == "US"
            assert aws_res["CertificateAuthorityConfiguration"]["Subject"]["State"] == "Virginia"
            assert aws_res["CertificateAuthorityConfiguration"]["Subject"]["Locality"] == "Arlington"
            assert aws_res["CertificateAuthorityConfiguration"]["Subject"]["CommonName"] == "www.example.com"
        except self.acmpca_client.exceptions.ClientError:
            pass
        try:
            aws_res = self.acmpca_client.list_tags(CertificateAuthorityArn=ca_arn)
            assert aws_res["Tags"]["Key"] == "Name"
            assert aws_res["Tags"]["Value"] == "Test CA"
        except self.acmpca_client.exceptions.ClientError:
            pass
    
    def assert_csr(self, ca_arn: str, exists=True):
        res_found = False
        try:
            aws_res = self.acmpca_client.get_certificate_authority_csr(CertificateAuthorityArn=ca_arn)
            res_found = aws_res["Csr"] != None
        except self.acmpca_client.exceptions.ClientError:
            pass
        assert res_found is exists

    def assert_certificate(self, ca_arn: str, c_arn: str, exists=True):
        res_found = False
        try:
            aws_res = self.acmpca_client.get_certificate(CertificateAuthorityArn=ca_arn, CertificateArn=c_arn)
            res_found = aws_res["Certificate"] != None
        except self.acmpca_client.exceptions.ClientError:
            pass
        assert res_found is exists 
