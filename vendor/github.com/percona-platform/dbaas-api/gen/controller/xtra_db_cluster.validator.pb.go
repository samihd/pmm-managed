// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/xtra_db_cluster.proto

package controllerv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *XtraDBClusterParams) Validate() error {
	if !(this.ClusterSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ClusterSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.ClusterSize))
	}
	if nil == this.Pxc {
		return github_com_mwitkow_go_proto_validators.FieldError("Pxc", fmt.Errorf("message must exist"))
	}
	if this.Pxc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pxc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pxc", err)
		}
	}
	if nil == this.Proxysql {
		return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", fmt.Errorf("message must exist"))
	}
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	return nil
}
func (this *XtraDBClusterParams_PXC) Validate() error {
	if nil == this.ComputeResources {
		return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", fmt.Errorf("message must exist"))
	}
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
func (this *XtraDBClusterParams_ProxySQL) Validate() error {
	if nil == this.ComputeResources {
		return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", fmt.Errorf("message must exist"))
	}
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
func (this *XtraDBCredentials) Validate() error {
	return nil
}
