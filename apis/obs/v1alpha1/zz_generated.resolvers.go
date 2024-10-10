// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BucketObject.
func (mg *BucketObject) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.ForProvider.BucketRef,
		Selector:     mg.Spec.ForProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
		Extract:      resource.ExtractParamPath("bucket", true),
		Reference:    mg.Spec.InitProvider.BucketRef,
		Selector:     mg.Spec.InitProvider.BucketSelector,
		To: reference.To{
			List:    &BucketList{},
			Managed: &Bucket{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	return nil
}
