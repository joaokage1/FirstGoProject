package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {

	t.Parallel() // executa o teste de forma paralela
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciiona em arquitetura amd64")
	}

}
