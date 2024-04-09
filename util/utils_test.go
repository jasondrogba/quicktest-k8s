package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

//import (
//	. "github.com/onsi/ginkgo/v2"
//	. "github.com/onsi/gomega"
//	"os/exec"
//)
//
//var _ = Describe("Util", func() {
//	Describe("HelmInstallAlluxioOperator", func() {
//		Context("when HelmInstallAlluxioOperator is called", func() {
//			It("should execute helm install command successfully", func() {
//				output, err := HelmInstallAlluxioOperator()
//				Expect(err).NotTo(HaveOccurred())
//				Expect(output).To(ContainSubstring("deploy/charts/alluxio-operator"))
//			})
//		})
//	})
//
//	Describe("Run", func() {
//		Context("when Run is called with a valid command", func() {
//			It("should execute the command successfully", func() {
//				cmd := exec.Command("echo", "test")
//				output, err := Run(cmd)
//				Expect(err).NotTo(HaveOccurred())
//				Expect(string(output)).To(Equal("test\n"))
//			})
//		})
//	})
//
//	Describe("GetProjectDir", func() {
//		Context("when GetProjectDir is called", func() {
//			It("should return the project directory", func() {
//				dir, err := GetProjectDir()
//				Expect(err).NotTo(HaveOccurred())
//				Expect(dir).NotTo(BeEmpty())
//				Expect(dir).To(ContainSubstring("/tests/go/e2e")) // Modify this to match your expected project directory
//			})
//		})
//	})
//})

func TestHelmInstallAlluxioOperator(t *testing.T) {

	t.Run("should execute helm install command successfully", func(t *testing.T) {
		output, err := HelmInstallAlluxioOperator()
		if err != nil {
			t.Errorf("HelmInstallAlluxioOperator returned an error: %v", err)
		}
		//expectedOutput := "deploy/charts/alluxio-operator\n"
		//if output != expectedOutput {
		//	t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, string(output))
		//}
		fmt.Print(output)
	})
}
func TestHelmUninstallAlluxioOperator(t *testing.T) {

	t.Run("should execute helm install command successfully", func(t *testing.T) {
		output, err := HelmUninstallAlluxioOperator()
		if err != nil {
			t.Errorf("HelmInstallAlluxioOperator returned an error: %v", err)
		}
		//expectedOutput := "deploy/charts/alluxio-operator\n"
		//if output != expectedOutput {
		//	t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, string(output))
		//}
		fmt.Print(output)
	})
}

func TestKubectlApplyAlluxioCluster(t *testing.T) {
	t.Parallel()

	t.Run("should execute helm install command successfully", func(t *testing.T) {
		output, err := KubectlApplyAlluxioCluster()
		if err != nil {
			t.Errorf("HelmInstallAlluxioOperator returned an error: %v", err)
		}
		//expectedOutput := "deploy/charts/alluxio-operator\n"
		//if output != expectedOutput {
		//	t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, string(output))
		//}
		fmt.Print(output)
	})
}

func TestKubectlApplyAppPod(t *testing.T) {
	t.Run("should execute helm install command successfully", func(t *testing.T) {
		output, err := KubectlApplyAppPod()
		if err != nil {
			t.Errorf("HelmInstallAlluxioOperator returned an error: %v", err)
		}
		//expectedOutput := "deploy/charts/alluxio-operator\n"
		//if output != expectedOutput {
		//	t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, string(output))
		//}
		fmt.Print(output)
	})
}

func TestKubectlDeleteAlluxioCluster(t *testing.T) {
	t.Parallel()

	t.Run("should execute helm install command successfully", func(t *testing.T) {
		output, err := KubectlDeleteAlluxioCluster()
		if err != nil {
			t.Errorf("HelmInstallAlluxioOperator returned an error: %v", err)
		}
		//expectedOutput := "deploy/charts/alluxio-operator\n"
		//if output != expectedOutput {
		//	t.Errorf("Expected output to contain '%s', but got '%s'", expectedOutput, string(output))
		//}
		fmt.Print(output)
	})
}

func TestGetAlluxioFusePodStatus(t *testing.T) {
	t.Parallel()
	t.Run("should execute helm install command successfully", func(t *testing.T) {
		pods, _ := GetPodStatus()
		for _, pod := range pods.Items {
			//annotationsJSON, err := json.MarshalIndent(pod.GetAnnotations(), "", "  ")
			//if err != nil {
			//	fmt.Printf("Error marshalling annotations to JSON: %v\n", err)
			//	continue
			//}
			//fmt.Println(string(annotationsJSON))

			if podNameStartsWithAlluxioFuse(pod.GetName()) {
				fmt.Println("Alluxio-Fuse Pods:")
				fmt.Printf(" Pod Name: %s\n", pod.GetName())
				for key, value := range pod.GetAnnotations() {
					fmt.Printf("  %s: %s\n", key, value)
				}
				fmt.Println("-----------------------")
			}

			//podJSON, err := json.MarshalIndent(pod, "", "	")
			//if err != nil {
			//	fmt.Printf("Error marshalling pod to JSON: %v\n", err)
			//	continue
			//}
			//fmt.Println(string(podJSON))
			//fmt.Println(pod.String())
		}

	})
}

func TestRun(t *testing.T) {
	t.Parallel()

	t.Run("should execute the command successfully", func(t *testing.T) {
		cmd := exec.Command("echo", "test")
		output, err := Run(cmd)
		if err != nil {
			t.Errorf("Run returned an error: %v", err)
		}
		expectedOutput := "test\n"
		if string(output) != expectedOutput {
			t.Errorf("Expected output to be '%s', but got '%s'", expectedOutput, string(output))
		}
	})
}

func TestGetProjectDir(t *testing.T) {
	t.Parallel()

	t.Run("should return the project directory", func(t *testing.T) {
		dir, err := GetProjectDir()
		if err != nil {
			t.Errorf("GetProjectDir returned an error: %v", err)
		}
		expectedSubstring, err := os.Getwd()
		if !strings.Contains(dir, expectedSubstring) {
			t.Errorf("Expected directory to contain '%s', but got '%s'", expectedSubstring, dir)
		}
	})
}
