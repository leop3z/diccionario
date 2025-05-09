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
	// Mi arbol esta vacio
	if abb == nil {
		return &nodoArbol[K, V]{clave: clave_nueva, dato: dato}
	}
	comparacion := abb.cmp(clave_nueva, nodo.clave)
	if comparacion == PRIMERA_ES_MENOR {
		nodo.hijo_izq = abb.guardarRecursivo(nodo.hijo_izq, clave_nueva, dato)
	} else if comparacion == SEGUNDA_ES_MENOR {
		nodo.hijo_der = abb.guardarRecursivo(nodo.hijo_der, clave_nueva, dato)
	} else {
		nodo.dato = dato
	}
	abb.cantidad++
	return nodo
}

func (abb *abbDiccionario[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.guardarRecursivo(abb.raiz, clave, dato)
}

func (abb *abbDiccionario[K, V]) busquedaRecursiva(nodo *nodoArbol[K, V], clave K) bool {
	return clave == nodo.clave
}

func (abb *abbDiccionario[K, V]) Pertenece(clave K) bool {
	return abb.raiz.clave == clave
}

func (hash *abbDiccionario[K, V]) Obtener(clave K) V {
	var dato V
	return dato
}

func (hash *abbDiccionario[K, V]) Borrar(clave K) V {
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
