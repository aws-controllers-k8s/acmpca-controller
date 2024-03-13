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

import logging

class ACMPCAValidator:
    def __init__(self, acmpca_client):
        self.acmpca_client = acmpca_client

    def assert_certificate_authority(self, ca_arn: str, status: str):
        try:
            aws_res = self.acmpca_client.describe_certificate_authority(CertificateAuthorityArn=ca_arn)
            ca = aws_res["CertificateAuthority"]
            logging.info(ca["Status"])
            assert ca["Status"] == status
            return ca
        except self.acmpca_client.exceptions.ClientError:
            pass

    def get_ca_tags(self, ca_arn: str):
        try:
            aws_res = self.acmpca_client.list_tags(CertificateAuthorityArn=ca_arn)
            ca_tags = aws_res["Tags"]
            logging.info(ca_tags)
            return ca_tags
        except self.acmpca_client.exceptions.ClientError:
            pass

    def get_csr(self, ca_arn: str):
        try:
            aws_res = self.acmpca_client.get_certificate_authority_csr(CertificateAuthorityArn=ca_arn)
            csr = aws_res["Csr"]
            assert csr is not None
            return csr
        except self.acmpca_client.exceptions.ClientError:
            pass
    
    def get_certificate(self, ca_arn: str, cert_arn: str):
        try:
            aws_res = self.acmpca_client.get_certificate(CertificateAuthorityArn=ca_arn, CertificateArn=cert_arn)
            certificate = aws_res["Certificate"]
            assert certificate is not None
            return certificate
        except self.acmpca_client.exceptions.ClientError:
            pass