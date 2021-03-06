/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	"testing"

	"k8s.io/utils/pointer"
)

func TestAWSMachine_ValidateUpdate(t *testing.T) {
	tests := []struct {
		name       string
		oldMachine *AWSMachine
		newMachine *AWSMachine
		wantErr    bool
	}{
		{
			name: "change in providerid, tags and securitygroups",
			oldMachine: &AWSMachine{
				Spec: AWSMachineSpec{
					ProviderID:               nil,
					AdditionalTags:           nil,
					AdditionalSecurityGroups: nil,
				},
			},
			newMachine: &AWSMachine{
				Spec: AWSMachineSpec{
					ProviderID: pointer.StringPtr("ID"),
					AdditionalTags: Tags{
						"key-1": "value-1",
					},
					AdditionalSecurityGroups: []AWSResourceReference{
						{
							ID: pointer.StringPtr("ID"),
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "change in fields other than providerid, tags and securitygroups",
			oldMachine: &AWSMachine{
				Spec: AWSMachineSpec{
					ProviderID:               nil,
					AdditionalTags:           nil,
					AdditionalSecurityGroups: nil,
				},
			},
			newMachine: &AWSMachine{
				Spec: AWSMachineSpec{
					ImageLookupOrg: "test",
					InstanceType:   "test",
					ProviderID:     pointer.StringPtr("ID"),
					AdditionalTags: Tags{
						"key-1": "value-1",
					},
					AdditionalSecurityGroups: []AWSResourceReference{
						{
							ID: pointer.StringPtr("ID"),
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.newMachine.ValidateUpdate(tt.oldMachine); (err != nil) != tt.wantErr {
				t.Errorf("ValidateUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
