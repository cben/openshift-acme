package util

import (
	"os"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/tnozicka/openshift-acme/test/e2e/framework"
)

var TestContext *framework.TestContextType = &framework.TestContext

func KubeConfigPath() string {
	return os.Getenv("KUBECONFIG")
}

func InitTest() {
	TestContext.KubeConfigPath = KubeConfigPath()
	framework.Logf("KubeConfigPath: %q", TestContext.KubeConfigPath)
	if TestContext.KubeConfigPath == "" {
		framework.Failf("You have to specify KubeConfigPath. (Use KUBECONFIG environment variable.)")
	}

	switch p := framework.DeleteTestingNSPolicyType(os.Getenv("DELETE_NS_POLICY")); p {
	case framework.DeleteTestingNSPolicyAlways,
		framework.DeleteTestingNSPolicyOnSuccess,
		framework.DeleteTestingNSPolicyNever:
		TestContext.DeleteTestingNSPolicy = framework.DeleteTestingNSPolicyType(p)
	case "":
		TestContext.DeleteTestingNSPolicy = framework.DeleteTestingNSPolicyAlways
	default:
		framework.Failf("Invalid DeleteTestingNSPolicy: %q", TestContext.DeleteTestingNSPolicy)
	}

	fixedNamespace := os.Getenv("FIXED_NAMESPACE")
	if fixedNamespace == "" {
		TestContext.DeleteTestingNSPolicy = framework.DeleteTestingNSPolicyNever
		TestContext.CreateTestingNS = func(f *framework.Framework, name string, labels map[string]string) (*corev1.Namespace, error) {
			return &corev1.Namespace{
				ObjectMeta: metav1.ObjectMeta{
					Name: fixedNamespace,
				},
			}, nil
		}
	} else {
		TestContext.CreateTestingNS = framework.CreateTestingProjectAndChangeUser
	}
}

func ExecuteTest(t *testing.T, suite string) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, suite)
}
