package votos

import (
	"main/errores"
	TDAPila "main/pila"
)

type votanteImplementacion struct {
	dni          int
	PilaVotos    TDAPila.Pila[[CANT_VOTACION]int]
	voto         Voto
	finalizoVoto bool
}

func CrearVotante(dni int) Votante {
	votante := new(votanteImplementacion)
	votante.dni = dni
	votante.PilaVotos = TDAPila.CrearPilaDinamica[[CANT_VOTACION]int]()
	votante.PilaVotos.Apilar([CANT_VOTACION]int{LISTA_BLANCO, LISTA_BLANCO, LISTA_BLANCO})
	votante.finalizoVoto = false
	return votante
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante votanteImplementacion) FinalizoVoto() bool {
	return votante.finalizoVoto
}

func (votante *votanteImplementacion) Votar(tipovoto TipoVoto, alternativa int) error {
	if votante.FinalizoVoto() {
		return errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
	}

	votante.voto.sobre = votante.PilaVotos.VerTope()
	votante.voto.sobre[tipovoto] = alternativa
	votante.PilaVotos.Apilar(votante.voto.sobre)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.PilaVotos.EstaVacia() {
		votante.PilaVotos.Apilar([CANT_VOTACION]int{LISTA_BLANCO, LISTA_BLANCO, LISTA_BLANCO})
		err := errores.ErrorNoHayVotosAnteriores{}
		return err
	}
	votante.PilaVotos.Desapilar()
	if votante.PilaVotos.EstaVacia() {
		votante.PilaVotos.Apilar([CANT_VOTACION]int{LISTA_BLANCO, LISTA_BLANCO, LISTA_BLANCO})
		err := errores.ErrorNoHayVotosAnteriores{}
		return err
	}
	votante.voto.sobre = votante.PilaVotos.VerTope()
	return nil
}

func (votante *votanteImplementacion) FinVoto() ([CANT_VOTACION]int, bool, error) {
	if votante.FinalizoVoto() {
		err := errores.ErrorVotanteFraudulento{Dni: votante.LeerDNI()}
		return votante.voto.sobre, votante.voto.impugnado, err
	}

	votante.voto.sobre = votante.PilaVotos.VerTope()
	for !votante.PilaVotos.EstaVacia() {
		for _, elem := range votante.PilaVotos.Desapilar() {
			if elem == LISTA_IMPUGNA {
				votante.voto.impugnado = true
			}
		}
	}
	votante.finalizoVoto = true
	return votante.voto.sobre, votante.voto.impugnado, nil
}
