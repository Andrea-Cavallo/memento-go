package main

import (
	"github.com/fatih/structs"
	"log"
	"mememento-go/internal/memento"
	"mememento-go/internal/models"
	"mememento-go/internal/utils"
	"mememento-go/internal/utils/constants"
)

func main() {
	// Creazione di nuovi oggetti person e server
	p := utils.NewPerson(constants.VALIDEMAIL)

	// Creazione di un Originator e impostazione del suo stato
	originator := memento.Originator[models.Person]{State: p}

	// Creazione di un Caretaker
	caretaker := memento.Caretaker[models.Person]{}

	// Il client applica una modifica (simulazione)
	originator.ApplyChange(p)

	// Salvataggio dello stato attuale nel Memento tramite l'Originator
	memento := originator.SaveToMemento()

	// Il Caretaker memorizza il Memento
	caretaker.AddMemento(memento)

	log.Println("=================================================")
	log.Println("Stato originale")

	// Log dello stato iniziale della persona
	log.Println("Person Input:", p)
	log.Println("Is Struct? ", structs.IsStruct(p))
	output := utils.ConvertToMap(p) // Conversione della persona in mappa
	log.Println("Person Output:", output)

	log.Println("=================================================")
	log.Println("Ripristino dello stato")

	// Il client richiede di ripristinare lo stato precedente
	originator.RestoreFromMemento(caretaker.GetMemento(0))

	// Eseguire un'asserzione di tipo per il tipo specifico del tuo progetto
	restoredState := originator.State
	log.Println("Restored Person State:", restoredState)
	log.Println("Is Struct? ", structs.IsStruct(restoredState))
	restoredMap := structs.Map(restoredState) // Conversione dello stato ripristinato in mappa
	log.Println("Restored Person Output:", restoredMap)
}
