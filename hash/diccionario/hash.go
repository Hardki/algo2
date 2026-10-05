package diccionario

import "fmt"

// Estados posibles de cada celda de la tabla.
const (
	_VACIA   = 0 // nunca se uso
	_OCUPADA = 1 // tiene un par clave-dato valido
	_BORRADA = 2 // tenia algo, pero se borro (borrado perezoso)
)

const (
	_TAM_INICIAL        = 17  // tamanio de arranque de la tabla
	_FACTOR_CARGA_MAX   = 0.7 // si se llena mas que esto, agrandamos
	_FACTOR_CARGA_MIN   = 0.2 // si queda mas vacio que esto, achicamos
	_FACTOR_REDIMENSION = 2   // cuanto agrandamos/achicamos
)

type celda[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celda[K, V]
	cantidad int // elementos OCUPADA
	borrados int // celdas BORRADA
	tam      int // tamanio de la tabla (len de tabla)
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashing(datos []byte) uint64 {
	const offset = 14695981039346656037
	const prime = 1099511628211
	var hash uint64 = offset
	for _, b := range datos {
		hash ^= uint64(b)
		hash *= prime
	}
	return hash
}

func (h *hashCerrado[K, V]) posicionInicial(clave K) int {
	return int(hashing(convertirABytes(clave)) % uint64(h.tam))
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{
		tabla: make([]celda[K, V], _TAM_INICIAL),
		tam:   _TAM_INICIAL,
	}
}

func (h *hashCerrado[K, V]) buscar(clave K) int {
	pos := h.posicionInicial(clave)
	for h.tabla[pos].estado != _VACIA {
		if h.tabla[pos].estado == _OCUPADA && h.tabla[pos].clave == clave {
			return pos
		}
		pos = (pos + 1) % h.tam
	}
	return -1
}

func (h *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if h.factorDeCarga() > _FACTOR_CARGA_MAX {
		h.redimensionar(h.tam * _FACTOR_REDIMENSION)
	}

	pos := h.buscar(clave)
	if pos != -1 {
		h.tabla[pos].dato = dato
		return
	}
	pos = h.posicionInicial(clave)
	for h.tabla[pos].estado == _OCUPADA {
		pos = (pos + 1) % h.tam
	}
	if h.tabla[pos].estado == _BORRADA {
		h.borrados--
	}
	h.tabla[pos] = celda[K, V]{clave: clave, dato: dato, estado: _OCUPADA}
	h.cantidad++
}

func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	return h.buscar(clave) != -1
}

func (h *hashCerrado[K, V]) Obtener(clave K) V {
	pos := h.buscar(clave)
	if pos == -1 {
		panic("La clave no pertenece al diccionario")
	}
	return h.tabla[pos].dato
}

func (h *hashCerrado[K, V]) Borrar(clave K) V {
	pos := h.buscar(clave)
	if pos == -1 {
		panic("La clave no pertenece al diccionario")
	}
	dato := h.tabla[pos].dato
	h.tabla[pos].estado = _BORRADA
	h.cantidad--
	h.borrados++
	if h.tam > _TAM_INICIAL && h.factorDeCarga() < _FACTOR_CARGA_MIN {
		h.redimensionar(h.tam / _FACTOR_REDIMENSION)
	}
	return dato
	return dato
}

func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashCerrado[K, V]) factorDeCarga() float64 {
	return float64(h.cantidad+h.borrados) / float64(h.tam)
}

func (h *hashCerrado[K, V]) redimensionar(tamNuevo int) {
	tablaVieja := h.tabla
	h.tabla = make([]celda[K, V], tamNuevo)
	h.tam = tamNuevo
	h.cantidad = 0
	h.borrados = 0

	for _, celdaVieja := range tablaVieja {
		if celdaVieja.estado == _OCUPADA {
			h.Guardar(celdaVieja.clave, celdaVieja.dato)
		}
	}
}
