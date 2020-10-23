// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/dbaas/psmdb_cluster.proto

package dbaasv1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PSMDBClusterParams) Validate() error {
	if !(this.ClusterSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.ClusterSize))
	}
	if nil == this.Replicaset {
		return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", fmt.Errorf("message must exist"))
	}
	if this.Replicaset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Replicaset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", err)
		}
	}
	return nil
}
func (this *PSMDBClusterParams_ReplicaSet) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	if !(this.DiskSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("DiskSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.DiskSize))
	}
	return nil
}
func (this *ListPSMDBClustersRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	return nil
}
func (this *ListPSMDBClustersResponse) Validate() error {
	for _, item := range this.Clusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Clusters", err)
			}
		}
	}
	return nil
}
func (this *ListPSMDBClustersResponse_Cluster) Validate() error {
	if this.Operation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operation", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *GetPSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *GetPSMDBClusterResponse) Validate() error {
	if this.Operation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operation", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	if this.ConnectionCredentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionCredentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionCredentials", err)
		}
	}
	return nil
}
func (this *GetPSMDBClusterResponse_PSMDBCredentials) Validate() error {
	return nil
}

var _regex_CreatePSMDBClusterRequest_Name = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)

func (this *CreatePSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if !_regex_CreatePSMDBClusterRequest_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]([-a-z0-9]*[a-z0-9])?$"`, this.Name))
	}
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *CreatePSMDBClusterResponse) Validate() error {
	return nil
}
func (this *UpdatePSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *UpdatePSMDBClusterRequest_UpdatePSMDBClusterParams) Validate() error {
	if !(this.ClusterSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.ClusterSize))
	}
	if nil == this.Replicaset {
		return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", fmt.Errorf("message must exist"))
	}
	if this.Replicaset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Replicaset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Replicaset", err)
		}
	}
	return nil
}
func (this *UpdatePSMDBClusterRequest_UpdatePSMDBClusterParams_ReplicaSet) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdatePSMDBClusterResponse) Validate() error {
	return nil
}
func (this *DeletePSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *DeletePSMDBClusterResponse) Validate() error {
	return nil
}
func (this *RestartPSMDBClusterRequest) Validate() error {
	if this.KubernetesClusterName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("KubernetesClusterName", fmt.Errorf(`value '%v' must not be an empty string`, this.KubernetesClusterName))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *RestartPSMDBClusterResponse) Validate() error {
	return nil
}