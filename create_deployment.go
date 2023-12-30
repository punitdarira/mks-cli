package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

func createDeployment(deploymentFromUser Deployment, clientset *kubernetes.Clientset) {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentFromUser.App,
			Namespace: deploymentFromUser.App,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentFromUser.Replica,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentFromUser.App,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentFromUser.App,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentFromUser.App,
							Image: deploymentFromUser.Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: deploymentFromUser.Port,
								},
							},
						},
					},
				},
			},
		},
	}

	result, err := clientset.AppsV1().Deployments(deploymentFromUser.App).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating deployment: %v", err)
		return
	}

	fmt.Printf("Deployment created: %s\n", result.GetObjectMeta().GetName())
}

func createService(deploymentFromUser Deployment, clientset *kubernetes.Clientset) {

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentFromUser.App,
			Namespace: deploymentFromUser.App,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeLoadBalancer,
			Selector: map[string]string{"app": deploymentFromUser.App},
			Ports: []corev1.ServicePort{
				{
					Port:       80,
					TargetPort: intstr.FromInt(int(deploymentFromUser.Port)),
				},
			},
		},
	}

	result, err := clientset.CoreV1().Services(deploymentFromUser.App).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating service: %v", err)
		return
	}

	fmt.Printf("Service created: %s\n", result.GetObjectMeta().GetName())
}

func createGrafanaDashboard(deploymentFromUser Deployment) {
	grafanaURL := "http://138.197.154.178:32000/api/dashboards/db"
	grafanaAPIKey := "glsa_N98QvHwSqNhhRE27foPVaXAd7aOCSYxS_f66eff97"

	data, err := ioutil.ReadFile("C:\\MKS\\dashboard.json")
	if err != nil {
		panic(err)
	}

	modifiedData := strings.ReplaceAll(string(data), "{{Namespace}}", deploymentFromUser.App)

	req, _ := http.NewRequest("POST", grafanaURL, bytes.NewBuffer([]byte(modifiedData)))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+grafanaAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error creating Grafana dashboard: %v", err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Grafana dashboard created for namespace: %s\n", deploymentFromUser.App)
}
