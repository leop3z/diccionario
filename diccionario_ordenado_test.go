package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	"tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

const VOLUMEN = 10000

func compararInt(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func funcComparacion(clave1 string, clave2 string) int {
	return strings.Compare(clave1, clave2)
}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := diccionario.CrearABB[string, string](funcComparacion)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.Panics(t, func() { dic.Obtener("A") }, "Si el diccionario esta vacio entonces no puede obtener nada")
	require.Panics(t, func() { dic.Borrar("A") }, "No se puede borrar de un diccionario vacio")
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := diccionario.CrearABB[string, int](funcComparacion)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := diccionario.CrearABB[string, string](funcComparacion)
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := diccionario.CrearABB[string, string](funcComparacion)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := diccionario.CrearABB[string, string](funcComparacion)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}
func TestDiccionarioBorrarABB(t *testing.T) {
	claves := []int{5, 3, 1, 4, 2, 8, 6, 9}
	valores := []string{"e", "c", "a", "d", "b", "h", "f", "i"}
	dic := diccionario.CrearABB[int, string](compararInt)
	for i := range len(claves) {
		dic.Guardar(claves[i], valores[i])
	}
	require.Equal(t, 8, dic.Cantidad())
	dic.Borrar(9)
	// Despues de borrar un nodo sin hijos, efectivamente tiene que no estar en el Arbol
	require.False(t, dic.Pertenece(9))
	require.Equal(t, 7, dic.Cantidad())
	// Despues de borrar un nodo con un hijo, efectivamente debe no pertenecer al Arbol, pero su hijo si.
	dic.Borrar(8)
	require.False(t, dic.Pertenece(8))
	require.Equal(t, 6, dic.Cantidad())
	require.True(t, dic.Pertenece(6))
	require.Equal(t, "f", dic.Obtener(6))
	// Despues de borrar un nodo con 2 hijos, efectivamente ya no pertenece al Arbol, pero si sus hijos.
	dic.Borrar(3)
	require.Equal(t, 5, dic.Cantidad())
	require.False(t, dic.Pertenece(3))
	require.True(t, dic.Pertenece(1))
	require.Equal(t, "a", dic.Obtener(1))
	require.True(t, dic.Pertenece(4))
	require.Equal(t, "d", dic.Obtener(4))
	// Borrar la raiz no rompe el arbol
	dic.Borrar(5)
	require.Equal(t, 4, dic.Cantidad())
	require.False(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(1))
	require.Equal(t, "a", dic.Obtener(1))
	require.True(t, dic.Pertenece(4))
	require.Equal(t, "d", dic.Obtener(4))
	require.True(t, dic.Pertenece(6))
	require.Equal(t, "f", dic.Obtener(6))
	require.True(t, dic.Pertenece(2))
	require.Equal(t, "b", dic.Obtener(2))
	// Borrar una clave inexistente no altera el arbol
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(10) })
	require.False(t, dic.Pertenece(10))
	require.Equal(t, 4, dic.Cantidad())
	require.False(t, dic.Pertenece(5))
	require.True(t, dic.Pertenece(1))
	require.Equal(t, "a", dic.Obtener(1))
	require.True(t, dic.Pertenece(4))
	require.Equal(t, "d", dic.Obtener(4))
	require.True(t, dic.Pertenece(6))
	require.Equal(t, "f", dic.Obtener(6))
	require.True(t, dic.Pertenece(2))
	require.Equal(t, "b", dic.Obtener(2))
	// Borrar los nodos efectivamente borrar el arbol
	dic.Borrar(4)
	dic.Borrar(1)
	dic.Borrar(6)
	dic.Borrar(2)
	require.False(t, dic.Pertenece(4))
	require.False(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(6))
	require.False(t, dic.Pertenece(2))
	require.Equal(t, 0, dic.Cantidad())
}
func TestReutlizacionDeBorrados(t *testing.T) {
	dic := diccionario.CrearABB[string, string](funcComparacion)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := diccionario.CrearABB[string, string](funcComparacion)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestCadenaLargaParticular(t *testing.T) {
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := diccionario.CrearABB[string, string](funcComparacion)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := diccionario.CrearABB[string, *int](funcComparacion)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoClaves(t *testing.T) {
	// Valida que todas las claves sean recorridas (en orden) con el iterador interno
	claves := []int{5, 3, 1, 4, 2, 8, 6, 9}
	valores := []string{"e", "c", "a", "d", "b", "h", "f", "i"}
	dic := diccionario.CrearABB[int, string](compararInt)
	for i := range len(claves) {
		dic.Guardar(claves[i], valores[i])
	}
	require.Equal(t, 8, dic.Cantidad())

	cs := []int{0, 0, 0, 0, 0, 0, 0, 0}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave int, dato string) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})
	require.EqualValues(t, 8, cantidad)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9}, cs)
	// Valida que los datos sean recorridas correctamente (en orden) con el iterador interno
	cv := []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	dic.Iterar(func(clave int, dato string) bool {
		cv[cantidad] = dato
		*cantPtr = *cantPtr + 1
		return true
	})
	require.EqualValues(t, 8, cantidad)
	require.Equal(t, []string{"a", "b", "c", "d", "e", "f", "h", "i"}, cv)
	// Valida que la funcion se aplique correctamente a medida que se itera y efectivamente detiene la iteracion.
	cs = []int{0, 0, 0, 0, 0, 0, 0, 0}
	cv = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	dic.Iterar(func(clave int, dato string) bool {
		if clave%2 != 0 {
			cs[cantidad] = clave
			cv[cantidad] = dato
			*cantPtr = *cantPtr + 1
		}
		if clave == 8 {
			return false
		}
		return true
	})
	require.EqualValues(t, 3, cantidad)
	require.Equal(t, []int{1, 3, 5, 0, 0, 0, 0, 0}, cs)
	require.Equal(t, []string{"a", "c", "e", "", "", "", "", ""}, cv)
	// Valida que los datos sean recorridas correctamente (en orden) con el iterador interno, sin recorrer datos borrados
	dic.Borrar(2)
	dic.Borrar(6)
	dic.Borrar(5)
	cb := ""
	cd := []int{0, 0, 0, 0, 0}
	cantidad = 0
	dic.Iterar(func(clave int, dato string) bool {
		cd[cantidad] = clave
		cb += dato
		*cantPtr = *cantPtr + 1
		return true
	})
	require.EqualValues(t, 5, cantidad)
	require.Equal(t, "acdhi", cb)
	require.Equal(t, []int{1, 3, 4, 8, 9}, cd)
}

