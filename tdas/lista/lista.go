package lista

type Lista[T any] interface {
	// Devuelve true en caso de que la lista no tenga elementos, false en caso de que si tenga aun elementos
	EstaVacia() bool
	// Inserta un elemento de tipo T en la primer posicion de la lista
	InsertarPrimero(T)
	// Inserta un elemento de tipo T en la ultima posicion de la lista
	InsertarUltimo(T)
	// Elimina el primer elemento de la lista y lo devuelve luego de eliminarlo
	BorrarPrimero() T
	// Devuelve el primer elemento de tipo T de la lista
	VerPrimero() T
	// Devuelve el ultimo elemento de tipo T de la lista
	VerUltimo() T
	// Devuelve la cantidad de posiciones de la lista
	Largo() int
	// Funcion que aplica la funcion visitar que es utilizada por el iterador interno para recorrer la lista y devolver
	// true en caso de haber recorrido todos los elementos de la lista o false en caso de no lograrlo
	Iterar(visitar func(T) bool)
	// Iterador devuelve un IteradorLista de tipo T posicionado al principio de la lista, que permite modificar la misma
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// Devuelve la posicion actual donde se encuentra posicionado el iterador externo
	// Si el iterador ya termino de iterar, entra en panico con un mensaje
	// "El iterador termino de iterar".
	VerActual() T
	// HayAlgoMas devuelve true si el iterador todavia no llego al final de la
	// lista, false en caso contrario.
	HayAlgoMas() bool
	// Avanzar mueve el iterador un elemento hacia adelante. Si el iterador ya
	// termino de iterar, entra en panico con un mensaje "El iterador termino de iterar".
	Avanzar()
	// Insertar agrega un nuevo elemento en la posicion donde esta parado el
	// iterador. El iterador queda parado sobre el elemento recien insertado.
	Insertar(T)
	// Borrar quita el elemento donde esta parado el iterador y devuelve su
	// valor. El iterador queda parado sobre el elemento que ocupaba la
	// siguiente posicion. Si el iterador ya termino de iterar, entra en
	// panico con un mensaje "El iterador termino de iterar".
	Borrar() T
}
