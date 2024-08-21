package language

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestNewNamespace(t *testing.T) {
	//ns := NewNamespace(PHP{})
	//if ns.name != "test" {
	//	t.Errorf("NewNamespace() failed: %s", ns.name)
	// }
	// f, _ := os.Open("testdata/php_namespace.proto")
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/service_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
		Package: proto.String("test"),
	}
	req := &pluginpb.CodeGeneratorRequest{
		ProtoFile: []*descriptorpb.FileDescriptorProto{fdp},
	}
	sdp := &descriptorpb.ServiceDescriptorProto{
		Name: proto.String("Test"),
		Method: []*descriptorpb.MethodDescriptorProto{
			{
				Name: proto.String("Test"),
			},
		},
	}
	ns := NewNamespace(PHP{}, req, fdp, sdp)
	if ns.Namespace != "Test\\CustomNamespace" {
		t.Errorf("NewNamespace() failed: %s", ns.Namespace)
	}
}
