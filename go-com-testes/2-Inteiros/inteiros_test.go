package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionar(t *testing.T) {
	verificarRespostaCorreta := func(t *testing.T, resposta, esperado int) {
		t.Helper()
		if resposta != esperado {
			t.Errorf("resposta %d, esperado %d", resposta, esperado)
		}
	}

	t.Run("testa o resultado da soma de dois inteiros", func(t *testing.T) {
		resposta := Adicionar(2, 2)
		esperado := 4

		verificarRespostaCorreta(t, resposta, esperado)
	})
}

// IMPORTANTE: o código desta função irá aparecer como exemplo na documentação.
// E o output será exibido como resultado na documentação. Para criar esse tipo de função ela irá
// precisar das palavras reservadas "Example" no começo e da palavra "Output" no comentada
// IMPORTANTE: este comentpario também aparecerá na documentação.
func ExampleAdicionar() {
	soma := Adicionar(2, 5)
	fmt.Println(soma)
	//Output: 7
}
