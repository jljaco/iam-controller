// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package open_id_connect_provider

import (
	"bytes"
	"reflect"
	"strings"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	compareTags(delta, a, b)

	if !ackcompare.SliceStringPEqual(a.ko.Spec.ClientIDList, b.ko.Spec.ClientIDList) {
		delta.Add("Spec.ClientIDList", a.ko.Spec.ClientIDList, b.ko.Spec.ClientIDList)
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.ThumbprintList, b.ko.Spec.ThumbprintList) {
		delta.Add("Spec.ThumbprintList", a.ko.Spec.ThumbprintList, b.ko.Spec.ThumbprintList)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.URL, b.ko.Spec.URL) {
		delta.Add("Spec.URL", a.ko.Spec.URL, b.ko.Spec.URL)
	} else if a.ko.Spec.URL != nil && b.ko.Spec.URL != nil {
		// the URL field must begin with "https://"
		// c.f. https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateOpenIDConnectProvider.html
		// however, when stored in the IAM backend, the "https://" prefix is stripped out
		// thus here we treat "https://" as a null token for the purposes of string comparison
		if (*a.ko.Spec.URL != *b.ko.Spec.URL) {
			a_url := strings.TrimPrefix(*a.ko.Spec.URL, "https://")
			b_url := strings.TrimPrefix(*b.ko.Spec.URL, "https://")
			if a_url != b_url {
				delta.Add("Spec.URL", a.ko.Spec.URL, b.ko.Spec.URL)
		   }
		}
	}

	return delta
}
