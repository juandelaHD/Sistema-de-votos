package main

import (
	"bufio"
	"fmt"
	TDACola "main/cola"
	"main/errores"
	lectura "main/lector"
	"main/votos"
	"os"
	"strings"
)

var args = os.Args[1:]

func main() {
	if len(args) != 2 {
		err := errores.ErrorParametros{}
		fmt.Println(err.Error())
		return
	}
	archivo_lista := args[0]
	archivo_padrones := args[1]

	padrones := lectura.LeerPadrones(archivo_padrones)
	lista_candidatos := lectura.LeerLista(archivo_lista)
	if padrones == nil || lista_candidatos == nil {
		err := errores.ErrorLeerArchivo{}
		fmt.Println(err.Error())
		return
	}

	fila := TDACola.CrearColaEnlazada[votos.Votante]()
	cant_impugnados := 0

	S := bufio.NewScanner(os.Stdin)
	for S.Scan() {
		lectura.LecturaDeComandos(strings.Fields(S.Text()), fila, padrones, lista_candidatos, &cant_impugnados)
	}
	// Aca termina el stdin (con ctrl + d)
	if !fila.EstaVacia() {
		err := errores.ErrorCiudadanosSinVotar{}
		fmt.Println(err.Error())
	}
	fmt.Println("Presidente:")
	for _, partido := range lista_candidatos {
		fmt.Println(partido.ObtenerResultado(votos.PRESIDENTE))
	}
	fmt.Println("\nGobernador:")
	for _, partido := range lista_candidatos {
		fmt.Println(partido.ObtenerResultado(votos.GOBERNADOR))
	}
	fmt.Println("\nIntendente:")
	for _, partido := range lista_candidatos {
		fmt.Println(partido.ObtenerResultado(votos.INTENDENTE))
	}
	var palabra string
	switch cant_impugnados {
	case 1:
		palabra = "voto"
	default:
		palabra = "votos"
	}
	fmt.Printf("\nVotos Impugnados: %d %s\n", cant_impugnados, palabra)
}