func TestDiccionarioIterarRango(t *testing.T) {
	// Iterar con Rango en un diccionario vacio no hace nada.
	dic := diccionario.CrearABB[int, string](compararInt)
	vacioc := []int{0, 0, 0}
	vaciov := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	desde := 2
	hasta := 5
	dic.IterarRango(&desde, &hasta, func(clave int, dato string) bool {
		vacioc[cantidad] = clave
		vaciov[cantidad] = dato
		*cantPtr++
		return true
	})
	require.Equal(t, 0, cantidad)
	require.Equal(t, []int{0, 0, 0}, vacioc)
	require.Equal(t, []string{"", "", ""}, vaciov)

	claves := []int{5, 3, 1, 4, 2, 8, 6, 9}
	valores := []string{"e", "c", "a", "d", "b", "h", "f", "i"}

	for i := range len(claves) {
		dic.Guardar(claves[i], valores[i])
	}
	require.Equal(t, 8, dic.Cantidad())
	// Si desde = nil Iterar con Rango itera desde el inicio
	sc := []int{0, 0, 0, 0, 0, 0, 0, 0}
	sv := []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	desde = 2
	hasta = 6
	dic.IterarRango(nil, &hasta, func(clave int, dato string) bool {
		sc[cantidad] = clave
		sv[cantidad] = dato
		*cantPtr++
		return true
	})
	require.EqualValues(t, 6, cantidad)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 0, 0}, sc)
	require.Equal(t, []string{"a", "b", "c", "d", "e", "f", "", ""}, sv)

	// Si hasta = nil Iterar con Rango itera hasta el final
	sc = []int{0, 0, 0, 0, 0, 0, 0, 0}
	sv = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	desde = 2
	hasta = 6
	dic.IterarRango(&desde, nil, func(clave int, dato string) bool {
		sc[cantidad] = clave
		sv[cantidad] = dato
		*cantPtr++
		return true
	})
	require.EqualValues(t, 7, cantidad)
	require.Equal(t, []int{2, 3, 4, 5, 6, 8, 9, 0}, sc)
	require.Equal(t, []string{"b", "c", "d", "e", "f", "h", "i", ""}, sv)

	// Si hasta y desde son nil, Iterar con rango se comporta igual que Iterar.
	sc = []int{0, 0, 0, 0, 0, 0, 0, 0}
	sv = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	dic.IterarRango(nil, nil, func(clave int, dato string) bool {
		sc[cantidad] = clave
		sv[cantidad] = dato
		*cantPtr++
		return true
	})
	require.EqualValues(t, 8, cantidad)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9}, sc)
	require.Equal(t, []string{"a", "b", "c", "d", "e", "f", "h", "i"}, sv)

	// Verifica que el iterador con rango itera correctamente (una vez y en orden) al diccionario.
	sc = []int{0, 0, 0, 0, 0, 0, 0, 0}
	sv = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	desde = 2
	hasta = 6
	dic.IterarRango(&desde, &hasta, func(clave int, dato string) bool {
		sc[cantidad] = clave
		sv[cantidad] = dato
		*cantPtr++
		return true
	})
	require.EqualValues(t, 5, cantidad)
	require.Equal(t, []int{2, 3, 4, 5, 6, 0, 0, 0}, sc)
	require.Equal(t, []string{"b", "c", "d", "e", "f", "", "", ""}, sv)
	// Verifica que el iterador con rango itera correctamente (una vez y en orden), efectivamente validando la funcion parametro.
	sc = []int{0, 0, 0, 0, 0, 0, 0, 0}
	sv = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	desde = 1
	hasta = 8

	dic.IterarRango(&desde, &hasta, func(clave int, dato string) bool {
		if clave > 6 {
			return false
		}
		if clave%2 == 0 {
			sc[cantidad] = clave
			sv[cantidad] = dato
			*cantPtr++
		}
		return true
	})
	require.EqualValues(t, 3, cantidad)
	require.Equal(t, []int{2, 4, 6, 0, 0, 0, 0, 0}, sc)
	require.Equal(t, []string{"b", "d", "f", "", "", "", "", ""}, sv)
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := diccionario.CrearABB[string, int](funcComparacion)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterar(t *testing.T) {
	// Guardamos valores en un Diccionario, e iteramos validando que las claves y datos sean recorridos en orden. Además los valores de VerActual y Siguiente van siendo correctos entre sí.
	claves := []int{5, 3, 1, 4, 2, 8, 6, 9}
	valores := []string{"e", "c", "a", "d", "b", "h", "f", "i"}
	dic := diccionario.CrearABB[int, string](compararInt)
	for i := range len(claves) {
		dic.Guardar(claves[i], valores[i])
	}
	require.Equal(t, 8, dic.Cantidad())

	cs := []int{0, 0, 0, 0, 0, 0, 0, 0}
	vs := []string{"", "", "", "", "", "", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iter.VerActual()
		*cantPtr += 1
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.EqualValues(t, 8, cantidad)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9}, cs)
	require.Equal(t, []string{"a", "b", "c", "d", "e", "f", "h", "i"}, vs)
	// El iterador recorre en orden los elementos del Diccionario correctamente (en orden), luego de haber sido borrado algunos de sus valores
	dic.Borrar(2)
	dic.Borrar(6)
	dic.Borrar(5)
	cs = []int{0, 0, 0, 0, 0}
	vs = []string{"", "", "", "", ""}
	cantidad = 0
	iter2 := dic.Iterador()
	for iter2.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iter2.VerActual()
		*cantPtr += 1
		require.True(t, iter2.HaySiguiente())
		iter2.HaySiguiente()
		iter2.Siguiente()
	}
	require.False(t, iter2.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter2.Siguiente() })
	require.EqualValues(t, 5, cantidad)
	require.Equal(t, []int{1, 3, 4, 8, 9}, cs)
	require.Equal(t, []string{"a", "c", "d", "h", "i"}, vs)
}

