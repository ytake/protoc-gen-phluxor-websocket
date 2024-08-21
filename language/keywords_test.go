package language

import (
	"testing"
)

func TestPHP_Camelize(t *testing.T) {
	p := PHP{}
	if p.Camelize("test_case") != "TestCase" {
		t.Errorf("Camelize() failed: %s", p.Camelize("test_case"))
	}
}

func TestPHP_Identifier(t *testing.T) {
	p := PHP{}
	if p.Identifier("test_case", "suffix") != "TestCaseSuffix" {
		t.Errorf("Identifier() failed: %s", p.Identifier("test_case", "suffix"))
	}
}

func TestPHP_Namespace(t *testing.T) {
	p := PHP{}
	pkg := "test\\Case"
	if p.Namespace(&pkg, "/") != "Test\\Case" {
		t.Errorf("Namespace() failed: %s", p.Namespace(&pkg, "/"))
	}
}
