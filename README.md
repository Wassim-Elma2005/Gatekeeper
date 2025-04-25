README – Gatekeeper Applicatie

Overzicht

De Gatekeeper-applicatie is ontwikkeld in Golang en is bedoeld om kentekens te controleren bij de ingang van een terrein.
De applicatie vraagt de gebruiker om een kenteken in te voeren en controleert of dit kenteken voorkomt in een vooraf ingestelde lijst.
Op basis van het kenteken geeft de applicatie toegang of weigert deze toegang.

Daarnaast geeft de applicatie op basis van het moment van de dag (ochtend, middag, avond) een begroeting.
De applicatie verwerkt gebruikersinput en bevat eenvoudige foutafhandeling.

Functionaliteit
	•	De gebruiker voert een kenteken in.
	•	De applicatie vergelijkt dit kenteken met een lijst van toegestane kentekens.
	•	Bij een geldig kenteken verschijnt een begroeting (gebaseerd op tijdstip) en wordt toegang verleend.
	•	Bij een ongeldig kenteken wordt toegang geweigerd.
	•	Alle belangrijke outputs worden weergegeven in de console.

Kenmerken
	•	Gebruikersinput: Scannen van het kenteken via de terminal.
	•	Controle op toegestane kentekens: Vergelijking via array/lijsten in Go.
	•	Tijdsafhankelijke begroeting: “Goedemorgen”, “Goedemiddag” of “Goedenavond” afhankelijk van de lokale tijd.
	•	Eenvoudige foutafhandeling: Ongeldige invoer wordt netjes afgehandeld.

Uitvoering
	•	Code ontwikkeld in Go (Golang).
	•	Uitvoerbaar via:

go run main.go


	•	De lijst met geldige kentekens kan in de code aangepast worden door de array allowedPlates te bewerken.

Gebruik van ChatGPT

Tijdens de ontwikkeling van de Gatekeeper-applicatie heb ik ChatGPT ingezet om hulp te krijgen bij de structuur van de applicatie, zoals bij het verwerken van gebruikersinput, foutafhandeling en het correct toepassen van tijdsfuncties.
