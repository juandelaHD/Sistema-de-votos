package votos

import (
	"fmt"
)

type partidoImplementacion struct {
	nombre string

	candidatos [CANT_VOTACION]string

	cantidadVotos [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	cantidadVotos [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	partido := new(partidoImplementacion)
	partido.nombre = nombre
	partido.candidatos = candidatos
	partido.cantidadVotos = [CANT_VOTACION]int{0, 0, 0}
	return partido
}

func CrearVotosEnBlanco() Partido {
	partidoBlanco := new(partidoEnBlanco)
	partidoBlanco.cantidadVotos = [CANT_VOTACION]int{0, 0, 0}
	return partidoBlanco
}

func (partido *partidoImplementacion) VotadoPara(tipo int) {
	partido.cantidadVotos[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var palabra string
	if partido.cantidadVotos[tipo] == 1 {
		palabra = "voto"
	} else {
		palabra = "votos"
	}
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], partido.cantidadVotos[tipo], palabra)
}

func (blanco *partidoEnBlanco) VotadoPara(tipo int) {
	blanco.cantidadVotos[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	var palabra string
	if blanco.cantidadVotos[tipo] == 1 {
		palabra = "voto"
	} else {
		palabra = "votos"
	}
	return fmt.Sprintf("Votos en Blanco: %d %s", blanco.cantidadVotos[tipo], palabra)
}
