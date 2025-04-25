package main

import (
	"fmt"
	"time"
)

func main() {
	var hour int = time.Now().Hour()
	var Kentekens [5]string
	Kentekens[0] = "AB-12-34"
	Kentekens[1] = "12-34-CD"
	Kentekens[2] = "XY-56-78"
	Kentekens[3] = "90-AB-12"
	Kentekens[4] = "CD-34-56"

	var Naam string

	//als de tijd lager is dan 7 dan foutmelding geven "We zijn helaas gesloten!"
	if hour < 7 {
		fmt.Println("Sorry,", Naam, "de parkeerplaats is ’s nachts gesloten!")
	}
	//als de tijd tussen 7 en 12 in zit dan melding geven "Goedemorgen"
	if hour >= 7 && hour <= 12 {
		fmt.Println("Goedemorgen", Naam, "! Het  is nu", hour, "uur")
	}
	//als de tijd tussen 12 en 18 zit dan melding geven "Goedemiddag"
	if hour >= 12 && hour <= 18 {
		fmt.Println("Goedemiddag", Naam, "! Het is nu", hour, "uur")
	}
	//als de tijd tussen 18 en 23 zit dan melding geven "goedeavond"
	if hour >= 18 && hour <= 23 {
		fmt.Println("Goedeavond!", Naam, "! Het is nu", hour, "uur")
	}

	fmt.Print("Voer alstublieft uw naam in: ")
	fmt.Scan(&Naam)

	fmt.Print("Voer alstublieft uw kenteken in (xx-xx-xx): ")
	var KentekenInput string

	for {
		fmt.Scan(&KentekenInput)

		if len(KentekenInput) != 8 {
			fmt.Println("Ongeldige invoer! Het kenteken moet precies 8 tekens zijn.")
			continue
		}
		break
	}

	for i := 0; i < len(Kentekens); i++ {
		if KentekenInput == Kentekens[i] {
			fmt.Println("Welkom", Naam, "! Uw kenteken is", KentekenInput)
			break
		} else {
			fmt.Println("Sorry,", Naam, "Uw kenteken staat niet geregistreerd in ons lijst!")
			break
		}
	}
}

//Kentekens
//AB-12-34
//12-34-CD
//XY-56-78
//90-AB-12
//CD-34-56
