package lector

import (
	"fmt"
	TDACola "main/cola"
	"main/errores"
	"main/ordenamiento_busqueda"
	"main/votos"
	"strconv"
)

func CMDingresar(padrones []votos.Votante, DNI string) votos.Votante {
	dni, errorconv := strconv.Atoi(DNI)
	if errorconv != nil || dni < 0 {
		err := errores.DNIError{}
		fmt.Println(err.Error())
		return nil
	}
	pos_votante_valido := ordenamiento_busqueda.Busqueda_binaria(padrones, dni)
	if pos_votante_valido == -1 {
		err := errores.DNIFueraPadron{}
		fmt.Println(err.Error())
		return nil
	}
	return padrones[pos_votante_valido]
}

func LecturaDeComandos(comando []string, fila TDACola.Cola[votos.Votante], padrones []votos.Votante, lista_candidatos []votos.Partido, cant_impugnados *int) {
	switch comando[0] {

	case "ingresar": // CASO: QUERER METER A UNA PERSONA DEL PADRON A LA COLA PARA VOTAR
		votante_valido := CMDingresar(padrones, comando[1])
		if votante_valido != nil {
			fila.Encolar(votante_valido)
			fmt.Println("OK")
		}

	case "votar": // CASO: HACER A UNA PERSONA DE LA FILA VOTAR A UN CANDIDATO
		if fila.EstaVacia() {
			err := errores.FilaVacia{}
			fmt.Println(err.Error())
			return
		}

		var tipo votos.TipoVoto
		switch comando[1] {
		case "Presidente":
			tipo = votos.PRESIDENTE
		case "Gobernador":
			tipo = votos.GOBERNADOR
		case "Intendente":
			tipo = votos.INTENDENTE
		default:
			err := errores.ErrorTipoVoto{}
			fmt.Println(err.Error())
			return
		}

		alternativa, errorconv := strconv.Atoi(comando[2])
		if alternativa > len(lista_candidatos)-1 || errorconv != nil || alternativa < 0 {
			err := errores.ErrorAlternativaInvalida{}
			fmt.Println(err.Error())
			return
		}

		err := fila.VerPrimero().Votar(tipo, alternativa)
		if err != nil {
			fmt.Println(err.Error())
			fila.Desencolar()
		} else {
			fmt.Println("OK")
		}

	case "deshacer": //CASO: HACER QUE UNA PERSONA DESHAGA SU VOTO
		if fila.EstaVacia() {
			err := errores.FilaVacia{}
			fmt.Println(err.Error())
			return
		}
		if fila.VerPrimero().FinalizoVoto() {
			err := errores.ErrorVotanteFraudulento{Dni: fila.VerPrimero().LeerDNI()}
			fmt.Println(err.Error())
			fila.Desencolar()
			return
		}
		err := fila.VerPrimero().Deshacer()
		if err != nil {
			fmt.Println(err.Error())
			return
		} else {
			fmt.Println("OK")
		}

	case "fin-votar": //CASO: LA PERSONA TERMINO DE VOTAR

		if fila.EstaVacia() {
			err := errores.FilaVacia{}
			fmt.Println(err.Error())
			return
		}

		sobreVoto, impugnado, err := fila.VerPrimero().FinVoto()

		if err != nil {
			fmt.Println(err.Error())
		}
		if impugnado {
			*cant_impugnados++
		} else {
			for index, elemento := range sobreVoto {
				if elemento == votos.LISTA_BLANCO {
					lista_candidatos[0].VotadoPara(index)
				} else {
					lista_candidatos[elemento].VotadoPara(index)
				}
			}
		}
		fila.Desencolar()
		fmt.Println("OK")
		return
	}
}
