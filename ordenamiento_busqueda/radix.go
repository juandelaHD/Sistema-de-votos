package ordenamiento_busqueda

const _CANT_ARRAY_MAXIMO int = 100

func RadixDni(arr []int) []int {
	ordenadoporA := countingSort(arr, func(num int) int {
		return (num % 100)
	})
	ordenadoporB := countingSort(ordenadoporA, func(num int) int {
		return ((num % 10000) / 100)
	})
	ordenadoporC := countingSort(ordenadoporB, func(num int) int {
		return ((num % 1000000) / 10000)
	})
	ordenadoporD := countingSort(ordenadoporC, func(num int) int {
		return (num / 1000000)
	})
	return ordenadoporD
}

func countingSort(arr []int, criterio func(int) int) []int {

	contador := make([]int, _CANT_ARRAY_MAXIMO)
	for _, elem := range arr {
		contador[criterio(elem)]++
	}

	suma_Acumulada := make([]int, _CANT_ARRAY_MAXIMO)
	suma := 0
	for index, elem := range contador {
		suma_Acumulada[index] += suma
		suma += elem
	}

	ordenado := make([]int, len(arr))
	for _, elem := range arr {
		index_suma := criterio(elem)
		index_ordenado := suma_Acumulada[index_suma]
		ordenado[index_ordenado] = elem
		suma_Acumulada[index_suma]++
	}

	return ordenado
}
