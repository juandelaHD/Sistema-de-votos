package pila

const (
	_CAPACIDAD_INICIAL = 10
	_MULTIPLICADOR     = 2
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAPACIDAD_INICIAL)
	return pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == cap(pila.datos) {
		pila.redimensionar(cap(pila.datos) * _MULTIPLICADOR)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	valor := pila.VerTope()
	pila.cantidad--
	if pila.cantidad*_MULTIPLICADOR*_MULTIPLICADOR <= cap(pila.datos) && cap(pila.datos) > _CAPACIDAD_INICIAL {
		pila.redimensionar(cap(pila.datos) / _MULTIPLICADOR)
	}
	return valor
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	arregloRedimensionado := make([]T, nuevaCapacidad)
	copy(arregloRedimensionado, pila.datos)
	pila.datos = arregloRedimensionado
}
