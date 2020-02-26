package e2eday2

import (
	"testing"

	"github.com/OpenLiberty/open-liberty-operator/pkg/apis"
	openlibertyv1beta1 "github.com/OpenLiberty/open-liberty-operator/pkg/apis/openliberty/v1beta1"
	framework "github.com/operator-framework/operator-sdk/pkg/test"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOpenLibertyApplication(t *testing.T) {
	openLibertyApplicationList := &openlibertyv1beta1.OpenLibertyApplicationList{
		TypeMeta: metav1.TypeMeta{
			Kind: "OpenLibertyApplication",
		},
	}
	openLibertyTraceList := &openlibertyv1beta1.OpenLibertyTraceList{
		TypeMeta: metav1.TypeMeta{
			Kind: "OpenLibertyTrace",
		},
	}

	err := framework.AddToFrameworkScheme(apis.AddToScheme, openLibertyApplicationList)
	if err != nil {
		t.Fatalf("Failed to add CR scheme to framework: %v", err)
	}

	err = framework.AddToFrameworkScheme(apis.AddToScheme, openLibertyTraceList)
	if err != nil {
		t.Fatalf("Failed to add Trace scheme to framework: %v", err)
	}

	t.Run("OpenLibertyDumpsTest", OpenLibertyDumpsTest)
	t.Run("OpenLibertyTraceTest", OpenLibertyTraceTest)
}
