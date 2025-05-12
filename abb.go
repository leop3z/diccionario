package diccionario

import (
	"fmt"
	TDAPila "tdas/pila"
)

const (
	ES_MENOR    = -1
	ES_MAYOR    = 1
	SON_IGUALES = 0
)

type nodoAbb[K comparable, V any] struct {
	clave    K
	dato     V
	hijo_izq *nodoAbb[K, V]
	hijo_der *nodoAbb[K, V]
}

type funcCmp[K comparable] func(a, b K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cmp      funcCmp[K]
	cantidad int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cmp: funcion_cmp, cantidad: 0}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodo := abb.buscar(&abb.raiz, clave)
	if *nodo == nil {
		*nodo = &nodoAbb[K, V]{clave: clave, dato: dato}
		abb.cantidad++
	} else {
		(*nodo).dato = dato
	}
}

// Retorna la direccion de memoria donde se almacena el puntero que apunta al nodo buscado (o apuntaria)
func (abb *abb[K, V]) buscar(nodo **nodoAbb[K, V], clave K) **nodoAbb[K, V] {
	if *nodo == nil {
		return nodo
	}
	switch abb.cmp(clave, (*nodo).clave) {
	case ES_MENOR:
		return abb.buscar(&(*nodo).hijo_izq, clave)
	case ES_MAYOR:
		return abb.buscar(&(*nodo).hijo_der, clave)
	default:
		return nodo
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return *abb.buscar(&abb.raiz, clave) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.buscar(&abb.raiz, clave)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	fmt.Println(*nodo)
	return (*nodo).dato
}

// Dado un nodo de un arbol, retorna el minimo de ese sub-arbol
func (abb *abb[K, V]) buscarMin(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	for (*nodo).hijo_izq != nil {
		nodo = &(*nodo).hijo_izq
	}
	return nodo
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo := abb.buscar(&abb.raiz, clave)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	borrado := (*nodo).dato
	abb.cantidad--
	switch {
	case (*nodo).hijo_izq == nil && (*nodo).hijo_der == nil:
		*nodo = nil

	case (*nodo).hijo_izq != nil && (*nodo).hijo_der == nil:
		*nodo = (*nodo).hijo_izq

	case (*nodo).hijo_izq == nil && (*nodo).hijo_der != nil:
		*nodo = (*nodo).hijo_der

	case (*nodo).hijo_izq != nil && (*nodo).hijo_der != nil:
		minDer := abb.buscarMin(&(*nodo).hijo_der)
		(*nodo).clave = (*minDer).clave
		(*nodo).dato = (*minDer).dato
		*minDer = (*minDer).hijo_der
	}
	return borrado
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.raiz.iterar(visitar)
}

func (nodo *nodoAbb[K, V]) iterar(visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	nodo.hijo_izq.iterar(visitar)
	if !visitar(nodo.clave, nodo.dato) {
		return
	}
	nodo.hijo_der.iterar(visitar)
}

type iteradorDiccionario[K comparable, V any] struct {
	pila TDAPila.Pila[*nodoAbb[K, V]]
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	nodo := abb.raiz
	for nodo != nil {
		pila.Apilar(nodo)
		nodo = nodo.hijo_izq
	}
	return &iteradorDiccionario[K, V]{pila: pila}
}

func (iter *iteradorDiccionario[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	if nodo.hijo_der != nil {
		actual := nodo.hijo_der
		for actual != nil {
			iter.pila.Apilar(actual)
			actual = actual.hijo_izq
		}
	}
}

func (iter *iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	hay_menores := desde == nil || abb.cmp(*desde, nodo.clave) == ES_MENOR
	hay_mayores := hasta == nil || abb.cmp(*hasta, nodo.clave) == ES_MAYOR
	es_limite := (desde == nil || abb.cmp(*desde, nodo.clave) <= 0) && (hasta == nil || abb.cmp(*hasta, nodo.clave) >= 0)
	if hay_menores {
		abb.iterarRango(nodo.hijo_izq, desde, hasta, visitar)
	}
	if (hay_menores && hay_mayores) || es_limite {
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}
	if hay_mayores {
		abb.iterarRango(nodo.hijo_der, desde, hasta, visitar)
	}
}

type iteradorDiccionarioRango[K comparable, V any] struct {
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	cmp   funcCmp[K]
	desde *K
	hasta *K
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if abb.raiz != nil {
		nodo := abb.raiz
		for nodo != nil {
			if desde == nil || abb.cmp(*desde, nodo.clave) <= 0 {
				pila.Apilar(nodo)
				nodo = nodo.hijo_izq
			} else {
				nodo = nodo.hijo_der
			}
		}
	}
	return &iteradorDiccionarioRango[K, V]{pila: pila, cmp: abb.cmp, desde: desde, hasta: hasta}
}

func (iter *iteradorDiccionarioRango[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorDiccionarioRango[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	nodo = nodo.hijo_der
	for nodo != nil {
		menOIg_hasta := iter.hasta == nil || iter.cmp(*iter.hasta, nodo.clave) >= 0
		mayOIg_desde := iter.desde == nil || iter.cmp(*iter.desde, nodo.clave) <= 0
		if mayOIg_desde && menOIg_hasta {
			iter.pila.Apilar(nodo)
		}
		nodo = nodo.hijo_izq
	}
}

func (iter *iteradorDiccionarioRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}
