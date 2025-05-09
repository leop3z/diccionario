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
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := diccionario.CrearABB[string, string](compararLongitud)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}
func TestObtenerDato(t *testing.T) {
	abb := diccionario.CrearABB[string, int](compararLongitud)
	abb.Guardar("Don Quijote", 34)
	abb.Guardar("Rosinante", 11)
	require.Equal(t, 34, abb.Obtener("Don Quijote"), "Busca el valor de la clave dada")
}
