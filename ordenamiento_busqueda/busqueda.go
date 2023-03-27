package ordenamiento_busqueda

import "main/votos"

func busqueda_binaria(arr []votos.Votante, inicio int, final int, DNI int) int {
	if inicio >= final {

		return -1
	}
	medio := (inicio + final) / 2
	if arr[medio].LeerDNI() == DNI {
		return medio
	}
	if arr[medio].LeerDNI() < DNI {
		return busqueda_binaria(arr, medio+1, final, DNI)
	} else {
		return busqueda_binaria(arr, inicio, medio, DNI)
	}
}

func Busqueda_binaria(arr []votos.Votante, DNI int) int {
	return busqueda_binaria(arr, 0, len(arr), DNI)
}
