package diccionario

const (
	PRIMERA_ES_MENOR = -1
	SEGUNDA_ES_MENOR = 1
	SON_IGUALES      = 0
)

type nodoArbol[K comparable, V any] struct {
	clave    K
	dato     V
	hijo_izq *nodoArbol[K, V]
	hijo_der *nodoArbol[K, V]
}

type abbDiccionario[K comparable, V any] struct {
	raiz     *nodoArbol[K, V]
	cmp      func(K, K) int
	cantidad int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abbDiccionario[K, V]{raiz: nil, cmp: funcion_cmp, cantidad: 0}
}

func (abb *abbDiccionario[K, V]) guardarRecursivo(nodo *nodoArbol[K, V], clave_nueva K, dato V) *nodoArbol[K, V] {
	if nodo == nil {
		abb.cantidad++
		return &nodoArbol[K, V]{clave: clave_nueva, dato: dato}
	}
	comparacion := abb.cmp(clave_nueva, nodo.clave)
	if comparacion > 0 {
		nodo.hijo_izq = abb.guardarRecursivo(nodo.hijo_izq, clave_nueva, dato)
	} else if comparacion < 0 {
		nodo.hijo_der = abb.guardarRecursivo(nodo.hijo_der, clave_nueva, dato)
	} else {
		nodo.dato = dato
	}
	return nodo
}

func (abb *abbDiccionario[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.guardarRecursivo(abb.raiz, clave, dato)
}

func (abb *abbDiccionario[K, V]) busquedaRecursiva(nodo *nodoArbol[K, V], clave K) bool {
	if nodo == nil {
		return false
	}
	comparacion := abb.cmp(clave, nodo.clave)
	if comparacion < 0 {
		return abb.busquedaRecursiva(nodo.hijo_izq, clave)
	} else if comparacion > 0 {
		return abb.busquedaRecursiva(nodo.hijo_der, clave)
	} else {
		return true
	}
}

func (abb *abbDiccionario[K, V]) obtenerNodo(nodo *nodoArbol[K, V], clave K) *nodoArbol[K, V] {
	comparacion := abb.cmp(clave, abb.raiz.clave)

	if comparacion < 0 {
		return abb.obtenerNodo(nodo.hijo_izq, clave)
	} else if comparacion > 0 {
		return abb.obtenerNodo(nodo.hijo_der, clave)
	}
	return nodo
}

func (abb *abbDiccionario[K, V]) Pertenece(clave K) bool {
	return abb.busquedaRecursiva(abb.raiz, clave)
}

func (abb *abbDiccionario[K, V]) Obtener(clave K) V {
	nodo := abb.obtenerNodo(abb.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abbDiccionario[K, V]) Borrar(clave K) V {
	var dato V
	return dato
}

func (abb *abbDiccionario[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abbDiccionario[K, V]) Iterar(visitar func(clave1 K, dato V) bool) {
}

func (abb *abbDiccionario[K, V]) Iterador() IterDiccionario[K, V] {
	panic("Hola")
}

func (abb *abbDiccionario[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	panic("IterarRango no implementado todavía")
}

func (abb *abbDiccionario[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	panic("IteradorRango no implementado todavía")
}
