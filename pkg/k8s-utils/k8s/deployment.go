package k8s

import (
	"context"
	"k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"keikoproj.intuit.com/Flippy/pkg/k8s-utils/utils"
)

func (K8sWrapper) DeleteDeployment(clientset kubernetes.Interface, namespace string, deploymentName string) error {
	deploymentClient := clientset.AppsV1().Deployments(namespace)
	err := deploymentClient.Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (K8sWrapper) GetDeployments(clientset kubernetes.Interface, namespace string) (*v1.DeploymentList, error) {
	deploymentClient := clientset.AppsV1().Deployments(namespace)
	deployments, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	return deployments, err
}

func (K8sWrapper) GetDeployment(clientset kubernetes.Interface, namespace string, deploymentName string) (v1.Deployment, error) {
	var nilDeployment v1.Deployment
	deployments, err := K8s.GetDeployments(clientset, namespace)
	if err != nil {
		return nilDeployment, err
	}
	for _, deployment := range deployments.Items {
		if deployment.Name == deploymentName {
			return deployment, nil
		}
	}
	return nilDeployment, nil
}

func (K8sWrapper) GetDeploymentWithSpecAnnotationFilter(clientset kubernetes.Interface, namespace string, annotations map[string]string) ([]string, error) {

	deploymentList := make([]string, 0)
	deployments, err := K8s.GetDeployments(clientset, namespace)

	if err != nil {
		return deploymentList, err
	}

	for _, deployment := range deployments.Items {
		if utils.IsStringMapSubset(deployment.Spec.Template.ObjectMeta.Annotations, annotations) {
			deploymentList = append(deploymentList, deployment.Name)
		}
	}
	return deploymentList, nil
}
