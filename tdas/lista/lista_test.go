package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestInsertarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("primero añadido, este debe quedar como último")
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, "primero añadido, este debe quedar como último", lista.VerPrimero())
	require.Equal(t, "primero añadido, este debe quedar como último", lista.VerUltimo())

	lista.InsertarPrimero("segundo añadido, debe quedar en el medio")
	lista.InsertarPrimero("tercero añadido, debe quedar de primero")
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, "tercero añadido, debe quedar de primero", lista.VerPrimero())
	require.Equal(t, "primero añadido, este debe quedar como último", lista.VerUltimo())
}

func TestInsertarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()

	lista.InsertarUltimo(true)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, true, lista.VerPrimero())
	require.Equal(t, true, lista.VerUltimo())

	lista.InsertarUltimo(false)
	lista.InsertarUltimo(false)
	require.Equal(t, 3, lista.Largo())

	require.Equal(t, true, lista.VerPrimero())
	require.Equal(t, false, lista.VerUltimo())
}

func TestInsertar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	lista.InsertarUltimo(1)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)
	lista.InsertarPrimero(4)

	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())
	require.False(t, lista.EstaVacia())
}

func TestInsertarBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()

	lista.InsertarPrimero("primero")
	lista.InsertarPrimero("segundo")
	lista.InsertarPrimero("tercero (queda primero)")

	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, "tercero (queda primero)", lista.BorrarPrimero())

	require.False(t, lista.EstaVacia())
	require.Equal(t, 2, lista.Largo())
	require.Equal(t, "segundo", lista.BorrarPrimero())

	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, "primero", lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestListaOrdenada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()

	lista.InsertarUltimo(true)
	lista.InsertarUltimo(false)
	lista.InsertarUltimo(true)
	lista.InsertarPrimero(true)
	lista.InsertarPrimero(false)

	require.Equal(t, false, lista.BorrarPrimero())
	require.Equal(t, true, lista.BorrarPrimero())
	require.Equal(t, true, lista.BorrarPrimero())
	require.Equal(t, false, lista.BorrarPrimero())
	require.Equal(t, true, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestVolumenLista(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	volumenTest := 10000
	for i := 0; i <= volumenTest; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 1+i, lista.Largo())
	}
	for k := volumenTest; k >= 0; k-- {
		require.Equal(t, k, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())

	for j := 0; j <= volumenTest; j++ {
		lista.InsertarUltimo(j)
		require.Equal(t, j, lista.VerUltimo())
		require.Equal(t, j+1, lista.Largo())
	}
}

func TestComportamientoVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	lista.InsertarUltimo(true)
	lista.InsertarUltimo(false)
	lista.InsertarUltimo(true)

	require.Equal(t, true, lista.BorrarPrimero())
	require.Equal(t, false, lista.BorrarPrimero())
	require.Equal(t, true, lista.BorrarPrimero())

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestIterInternoCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("5")
	lista.InsertarPrimero("4")
	lista.InsertarPrimero("3")
	lista.InsertarPrimero("2")
	lista.InsertarPrimero("1")

	cont := 0
	lista.Iterar(func(c string) bool {
		if cont >= 4 {
			return false
		}
		cont += 1
		return true
	})
	require.Equal(t, 4, cont)
}

func TestIterInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	cont, suma := 0, 0
	lista.Iterar(func(valor int) bool {
		cont += 1
		suma += valor
		return true
	})
	require.Equal(t, 5, cont)
	require.Equal(t, 15, suma)
}

func TestIterInternoSinElems(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	lista.InsertarPrimero(true)
	lista.InsertarPrimero(true)
	lista.InsertarPrimero(true)
	lista.InsertarPrimero(true)
	lista.InsertarPrimero(false)

	cont, valorOpuesto := 0, true
	lista.Iterar(func(valor bool) bool {
		cont += 1
		return valor == valorOpuesto
	})

	require.Equal(t, 1, cont)
}

func TestIterExterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)

	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() {
		dato := iter.VerActual()
		encontrado := false
		for i := 0; i <= 6; i++ {
			if dato == i+1 {
				encontrado = true
				break
			}
		}
		require.Equal(t, true, encontrado)
	}
}

func TestIterExtVacio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Avanzar() })
}

func TestIterExtInsertarPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iter := lista.Iterador()

	require.False(t, iter.HayAlgoMas())

	iter.Insertar("H")
	require.Equal(t, "H", iter.VerActual())
	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestIterExtInsertarFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(1)
	iter := lista.Iterador()

	require.Equal(t, 1, iter.VerActual())
	require.True(t, iter.HayAlgoMas())
	iter.Avanzar()

	require.Equal(t, 2, iter.VerActual())
	require.True(t, iter.HayAlgoMas())
	iter.Avanzar()

	require.False(t, iter.HayAlgoMas())

	iter.Insertar(3)
	require.True(t, iter.HayAlgoMas())
	require.Equal(t, 3, iter.VerActual())

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestIterExtInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("la")
	lista.InsertarPrimero("Ho")
	iter := lista.Iterador()

	require.Equal(t, "Ho", iter.VerActual())
	require.True(t, iter.HayAlgoMas())
	iter.Avanzar()

	iter.Insertar("Esto va en el medio")

	require.Equal(t, "Esto va en el medio", iter.VerActual())
	require.True(t, iter.HayAlgoMas())
	iter.Avanzar()

	require.Equal(t, "la", iter.VerActual())
	require.True(t, iter.HayAlgoMas())

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestIterExtBorrarPrimerElem(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("2")
	lista.InsertarPrimero("1")

	iter := lista.Iterador()
	require.Equal(t, "1", iter.Borrar())
	require.Equal(t, "2", iter.VerActual())

	iter.Avanzar()
	require.False(t, iter.HayAlgoMas())
}

func TestIterExtBorrarUltimoElem(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	lista.InsertarUltimo(false)
	lista.InsertarUltimo(true)
	lista.InsertarUltimo(true)

	iter := lista.Iterador()

	require.Equal(t, false, iter.VerActual())
	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())

	require.Equal(t, true, iter.VerActual())
	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())

	require.Equal(t, true, iter.VerActual())
	require.Equal(t, true, iter.Borrar())
}

func TestIterExtBorrarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	lista.InsertarUltimo(false)
	lista.InsertarUltimo(true)
	lista.InsertarUltimo(false)

	iter := lista.Iterador()

	require.Equal(t, false, iter.VerActual())
	iter.Avanzar()
	require.True(t, iter.HayAlgoMas())

	require.Equal(t, true, iter.VerActual())
	require.Equal(t, true, iter.Borrar())

	require.Equal(t, false, iter.VerActual())
	require.True(t, iter.HayAlgoMas())
}
