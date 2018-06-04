/*
  Inspired by:
  - https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet_getters.go
*/

package agent

import (
	nodeutil "k8s.io/kubernetes/pkg/util/node"
)

// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet_getters.go#187

// GetHostname Returns the hostname as the kubelet sees it.
//func (kl *Kubelet) GetHostname() string {
//	return kl.hostname
func (a *Agent) GetHostname() string {
	hostnameOverride := ""
	hostname := nodeutil.GetHostname(hostnameOverride)
	if a.hostname == "" {
		a.hostname = hostname
	}
	return hostname
}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet_getters.go#197

// GetNode returns the node info for the configured node name of this Kubelet.
//func (kl *Kubelet) GetNode() (*v1.Node, error) {
//	if kl.kubeClient == nil {
//		return kl.initialNode()
//	}
//	return kl.nodeInfo.GetNodeInfo(string(kl.nodeName))
func (a *Agent) GetNode() (*v1.Node, error) {

}

// Refer to https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/kubelet_getters.go#205

// getNodeAnyWay() must return a *v1.Node which is required by RunGeneralPredicates().
// The *v1.Node is obtained as follows:
// Return kubelet's nodeInfo for this node, except on error or if in standalone mode,
// in which case return a manufactured nodeInfo representing a node with no pods,
// zero capacity, and the default labels.
func (kl *Kubelet) getNodeAnyWay() (*v1.Node, error) {
	if kl.kubeClient != nil {
		if n, err := kl.nodeInfo.GetNodeInfo(string(kl.nodeName)); err == nil {
			return n, nil
		}
	}
	return kl.initialNode()
}
