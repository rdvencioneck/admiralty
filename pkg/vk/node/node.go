/*
 * Copyright 2022 The Multicluster-Scheduler Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package node

import (
	"os"

	"admiralty.io/multicluster-scheduler/pkg/config/agent"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"admiralty.io/multicluster-scheduler/pkg/common"
	"admiralty.io/multicluster-scheduler/pkg/model/virtualnode"
)

func NodeFromOpts(t agent.Target) *v1.Node {
	node_address := os.Getenv("NODE_IP")
	if node_address == "" {
		node_address = os.Getenv("VKUBELET_POD_IP")
	}
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name:   t.VirtualNodeName,
			Labels: virtualnode.BaseLabels(t.Namespace, t.Name),
		},
		Spec: v1.NodeSpec{
			Taints: []v1.Taint{
				{
					Key:    common.LabelAndTaintKeyVirtualKubeletProvider,
					Value:  common.VirtualKubeletProviderName,
					Effect: v1.TaintEffectNoSchedule,
				},
			},
		},
		Status: v1.NodeStatus{
			Conditions: []v1.NodeCondition{
				{
					Type:               v1.NodeReady,
					Status:             v1.ConditionTrue,
					LastHeartbeatTime:  metav1.Now(),
					LastTransitionTime: metav1.Now(),
				},
				//{
				//	Type:               v1.NodeMemoryPressure,
				//	Status:             v1.ConditionFalse,
				//	LastHeartbeatTime:  metav1.Now(),
				//	LastTransitionTime: metav1.Now(),
				//},
				//{
				//	Type:               v1.NodeDiskPressure,
				//	Status:             v1.ConditionFalse,
				//	LastHeartbeatTime:  metav1.Now(),
				//	LastTransitionTime: metav1.Now(),
				//},
				//{
				//	Type:               v1.NodePIDPressure,
				//	Status:             v1.ConditionFalse,
				//	LastHeartbeatTime:  metav1.Now(),
				//	LastTransitionTime: metav1.Now(),
				//},
				//{
				//	Type:               v1.NodeNetworkUnavailable,
				//	Status:             v1.ConditionFalse,
				//	LastHeartbeatTime:  metav1.Now(),
				//	LastTransitionTime: metav1.Now(),
				//},
			},
			Addresses: []v1.NodeAddress{
				{
					Type:    "InternalIP",
					Address: node_address,
				},
			},
			//DaemonEndpoints: v1.NodeDaemonEndpoints{
			//	KubeletEndpoint: v1.DaemonEndpoint{
			//		Port: int32(c.ListenPort),
			//	},
			//},
		},
	}
	return node
}