func TestDiccionarioItExternoRango(t *testing.T) {
	// Un iterador externo con rango en un diccionario vacio no hace nada, al no tener elementos para iterar.
	dic := diccionario.CrearABB[int, string](compararInt)
	vacioc := []int{0, 0, 0}
	vaciov := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad
	desde := 2
	hasta := 5
	iterVacio := dic.IteradorRango(&desde, &hasta)
	require.False(t, iterVacio.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterVacio.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterVacio.Siguiente() })
	for iterVacio.HaySiguiente() {
		vacioc[cantidad], vaciov[cantidad] = iterVacio.VerActual()
		*cantPtr++
		iterVacio.Siguiente()
	}
	require.Equal(t, 0, cantidad)
	require.Equal(t, []int{0, 0, 0}, vacioc)
	require.Equal(t, []string{"", "", ""}, vaciov)

	claves := []int{5, 3, 1, 4, 2, 8, 6, 9}
	valores := []string{"e", "c", "a", "d", "b", "h", "f", "i"}

	for i := range len(claves) {
		dic.Guardar(claves[i], valores[i])
	}
	require.Equal(t, 8, dic.Cantidad())

	// Si desde es nil, el Iterador Externo con rango itera desde el inicio.
	cs := []int{0, 0, 0, 0, 0, 0, 0, 0}
	vs := []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	hasta = 5
	iterSinInicio := dic.IteradorRango(nil, &hasta)
	for iterSinInicio.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iterSinInicio.VerActual()
		*cantPtr += 1
		require.True(t, iterSinInicio.HaySiguiente())
		iterSinInicio.Siguiente()
	}
	require.False(t, iterSinInicio.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterSinInicio.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterSinInicio.Siguiente() })
	require.EqualValues(t, 5, cantidad)
	require.Equal(t, []int{1, 2, 3, 4, 5, 0, 0, 0}, cs)
	require.Equal(t, []string{"a", "b", "c", "d", "e", "", "", ""}, vs)

	// Si hasta es nil, el Iterador Externo con rango itera hasta el final.
	cs = []int{0, 0, 0, 0, 0, 0, 0, 0}
	vs = []string{"", "", "", "", "", "", "", ""}
	desde = 2
	cantidad = 0
	iterSinFinal := dic.IteradorRango(&desde, nil)
	for iterSinFinal.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iterSinFinal.VerActual()
		*cantPtr += 1
		require.True(t, iterSinFinal.HaySiguiente())
		iterSinFinal.Siguiente()
	}
	require.False(t, iterSinFinal.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterSinFinal.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterSinFinal.Siguiente() })
	require.EqualValues(t, 7, cantidad)
	require.Equal(t, []int{2, 3, 4, 5, 6, 8, 9, 0}, cs)
	require.Equal(t, []string{"b", "c", "d", "e", "f", "h", "i", ""}, vs)

	// Si desde y hasta son nil, el Iterador Externo con rango se comporta igual que el  Iterador externo sin rango.
	cs = []int{0, 0, 0, 0, 0, 0, 0, 0}
	vs = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	iterSinRango := dic.IteradorRango(nil, nil)
	for iterSinRango.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iterSinRango.VerActual()
		*cantPtr += 1
		require.True(t, iterSinRango.HaySiguiente())
		iterSinRango.Siguiente()
	}
	require.False(t, iterSinRango.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterSinRango.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterSinRango.Siguiente() })
	require.EqualValues(t, 8, cantidad)
	require.Equal(t, []int{1, 2, 3, 4, 5, 6, 8, 9}, cs)
	require.Equal(t, []string{"a", "b", "c", "d", "e", "f", "h", "i"}, vs)
	// Guardamos valores en un Diccionario, e iteramos en un rango validando que las claves y datos sean recorridos en orden. Además los valores de VerActual y Siguiente van siendo correctos entre sí.

	cs = []int{0, 0, 0, 0, 0, 0, 0, 0}
	vs = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	desde = 2
	hasta = 6
	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iter.VerActual()
		*cantPtr += 1
		require.True(t, iter.HaySiguiente())
		iter.HaySiguiente()
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.EqualValues(t, 5, cantidad)
	require.Equal(t, []int{2, 3, 4, 5, 6, 0, 0, 0}, cs)
	require.Equal(t, []string{"b", "c", "d", "e", "f", "", "", ""}, vs)

	// Iterador Externo con rango, itera correctamente despues de haber borrado valores.
	dic.Borrar(2)
	dic.Borrar(4)
	dic.Borrar(6)
	cs = []int{0, 0, 0, 0, 0, 0, 0, 0}
	vs = []string{"", "", "", "", "", "", "", ""}
	cantidad = 0
	desde = 1
	hasta = 7
	iterBorrado := dic.IteradorRango(&desde, &hasta)
	for iterBorrado.HaySiguiente() {
		cs[cantidad], vs[cantidad] = iterBorrado.VerActual()
		*cantPtr += 1
		require.True(t, iterBorrado.HaySiguiente())
		iterBorrado.HaySiguiente()
		iterBorrado.Siguiente()
	}
	require.False(t, iterBorrado.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterBorrado.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterBorrado.Siguiente() })
	require.EqualValues(t, 3, cantidad)
	require.Equal(t, []int{1, 3, 5, 0, 0, 0, 0, 0}, cs)
	require.Equal(t, []string{"a", "c", "e", "", "", "", "", ""}, vs)
}

