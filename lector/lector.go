package lector

import (
	Sort "main/ordenamiento_busqueda"
	"main/votos"

	"bufio"
	"os"
	"strconv"
	"strings"
)

func LeerPadrones(ruta string) []votos.Votante {
	lista_padrones := []int{}
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil
	}
	defer archivo.Close()
	lector := bufio.NewScanner(archivo)
	// Leemos el archivo de los padrones y cada numero lo metemos en la lista enlazada
	for lector.Scan() {
		DNI, errorconv := strconv.Atoi(lector.Text())
		if errorconv != nil {
			return nil
		}
		lista_padrones = append(lista_padrones, DNI)
	}
	err = lector.Err()
	if err != nil {
		return nil
	}

	// Ordenamos mediante RadixSort

	padrones_ordenados := Sort.RadixDni(lista_padrones)
	resultado := make([]votos.Votante, len(padrones_ordenados))
	for index, DNI := range padrones_ordenados {
		resultado[index] = votos.CrearVotante(DNI)
	}
	return resultado
}

func LeerLista(ruta string) []votos.Partido {
	lista_partidos := []votos.Partido{}
	lista_partidos = append(lista_partidos, votos.CrearVotosEnBlanco())
	listacsv, err := os.Open(ruta)
	if err != nil {
		return nil
	}
	defer listacsv.Close()
	lector := bufio.NewScanner(listacsv)
	for lector.Scan() {
		linea := strings.Split(lector.Text(), ",")
		partido := votos.CrearPartido(linea[0], [3]string{linea[1], linea[2], linea[3]})
		lista_partidos = append(lista_partidos, partido)
	}
	err = lector.Err()
	if err != nil {
		return nil
	}
	return lista_partidos
}
