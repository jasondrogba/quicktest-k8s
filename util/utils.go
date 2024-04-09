package util

import (
	"context"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func HelmInstallAlluxioOperator() (string, error) {
	//alluxioOperatorPath := "deploy/charts/alluxio-operator"
	alluxioOperatorPath := "/Users/sunbury/alluxio-operator/alluxio-1.0.6/alluxio-operator"
	//cmd := exec.Command("helm", "install", "alluxio-operator", alluxioOperatorPath)
	cmd := exec.Command("helm", "install", "operator", alluxioOperatorPath)
	output, err := Run(cmd)
	return string(output), err
}

func HelmUninstallAlluxioOperator() (string, error) {

	cmd := exec.Command("helm", "uninstall", "operator")
	output, err := Run(cmd)
	return string(output), err
}

func KubectlApplyAlluxioCluster() (string, error) {
	datasetNullPath := "/Users/sunbury/alluxio-operator/alluxio-operator_1.0.4/dataset.yaml"
	etcdPVPath := "/Users/sunbury/alluxio-operator/alluxio-operator_1.0.4/etcd-pv.yaml"
	alluxioCluserPath := "/Users/sunbury/alluxio-operator/alluxio-1.0.6/alluxio-cluster-jiasen.yaml"
	//cmd := exec.Command("helm", "install", "alluxio-operator", alluxioOperatorPath)

	cmdDataset := exec.Command("kubectl", "apply", "-f", datasetNullPath)
	_, err := Run(cmdDataset)
	if err != nil {
		fmt.Print(err)
	}

	cmdETCDPV := exec.Command("kubectl", "apply", "-f", etcdPVPath)
	_, err = Run(cmdETCDPV)
	if err != nil {
		fmt.Print(err)
	}

	cmd := exec.Command("kubectl", "apply", "-f", alluxioCluserPath)
	output, err := Run(cmd)
	return string(output), err
}

func KubectlApplyAppPod() (string, error) {
	appPath := "/Users/sunbury/alluxio-operator/alluxio-operator_1.0.4/ds-XHS.yaml"
	cmdETCDPV := exec.Command("kubectl", "apply", "-f", appPath)
	output, err := Run(cmdETCDPV)
	if err != nil {
		fmt.Print(err)
	}
	return string(output), err
}

func KubectlDeleteAlluxioCluster() (string, error) {
	alluxioCluserPath := "/Users/sunbury/alluxio-operator/alluxio-operator_1.0.4/alluxio-cluster-jiasen.yaml"
	//cmd := exec.Command("helm", "install", "alluxio-operator", alluxioOperatorPath)
	cmd := exec.Command("kubectl", "delete", "-f", alluxioCluserPath)
	output, err := Run(cmd)
	return string(output), err
}

// Run executes the provided command within this context
func Run(cmd *exec.Cmd) ([]byte, error) {
	dir, _ := GetProjectDir()
	cmd.Dir = dir
	fmt.Fprintf(ginkgo.GinkgoWriter, "running dir: %s\n", cmd.Dir)

	// To allow make commands be executed from the project directory which is subdir on SDK repo
	// TODO:(user) You might not need the following code
	if err := os.Chdir(cmd.Dir); err != nil {
		fmt.Fprintf(ginkgo.GinkgoWriter, "chdir dir: %s\n", err)
	}

	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	command := strings.Join(cmd.Args, " ")
	fmt.Fprintf(ginkgo.GinkgoWriter, "running: %s\n", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, fmt.Errorf("%s failed with error: (%v) %s", command, err, string(output))
	}

	return output, nil
}

// GetProjectDir will return the directory where the project is
func GetProjectDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, err
	}
	wd = strings.Replace(wd, "/tests/go/e2e", "", -1)
	return wd, nil
}

// GetPodStatus will return pod status
func GetPodStatus() (*v1.PodList, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// get all pods
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pods, err
}

// DeletePod delete Pod
func DeletePod() {

}

func podNameStartsWithAlluxioFuse(name string) bool {
	prefix := "alluxio-fuse"
	return len(name) >= len(prefix) && name[:len(prefix)] == prefix
}
