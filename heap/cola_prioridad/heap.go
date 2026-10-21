package cola_prioridad

type colaPrioridad[T any] struct {
	datos []T
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaPrioridad[T]{
		datos: []T{},
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arreglo))
	copy(datos, arreglo)
	heapify(datos, funcion_cmp)
	return &colaPrioridad[T]{
		datos: datos,
		cmp:   funcion_cmp,
	}
}

func (c *colaPrioridad[T]) EstaVacia() bool {
	return len(c.datos) == 0
}

func (c *colaPrioridad[T]) Cantidad() int {
	return len(c.datos)
}

func (c *colaPrioridad[T]) VerMax() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.datos[0]
}

func (c *colaPrioridad[T]) Encolar(elemento T) {
	c.datos = append(c.datos, elemento)
	upHeap(c.datos, len(c.datos)-1, c.cmp)
}

func (c *colaPrioridad[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	maximo := c.datos[0]
	ultimo := len(c.datos) - 1
	c.datos[0] = c.datos[ultimo]
	c.datos = c.datos[:ultimo]
	downHeap(c.datos, 0, len(c.datos), c.cmp)
	return maximo
}

func upHeap[T any](datos []T, i int, cmp func(T, T) int) {
	for i > 0 {
		padre := (i - 1) / 2
		if cmp(datos[i], datos[padre]) <= 0 {
			break
		}
		datos[i], datos[padre] = datos[padre], datos[i]
		i = padre
	}
}

func downHeap[T any](datos []T, i, tam int, cmp func(T, T) int) {
	for {
		izquierdo := 2*i + 1
		derecho := 2*i + 2
		mayor := i
		if izquierdo < tam && cmp(datos[izquierdo], datos[mayor]) > 0 {
			mayor = izquierdo
		}
		if derecho < tam && cmp(datos[derecho], datos[mayor]) > 0 {
			mayor = derecho
		}
		if mayor == i {
			break
		}
		datos[i], datos[mayor] = datos[mayor], datos[i]
		i = mayor
	}
}

func heapify[T any](datos []T, cmp func(T, T) int) {
	for i := len(datos)/2 - 1; i >= 0; i-- {
		downHeap(datos, i, len(datos), cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, funcion_cmp)
	for fin := len(elementos) - 1; fin > 0; fin-- {
		elementos[0], elementos[fin] = elementos[fin], elementos[0]
		downHeap(elementos, 0, fin, funcion_cmp)
	}
}