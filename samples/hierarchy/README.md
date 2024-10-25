# CA Hierarchy Sample

This sample demonstrates how to set up a CA hierarchy using the AWS Controllers for Kubernetes (ACK) service controller for AWS Private Certificate Authority.

Apply the manifest, `ca_hierarchy.yaml`, to your Kubernetes cluster using the following command:

```
kubectl apply -f ca_hierarchy.yaml
```

Applying this manifest should create the following resources:
- Root `CertificateAuthority` resource
- Root CA `Certificate` resource
- Root CA `CertificateAuthorityActivation` resource
- 2 `Secret` resources, storing the root CA certificate and certificate chain
- Subordinate `CertificateAuthority` resource
- Subordinate CA `Certificate` resource
- Subordinate CA `CertificateAuthorityActivation` resource
- 2 `Secret` resources, storing the subordinate CA certificate and certificate chain
- End entity `Certificate` resource
- `Secret` resource, storing the end entity certificate. 

The following commands will describe the resources.
```
kubectl describe certificateAuthority/root-ca
```
```
kubectl describe certificate/root-ca-certificate
```
```
kubectl describe certificateAuthorityActivation/root-ca-activation
```
```
kubectl describe certificateAuthority/sub-ca
```
```
kubectl describe certificate/sub-ca-certificate
```
```
kubectl describe certificateAuthorityActivation/sub-ca-activation
```
```
kubectl describe certificate/end-entity-certificate
```
```
kubectl describe certificateauthority/sub-ca
```