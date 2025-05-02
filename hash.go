package diccionario

import (
	"fmt"
	"hash/fnv"
)

const (
	VACIO              = 0
	OCUPADO            = 1
	BORRADO            = 2
	FACTOR_CRE         = 0.70
	FACTOR_DEC         = 0.20
	TAMANO_INICIAL     = 13
	FACTOR_REDIMENSION = 2
)

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	borrados int
	tam      int
}

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashCerrado[K, V]{}

	hash.tabla = make([]celdaHash[K, V], TAMANO_INICIAL)
	hash.cantidad = 0
	hash.borrados = 0
	hash.tam = TAMANO_INICIAL
	return hash
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashing(b []byte) uint64 {
	hasher := fnv.New64()
	hasher.Write(b)
	return hasher.Sum64()
}

func buscar[K comparable, V any](clave any, tabla []celdaHash[K, V]) uint64 {
	pos := hashing(convertirABytes(clave)) % uint64(len(tabla))
	for !(tabla[pos].estado == VACIO || tabla[pos].clave == clave) {
		if int(pos) == len(tabla)-1 {
			pos = 0
		} else {
			pos++
		}
	}
	return pos
}

func (hash *hashCerrado[K, V]) redimensionar(nuevo_tam int) {
	hash.tam = nuevo_tam
	nueva := make([]celdaHash[K, V], nuevo_tam)
	for i := range hash.tabla {
		if hash.tabla[i].estado == OCUPADO {
			pos := buscar(hash.tabla[i].clave, nueva)
			nueva[pos] = hash.tabla[i]
		}
	}
	hash.borrados = 0
	hash.tabla = nueva
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	carga := float64(hash.cantidad+hash.borrados) / float64(hash.tam)
	if FACTOR_CRE < carga {
		hash.redimensionar(hash.tam * FACTOR_REDIMENSION)
	}
	pos := buscar(clave, hash.tabla)
	hash.tabla[pos].dato = dato
	if !hash.Pertenece(clave) {
		hash.tabla[pos].clave = clave
		hash.tabla[pos].estado = OCUPADO
		hash.cantidad++
	}
}

func (hash hashCerrado[K, V]) Pertenece(clave K) bool {
	celda := hash.tabla[buscar(clave, hash.tabla)]
	return celda.clave == clave && celda.estado == OCUPADO
}

func (hash hashCerrado[K, V]) Obtener(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[buscar(clave, hash.tabla)].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	hash.cantidad--
	hash.borrados++
	pos := buscar(clave, hash.tabla)
	dato := hash.tabla[pos].dato
	hash.tabla[pos].estado = BORRADO
	carga := float64(hash.cantidad+hash.borrados) / float64(hash.tam)
	if FACTOR_DEC > carga && hash.tam > TAMANO_INICIAL {
		hash.redimensionar(hash.tam / FACTOR_REDIMENSION)
	}
	return dato
}

func (hash hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(K, V) bool) {
	pos := 0
	for pos < len(hash.tabla) {
		if hash.tabla[pos].estado == OCUPADO && !visitar(hash.tabla[pos].clave, hash.tabla[pos].dato) {
			break
		}
		pos++
	}
}

type iteradorHash[K comparable, V any] struct {
	pos_actual int
	tabla      []celdaHash[K, V]
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	pos := 0
	for pos < len(hash.tabla) && hash.tabla[pos].estado != OCUPADO {
		pos++
	}
	return &iteradorHash[K, V]{pos_actual: pos, tabla: hash.tabla}
}

func (iter iteradorHash[K, V]) HaySiguiente() bool {
	return iter.pos_actual < len(iter.tabla)
}

func (iter iteradorHash[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.tabla[iter.pos_actual].clave, iter.tabla[iter.pos_actual].dato
}

func (iter *iteradorHash[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.pos_actual++
	for iter.pos_actual < len(iter.tabla) && iter.tabla[iter.pos_actual].estado != OCUPADO {
		iter.pos_actual++
	}
}
