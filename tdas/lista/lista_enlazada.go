package lista

type listaEnlazada[T any] struct{
	primero *nodoLista[T]
	ultimo *nodoLista[T]
	longitud int
}

type nodoLista[T any] struct{
	dato T
	prox *nodoLista[T]
}

type iteradorListaEnlazada[T any] struct{
	lista *listaEnlazada[T]
	actual *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
    return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool{
	return lista.longitud==0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T){
	nodo:=&nodoLista[T]{dato:elemento, prox:nil}

	if lista.EstaVacia(){
		lista.primero = nodo
		lista.ultimo = nodo
	}else{
		nodo.prox = lista.primero
		lista.primero = nodo
	}
	lista.longitud+=1
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T){
	nodo:=&nodoLista[T]{dato:elemento, prox:nil}

	if lista.EstaVacia(){
		lista.primero = nodo
		lista.ultimo = nodo
	}else{
		lista.ultimo.prox = nodo
		lista.ultimo = nodo
	}
	lista.longitud += 1
}

func (lista *listaEnlazada[T]) BorrarPrimero() (T){
	if lista.EstaVacia(){
		panic("La lista esta vacia")
	}
	elemento:=lista.primero.dato
	lista.primero = lista.primero.prox
	lista.longitud--
	if lista.EstaVacia(){
		lista.ultimo=nil
	}
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() (T){
	if lista.EstaVacia(){
		panic("La lista esta vacia")
	}
	elemento:=lista.primero.dato
	return elemento
}

func (lista *listaEnlazada[T]) VerUltimo() (T){
	if lista.EstaVacia(){
		panic("La lista esta vacia")
	}
	elemento:=lista.ultimo.dato
	return elemento
}

func (lista *listaEnlazada[T]) Largo()(int){
	return lista.longitud
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual:=lista.primero
	for actual != nil{
		sigue:=visitar(actual.dato)
		if sigue == false{
			break
		}
		actual = actual.prox
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T]{
	return &iteradorListaEnlazada[T]{lista: lista, actual: lista.primero, anterior: nil}
}


func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	panic("TODO: implementar VerActual")
}

func (iterador *iteradorListaEnlazada[T]) HayAlgoMas() bool {
	panic("TODO: implementar HayAlgoMas")
}

func (iterador *iteradorListaEnlazada[T]) Avanzar() {
	panic("TODO: implementar Avanzar")
}

func (iterador *iteradorListaEnlazada[T]) Insertar(elemento T) {
	panic("TODO: implementar Insertar")
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	panic("TODO: implementar Borrar")
}
