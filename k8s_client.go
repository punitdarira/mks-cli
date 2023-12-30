package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func connectCluster(deployment Deployment) {

	// Create a Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", getKubeconfigPath())
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v", err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %v", err)
		return
	}

	// Create the deployment

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: deployment.App,
		},
	}
	_, err = clientset.CoreV1().Namespaces().Create(context.TODO(), ns, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating namespace: %v", err)
		return
	}

	createDeployment(deployment, clientset)
	createService(deployment, clientset)
	createGrafanaDashboard(deployment)
}

// getKubeconfigPath returns the path to the kubeconfig file.
func getKubeconfigPath() string {
	//home := homedir.HomeDir()
	//return home + "/.kube/config"
	return "C:\\MKS\\admin.conf"
}
