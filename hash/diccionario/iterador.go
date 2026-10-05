package diccionario

// ----- Iterador interno -----

func (h *hashCerrado[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for _, c := range h.tabla {
		if c.estado == _OCUPADA {
			if !visitar(c.clave, c.dato) {
				return
			}
		}
	}
}

// ----- Iterador externo -----

type iterHashCerrado[K comparable, V any] struct {
	hash *hashCerrado[K, V]
	pos  int
}

func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterHashCerrado[K, V]{hash: h, pos: 0}
	iter.avanzarHastaOcupada()
	return iter
}

func (iter *iterHashCerrado[K, V]) avanzarHastaOcupada() {
	for iter.pos < iter.hash.tam && iter.hash.tabla[iter.pos].estado != _OCUPADA {
		iter.pos++
	}
}

func (iter *iterHashCerrado[K, V]) HayAlgoMas() bool {
	return iter.pos < iter.hash.tam
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	actual := iter.hash.tabla[iter.pos]
	return actual.clave, actual.dato
}

func (iter *iterHashCerrado[K, V]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iter.pos++
	iter.avanzarHastaOcupada()
}
