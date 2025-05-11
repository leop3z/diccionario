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
	if comparacion < 0 {
		nodo.hijo_izq = abb.guardarRecursivo(nodo.hijo_izq, clave_nueva, dato)
	} else if comparacion > 0 {
		nodo.hijo_der = abb.guardarRecursivo(nodo.hijo_der, clave_nueva, dato)
	} else {
		nodo.dato = dato
	}
	return nodo
}

func (abb *abbDiccionario[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.guardarRecursivo(abb.raiz, clave, dato)
}

func (abb *abbDiccionario[K, V]) busquedaRecursiva(nodo *nodoArbol[K, V], clave K) *nodoArbol[K, V] {
	if nodo == nil {
		return nil
	}
	comparacion := abb.cmp(clave, nodo.clave)
	if comparacion < 0 {
		return abb.busquedaRecursiva(nodo.hijo_izq, clave)
	} else if comparacion > 0 {
		return abb.busquedaRecursiva(nodo.hijo_der, clave)
	}
	return nodo
}

func (abb *abbDiccionario[K, V]) Pertenece(clave K) bool {
	return abb.busquedaRecursiva(abb.raiz, clave) != nil
}

func (abb *abbDiccionario[K, V]) Obtener(clave K) V {
	nodo := abb.busquedaRecursiva(abb.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

// Funcion para buscar el minimo del sub-arbol derecho
func (abb *abbDiccionario[K, V]) buscarMinimo(nodo *nodoArbol[K, V]) *nodoArbol[K, V] {
	for nodo != nil && nodo.hijo_izq != nil {
		nodo = nodo.hijo_izq
	}
	return nodo
}

// Recorre recursivamente el arbol xq debe ir actualizando los punteros
func (abb *abbDiccionario[K, V]) borrarNodo(nodo *nodoArbol[K, V], clave K) (*nodoArbol[K, V], V) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	comparacion := abb.cmp(clave, nodo.clave)
	if comparacion < 0 {
		var valor V
		nodo.hijo_izq, valor = abb.borrarNodo(nodo.hijo_izq, clave)
		return nodo, valor
	} else if comparacion > 0 {
		var valor V
		nodo.hijo_der, valor = abb.borrarNodo(nodo.hijo_der, clave)
		return nodo, valor
	}

	dato_original := nodo.dato
	abb.cantidad--
	// No tiene hijos
	if nodo.hijo_izq == nil && nodo.hijo_der == nil {
		return nil, dato_original
	}
	// tiene un hijo
	if nodo.hijo_izq == nil {
		return nodo.hijo_der, dato_original
	}
	if nodo.hijo_der == nil {
		return nodo.hijo_izq, dato_original
	}
	// tiene 2 hijos
	sucesor := abb.buscarMinimo(nodo.hijo_der)
	nodo.clave = sucesor.clave
	nodo.dato = sucesor.dato
	nodo.hijo_der, _ = abb.borrarNodo(nodo.hijo_der, sucesor.clave)
	return nodo, dato_original
}

func (abb *abbDiccionario[K, V]) Borrar(clave K) V {
	nueva_raiz, dato := abb.borrarNodo(abb.raiz, clave)
	abb.raiz = nueva_raiz
	return dato
}

func (abb *abbDiccionario[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abbDiccionario[K, V]) Iterar(visitar func(clave1 K, dato V) bool) {
}

type iteradorDiccionario[K comparable, V any] struct {
}

func (abb *abbDiccionario[K, V]) Iterador() IterDiccionario[K, V] {
	return iteradorDiccionario[K, V]{}
}

func (abb iteradorDiccionario[K, V]) HaySiguiente() bool {
	return false
}

func (abb iteradorDiccionario[K, V]) Siguiente() {

}

func (abb iteradorDiccionario[K, V]) VerActual() (K, V) {
	var (
		clave K
		dato  V
	)
	return clave, dato
}

func (abb *abbDiccionario[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	panic("IterarRango no implementado todavía")
}

func (abb *abbDiccionario[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	panic("IteradorRango no implementado todavía")
}

/*

func (abb *abbDiccionario[K,V])borrarNodo(nodo *nodoArbol[K,V], clave K) V{
	nodo_borrar := abb.buquedaRecursiva(nodo, clave)

	// demas operaciones para manejar los casos de la cantidad de hijos
}

*/
