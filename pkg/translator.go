package pkg

var symptoms = map[string]string{
	"feverHigher39":          "Fiebre mayor a 39",
	"painEyes":               "Dolor detras de los ojos",
	"musclePain":             "Dolor muscular",
	"persistentVomiting":     "Vomitos persistentes",
	"bloodyVomit":            "Vomitos con sangre",
	"skinSpots":              "Manchas en la piel",
	"lymphNodes":             "Ganglios linfaticos",
	"abdominalPain":          "Dolor abdominal",
	"acceleratedBreathing":   "Respiracion acelerada",
	"bleeding":               "Sangrado en las encías o la nariz",
	"fatigue":                "Cansancio",
	"agitation":              "Agitacion",
	"bloodInStool":           "Sangre en heces",
	"headache":               "Dolor de cabeza",
	"diarrhea":               "Diarrea",
	"weightLoss":             "Perdida de peso",
	"cough":                  "Tos",
	"feverMinus39":           "Fiebre menor a 39",
	"nasalCongestion":        "Congestion nasal",
	"whiteStools":            "Heces blancas o amarrillo claro",
	"brownUrine":             "Orina color “coca cola”",
	"jaundice":               "Color amarillo en mucosas, piel y/o ojos",
	"rightHypochondriumPain": "Dolor en el hipocondrio derecho",
	"abdominalDistention":    "Distencion abdominal",
	"nausea":                 "Nauseas",
}

func SymptomsToSpanish(symptomsMarked []string) string {
	symptomsInSpanish := ""
	for i := 0; i < len(symptomsMarked); i++ {
		for symptoms, val := range symptoms {
			if symptomsMarked[i] == symptoms {
				if symptomsInSpanish == "" {
					symptomsInSpanish = val
				} else {
					symptomsInSpanish = symptomsInSpanish + ", " + val
				}
			}
		}
	}
	return symptomsInSpanish
}