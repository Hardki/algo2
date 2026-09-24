package lista

type listaEnlazada[T any] struct {
	primero  *nodoLista[T]
	ultimo   *nodoLista[T]
	longitud int
}

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.longitud == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nodo := &nodoLista[T]{dato: elemento, prox: nil}

	if lista.EstaVacia() {
		lista.primero = nodo
		lista.ultimo = nodo
	} else {
		nodo.prox = lista.primero
		lista.primero = nodo
	}
	lista.longitud += 1
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nodo := &nodoLista[T]{dato: elemento, prox: nil}

	if lista.EstaVacia() {
		lista.primero = nodo
		lista.ultimo = nodo
	} else {
		lista.ultimo.prox = nodo
		lista.ultimo = nodo
	}
	lista.longitud += 1
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	elemento := lista.primero.dato
	lista.primero = lista.primero.prox
	lista.longitud--
	if lista.EstaVacia() {
		lista.ultimo = nil
	}
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	elemento := lista.primero.dato
	return elemento
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	elemento := lista.ultimo.dato
	return elemento
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.longitud
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		sigue := visitar(actual.dato)
		if sigue == false {
			break
		}
		actual = actual.prox
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{lista: lista, actual: lista.primero, anterior: nil}
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) HayAlgoMas() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Avanzar() {
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.prox
}

func (iterador *iteradorListaEnlazada[T]) Insertar(elemento T) {
	nodo := &nodoLista[T]{dato: elemento, prox: iterador.actual}
	// Caso lista vacia
	if iterador.lista.longitud == 0 {
		iterador.lista.primero, iterador.lista.ultimo = nodo, nodo
		// Iterador posicionado en el primer elemento
	} else if iterador.anterior == nil {
		iterador.lista.primero = nodo
		// Iterador posicionado en el último elemento
	} else if !iterador.HayAlgoMas() {
		iterador.anterior.prox, iterador.lista.ultimo = nodo, nodo
	} else {
		iterador.anterior.prox = nodo
	}
	iterador.actual = nodo
	iterador.lista.longitud += 1
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	// Caso se finalizó de iterar o no tiene elementos
	if !iterador.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	elemBorrado := iterador.VerActual()
	// Caso solo había 1 elemento
	if iterador.lista.longitud == 1 {
		iterador.lista.primero, iterador.lista.ultimo = nil, nil
		// Caso el siguiente elemento está vacío
	} else if iterador.actual.prox == nil {
		iterador.anterior.prox, iterador.lista.ultimo = nil, iterador.anterior
		// Caso iterador en el primer elemento
	} else if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.prox
	} else {
		iterador.anterior.prox = iterador.actual.prox
	}
	iterador.actual = iterador.actual.prox
	iterador.lista.longitud -= 1
	return elemBorrado
}
