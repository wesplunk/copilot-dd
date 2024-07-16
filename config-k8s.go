package main

import (
    "context"
    "fmt"
    "path/filepath"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/util/homedir"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/apimachinery/pkg/types"
)

func scaleDeployment(namespace, deploymentName string, replicas int32) {
    // Use the current user's home directory to find the kubeconfig file
    kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

    // Build the configuration from the kubeconfig file
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(err.Error())
    }

    // Create a new clientset for the given config
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }

    // Define the patch body to update the replicas
    patch := fmt.Sprintf(`{"spec":{"replicas":%d}}`, replicas)

    // Scale the deployment
    _, err = clientset.AppsV1().Deployments(namespace).Patch(context.TODO(), deploymentName, types.StrategicMergePatchType, []byte(patch), metav1.PatchOptions{})
    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("Deployment %s scaled to %d replicas.\n", deploymentName, replicas)
}

func main() {
    // Example usage
    scaleDeployment("default", "myapp-deployment", 5)
}