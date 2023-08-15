package k8s

import (
	"base_k8s/cloud/global"
	"context"
	"flag"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetClientOutside() (*kubernetes.Clientset, error) {
	// 支持以 Pod 形式或者在宿主机上运行代码的形式获取 kubeconfig 配置
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "./kube/config", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, err

	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}

// GetClientSetInside （当应用在集群中，比如 Pod 内使用时调用此函数）获取 clientSet 对象，以对 k8s 中的资源进行 CURD
func GetClientSetInside() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientSet
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

// GetNamespaceList 获取namespace列表
func GetNamespaceList(clientSet *kubernetes.Clientset) (*corev1.NamespaceList, error) {
	namespaceList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaceList, nil

}

// CreateNamespace 创建k8s的namespace
func CreateNamespace(clientSet *kubernetes.Clientset, namespace string) error {
	_, err := clientSet.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		nsSpec := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}
		_, err := clientSet.CoreV1().Namespaces().Create(context.TODO(), nsSpec, metav1.CreateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteNamespace 删除k8s的namespace
func DeleteNamespace(clientSet *kubernetes.Clientset, namespace string) error {
	err := clientSet.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
} // PatchNodeLabel 为某个节点新增/更新标签, 需要指定节点的名称
func ParchNodeLabel(clientSet *kubernetes.Clientset, nodeName string, labels map[string]interface{}) error {
	patchTemplate := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": labels,
		},
	}
	// 通过 Patch 操作实现对某个节点的 Label 的新建或者更新。
	patchData, _ := json.Marshal(patchTemplate)
	_, err := clientSet.CoreV1().Nodes().Patch(context.TODO(), nodeName, types.StrategicMergePatchType, patchData, metav1.PatchOptions{})
	if err != nil {
		return err
	}
	return nil
}

// DeleteNodeLabel 为某个节点删除某标签，需要指定节点的名称以及标签的 key
func DeleteNodeLabel(clientSet *kubernetes.Clientset, nodeName, labelKey string) error {
	patchTemplate := map[string]interface{}{
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				labelKey: nil,
			},
		},
	}
	// 通过 Patch 操作实现对某个节点的 Label 的删除。这是声明式的。
	patchData, _ := json.Marshal(patchTemplate)
	_, err := clientSet.CoreV1().Nodes().Patch(context.TODO(), nodeName, types.StrategicMergePatchType, patchData, metav1.PatchOptions{})
	if err != nil {
		return err
	}
	return nil
}

// ApplyConfigMapWithData 利用 Apply 方式去创建/更新 configmap 中的内容，需要给定 data 数据
func ApplyConfigMapWithData(clientSet *kubernetes.Clientset, configMapName string, namespace string, data map[string]string) (*corev1.ConfigMap, error) {
	cm := v1.ConfigMap(configMapName, global.CONFIG.K8s.Namespace)
	cm.WithData(data)
	resultCM, err := clientSet.CoreV1().ConfigMaps(namespace).Apply(context.TODO(), cm, metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		return nil, err
	}
	return resultCM, nil
}

// GetNodePodList 获取pod列表
func GetNodePodList(clientSet *kubernetes.Clientset, namespace string) (*corev1.PodList, error) {
	podList, err := clientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return podList, nil
}

// GetNodePod 获取某一个pod
func GetNodePod(clientSet *kubernetes.Clientset, namespace, name string) (*corev1.Pod, error) {
	pod, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}

// GetServiceList 获取Service列表
func GetServiceList(clientSet *kubernetes.Clientset, namespace string) (*corev1.ServiceList, error) {
	serviceList, err := clientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return serviceList, err
}

// GetNodeService 获取某一个Service
func GetNodeService(clientSet *kubernetes.Clientset, namespace, name string) (*corev1.Service, error) {
	service, err := clientSet.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return service, nil
}
