package k8s

import (
	"testing"
)

func TestGetNamespaceList(t *testing.T) {
	clientSet, _ := GetClientOutside()
	nsList, err := GetNamespaceList(clientSet)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(nsList)
}
func TestCreateNamespace(t *testing.T) {
	clientSet, _ := GetClientOutside()
	err := CreateNamespace(clientSet, "ns-test")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteNamespace(t *testing.T) {
	clientSet, _ := GetClientOutside()
	err := DeleteNamespace(clientSet, "ns-test")
	if err != nil {
		t.Error(err)
	}
}
func TestParchNodeLabel(t *testing.T) {
	clientSet, _ := GetClientOutside()
	label := map[string]interface{}{
		"k8s": "k8s-1",
	}
	err := ParchNodeLabel(clientSet, "k8s-master", label)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteNodeLabel(t *testing.T) {
	clientSet, _ := GetClientOutside()
	label := "k8s"
	err := DeleteNodeLabel(clientSet, "k8s-master", label)
	if err != nil {
		t.Error(err)
	}
}
func TestGetNodePodList(t *testing.T) {
	clientSet, _ := GetClientOutside()
	pods, err := GetNodePodList(clientSet, "default")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", pods)

}

func TestGetNodePod(t *testing.T) {
	clientSet, _ := GetClientOutside()
	pod, err := GetNodePod(clientSet, "default", "kubia-deployment-59cf6d5ccf-55b7t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", pod)
}
func TestGetServiceList(t *testing.T) {
	clientSet, _ := GetClientOutside()
	serviceList, err := GetServiceList(clientSet, "default")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", serviceList)
}
func TestGetNodeService(t *testing.T) {
	clientSet, _ := GetClientOutside()
	service, err := GetNodeService(clientSet, "default", "kubia-service")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", service)
}
