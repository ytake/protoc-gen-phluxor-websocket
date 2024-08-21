package language

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestPHP_DetectNamespace(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/php_namespace.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
	}
	p := PHP{}
	if p.DetectNamespace(fdp) != "Test/CustomNamespace" {
		t.Error("DetectNamespace() failed")
	}
}

func TestClient_Filename(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/client_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
	}
	c := &Client{p: PHP{}}
	if c.Filename(fdp, proto.String("Test")) != "Test/CustomNamespace/TestClient.php" {
		t.Errorf("Client.Filename() failed: %s", c.Filename(fdp, proto.String("Test")))
	}
}

func TestInterfaceName_Filename(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/interface_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
	}
	i := &InterfaceName{p: PHP{}}
	if i.Filename(fdp, proto.String("Test")) != "Test/CustomNamespace/TestInterface.php" {
		t.Errorf("InterfaceName.Filename() failed: %s", i.Filename(fdp, proto.String("Test")))
	}
}

func TestServiceName_Filename(t *testing.T) {
	fdp := &descriptorpb.FileDescriptorProto{
		Name: proto.String("testdata/service_filename.proto"),
		Options: &descriptorpb.FileOptions{
			PhpNamespace: proto.String("Test\\CustomNamespace"),
		},
	}
	s := &ServiceName{p: PHP{}}
	if s.Filename(fdp, proto.String("Test")) != "Test/CustomNamespace/TestService.php" {
		t.Errorf("ServiceName.Filename() failed: %s", s.Filename(fdp, proto.String("Test")))
	}
}
