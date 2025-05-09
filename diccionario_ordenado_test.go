package diccionario_test

import (
	"tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararLongitud(clave1 string, clave2 string) int {
	return len(clave1) - len(clave2)
}

func TestDiccionarioVacio(t *testing.T) {
	abb := diccionario.CrearABB[string, string](compararLongitud)
	require.Equal(t, abb.Cantidad(), 0, "Como recien se creo el arbol no tiene elementos")
	abb.Guardar("Mordecai", "Azul")
	require.Equal(t, abb.Cantidad(), 1, "El arbol ya no está vacío")
}