func TestDiccionarioVolumen(t *testing.T) {
	arreglo := make([]int, VOLUMEN)
	for i := range VOLUMEN {
		arreglo[i] = i
	}
	rand.Shuffle(len(arreglo), func(a, b int) {
		arreglo[a], arreglo[b] = arreglo[b], arreglo[a]
	})
	// Verifica que los elementos esten ahi luego de guardados
	dic := diccionario.CrearABB[string, int](strings.Compare)
	for i := range VOLUMEN {
		dic.Guardar(fmt.Sprintf("%d", arreglo[i]), arreglo[i])
		require.True(t, dic.Pertenece(fmt.Sprintf("%d", arreglo[i])))
	}
	// Verifica que la cantidad sea la correcta
	require.Equal(t, VOLUMEN, dic.Cantidad())
	// Se verifica que efectivamente esten todos los elementos
	for i := range VOLUMEN {
		require.Equal(t, arreglo[i], dic.Obtener(fmt.Sprintf("%d", arreglo[i])))
	}
	// Se borran los elementos y se verifica que efectivamente ahora el diccionario esta vacio.
	for i := range VOLUMEN {
		dic.Borrar(fmt.Sprintf("%d", arreglo[i]))
		require.False(t, dic.Pertenece(fmt.Sprintf("%d", arreglo[i])))
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestIteradorInternoVolumen(t *testing.T) {
	aleatorio := make([]int, VOLUMEN)
	for i := range VOLUMEN {
		aleatorio[i] = i
	}
	rand.Shuffle(len(aleatorio), func(a, b int) {
		aleatorio[a], aleatorio[b] = aleatorio[b], aleatorio[a]
	})
	ordenado := make([]int, VOLUMEN)
	dic := diccionario.CrearABB[int, int](compararInt)
	for i := range VOLUMEN {
		dic.Guardar(aleatorio[i], aleatorio[i])
	}
	iteraciones := 0
	// Verifica que efecticamente el iterador itera las claves en orden y una unica vez.
	dic.Iterar(func(clave int, dato int) bool {
		ordenado[iteraciones] = dato
		iteraciones++
		return true
	})
	require.True(t, iteraciones == VOLUMEN)
	for i := 1; i < len(ordenado); i++ {
		require.True(t, ordenado[i] > ordenado[i-1])
	}
}

func TestIteradorExternoVolumen(t *testing.T) {
	aleatorio := make([]int, VOLUMEN)
	for i := range VOLUMEN {
		aleatorio[i] = i
	}
	rand.Shuffle(len(aleatorio), func(a, b int) {
		aleatorio[a], aleatorio[b] = aleatorio[b], aleatorio[a]
	})
	ordenado := make([]int, VOLUMEN)
	for i := range VOLUMEN {
		ordenado[i] = i
	}
	dic := diccionario.CrearABB[int, int](compararInt)
	for i := range VOLUMEN {
		dic.Guardar(aleatorio[i], aleatorio[i])
	}
	iter := dic.Iterador()
	iteraciones := 0
	// Verifica que el Iterador externo pase solo una vez por los elementos y que esto sea en el orden correcto.
	for iter.HaySiguiente() {
		actual, _ := iter.VerActual()
		fmt.Println(actual, ordenado[iteraciones])
		require.Equal(t, actual, ordenado[iteraciones])
		iteraciones++
		iter.Siguiente()
	}
}

func TestIteradorInternoRangoVolumen(t *testing.T) {
	aleatorio := make([]int, VOLUMEN)
	for i := range VOLUMEN {
		aleatorio[i] = i
	}
	rand.Shuffle(len(aleatorio), func(a, b int) {
		aleatorio[a], aleatorio[b] = aleatorio[b], aleatorio[a]
	})
	hasta := rand.Intn(VOLUMEN)
	desde := rand.Intn(hasta)
	total_acotado := (hasta - desde) + 1
	ordenado := make([]int, total_acotado)
	dic := diccionario.CrearABB[int, int](compararInt)
	for i := range VOLUMEN {
		dic.Guardar(aleatorio[i], aleatorio[i])
	}
	iteraciones := 0
	dic.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		ordenado[iteraciones] = dato
		iteraciones++
		return true
	})
	// Verifica que se itero una sola vez dentro del rango y en orden.
	require.True(t, iteraciones == total_acotado)
	for i := range len(ordenado) {
		require.Equal(t, desde, ordenado[i])
		desde++
	}
}

func TestIteradorExternoRangoVolumen(t *testing.T) {
	aleatorio := make([]int, VOLUMEN)
	for i := range VOLUMEN {
		aleatorio[i] = i
	}
	rand.Shuffle(len(aleatorio), func(a, b int) {
		aleatorio[a], aleatorio[b] = aleatorio[b], aleatorio[a]
	})
	hasta := rand.Intn(VOLUMEN)
	desde := rand.Intn(hasta)
	total_acotado := (hasta - desde) + 1
	ordenado := make([]int, total_acotado)
	for i := range total_acotado {
		ordenado[i] = desde + i
	}
	dic := diccionario.CrearABB[int, int](compararInt)
	for i := range VOLUMEN {
		dic.Guardar(aleatorio[i], aleatorio[i])
	}
	hasta--
	iter := dic.IteradorRango(&desde, &hasta)
	iteraciones := 0
	// Verifica que el Iterador externo pase solo una vez por los elementos y que esto sea en el orden correcto.
	for iter.HaySiguiente() {
		actual, _ := iter.VerActual()
		fmt.Println(actual, ordenado[iteraciones])
		require.Equal(t, actual, ordenado[iteraciones])
		iteraciones++
		iter.Siguiente()
	}
}
