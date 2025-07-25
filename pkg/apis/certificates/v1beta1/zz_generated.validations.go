//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by validation-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"
	fmt "fmt"

	certificatesv1beta1 "k8s.io/api/certificates/v1beta1"
	equality "k8s.io/apimachinery/pkg/api/equality"
	operation "k8s.io/apimachinery/pkg/api/operation"
	safe "k8s.io/apimachinery/pkg/api/safe"
	validate "k8s.io/apimachinery/pkg/api/validate"
	runtime "k8s.io/apimachinery/pkg/runtime"
	field "k8s.io/apimachinery/pkg/util/validation/field"
)

func init() { localSchemeBuilder.Register(RegisterValidations) }

// RegisterValidations adds validation functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterValidations(scheme *runtime.Scheme) error {
	scheme.AddValidationFunc((*certificatesv1beta1.CertificateSigningRequest)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/", "/approval", "/status":
			return Validate_CertificateSigningRequest(ctx, op, nil /* fldPath */, obj.(*certificatesv1beta1.CertificateSigningRequest), safe.Cast[*certificatesv1beta1.CertificateSigningRequest](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	scheme.AddValidationFunc((*certificatesv1beta1.CertificateSigningRequestList)(nil), func(ctx context.Context, op operation.Operation, obj, oldObj interface{}) field.ErrorList {
		switch op.Request.SubresourcePath() {
		case "/":
			return Validate_CertificateSigningRequestList(ctx, op, nil /* fldPath */, obj.(*certificatesv1beta1.CertificateSigningRequestList), safe.Cast[*certificatesv1beta1.CertificateSigningRequestList](oldObj))
		}
		return field.ErrorList{field.InternalError(nil, fmt.Errorf("no validation found for %T, subresource: %v", obj, op.Request.SubresourcePath()))}
	})
	return nil
}

func Validate_CertificateSigningRequest(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *certificatesv1beta1.CertificateSigningRequest) (errs field.ErrorList) {
	// field certificatesv1beta1.CertificateSigningRequest.TypeMeta has no validation
	// field certificatesv1beta1.CertificateSigningRequest.ObjectMeta has no validation
	// field certificatesv1beta1.CertificateSigningRequest.Spec has no validation

	// field certificatesv1beta1.CertificateSigningRequest.Status
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj *certificatesv1beta1.CertificateSigningRequestStatus) (errs field.ErrorList) {
			errs = append(errs, Validate_CertificateSigningRequestStatus(ctx, op, fldPath, obj, oldObj)...)
			return
		}(fldPath.Child("status"), &obj.Status, safe.Field(oldObj, func(oldObj *certificatesv1beta1.CertificateSigningRequest) *certificatesv1beta1.CertificateSigningRequestStatus {
			return &oldObj.Status
		}))...)

	return errs
}

func Validate_CertificateSigningRequestList(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *certificatesv1beta1.CertificateSigningRequestList) (errs field.ErrorList) {
	// field certificatesv1beta1.CertificateSigningRequestList.TypeMeta has no validation
	// field certificatesv1beta1.CertificateSigningRequestList.ListMeta has no validation

	// field certificatesv1beta1.CertificateSigningRequestList.Items
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj []certificatesv1beta1.CertificateSigningRequest) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			errs = append(errs, validate.EachSliceVal(ctx, op, fldPath, obj, oldObj, nil, nil, Validate_CertificateSigningRequest)...)
			return
		}(fldPath.Child("items"), obj.Items, safe.Field(oldObj, func(oldObj *certificatesv1beta1.CertificateSigningRequestList) []certificatesv1beta1.CertificateSigningRequest {
			return oldObj.Items
		}))...)

	return errs
}

var zeroOrOneOfMembershipFor_k8s_io_api_certificates_v1beta1_CertificateSigningRequestStatus_Conditions_ = validate.NewUnionMembership([2]string{"Conditions[{\"type\": \"Approved\"}]", ""}, [2]string{"Conditions[{\"type\": \"Denied\"}]", ""})

func Validate_CertificateSigningRequestStatus(ctx context.Context, op operation.Operation, fldPath *field.Path, obj, oldObj *certificatesv1beta1.CertificateSigningRequestStatus) (errs field.ErrorList) {
	// field certificatesv1beta1.CertificateSigningRequestStatus.Conditions
	errs = append(errs,
		func(fldPath *field.Path, obj, oldObj []certificatesv1beta1.CertificateSigningRequestCondition) (errs field.ErrorList) {
			if op.Type == operation.Update && equality.Semantic.DeepEqual(obj, oldObj) {
				return nil // no changes
			}
			if e := validate.OptionalSlice(ctx, op, fldPath, obj, oldObj); len(e) != 0 {
				return // do not proceed
			}
			errs = append(errs, validate.ZeroOrOneOfUnion(ctx, op, fldPath, obj, oldObj, zeroOrOneOfMembershipFor_k8s_io_api_certificates_v1beta1_CertificateSigningRequestStatus_Conditions_, func(list []certificatesv1beta1.CertificateSigningRequestCondition) bool {
				for i := range list {
					if list[i].Type == "Approved" {
						return true
					}
				}
				return false
			}, func(list []certificatesv1beta1.CertificateSigningRequestCondition) bool {
				for i := range list {
					if list[i].Type == "Denied" {
						return true
					}
				}
				return false
			})...)
			return
		}(fldPath.Child("conditions"), obj.Conditions, safe.Field(oldObj, func(oldObj *certificatesv1beta1.CertificateSigningRequestStatus) []certificatesv1beta1.CertificateSigningRequestCondition {
			return oldObj.Conditions
		}))...)

	// field certificatesv1beta1.CertificateSigningRequestStatus.Certificate has no validation
	return errs
}
