package pkg

type Desease struct {
	Name     string
	Symptoms map[string]map[string]string
}

var Dengue = &Desease{
	Name: "Dengue",
	Symptoms: map[string]map[string]string{
		"feverHigher39": {
			"importance": "low",
		},
		"painEyes": {
			"importance": "high",
		},
		"musclePain": {
			"importance": "low",
		},
		"nausea": {
			"importance": "low",
		},
		"persistentVomiting": {
			"importance": "low",
		},
		"bloodyVomit": {
			"importance": "high",
		},
		"skinSpots": {
			"importance": "high",
		},
		"lymphNodes": {
			"importance": "moderate",
		},
		"abdominalPain": {
			"importance": "moderate",
		},
		"acceleratedBreathing": {
			"importance": "moderate",
		},
		"bleeding": {
			"importance": "high",
		},
		"fatigue": {
			"importance": "low",
		},
		"agitation": {
			"importance": "moderate",
		},
		"bloodInStool": {
			"importance": "moderate",
		},
		"headache": {
			"importance": "low",
		},
		"diarrhea": {
			"importance": "low",
		},
		"weightLoss": {
			"importance": "low",
		},
	},
}

var RespiratoryProblems = &Desease{
	Name: "RespiratoryProblems",
	Symptoms: map[string]map[string]string{
		"cough": {
			"importance": "high",
		},
		"feverMinus39": {
			"importance": "low",
		},
		"musclePain": {
			"importance": "low",
		},
		"headache": {
			"importance": "low",
		},
		"nasalCongestion": {
			"importance": "high",
		},
		"fatigue": {
			"importance": "low",
		},
	},
}

var Hepatitis = &Desease{
	Name: "Hepatitis",
	Symptoms: map[string]map[string]string{
		"fatigue": {
			"importance": "moderate",
		},
		"whiteStools": {
			"importance": "high",
		},
		"brownUrine": {
			"importance": "high",
		},
		"jaundice": {
			"importance": "high",
		},
		"rightHypochondriumPain": {
			"importance": "moderate",
		},
		"abdominalDistention": {
			"importance": "moderate",
		},
		"nausea": {
			"importance": "low",
		},
		"persistentVomiting": {
			"importance": "low",
		},
		"feverMinus39": {
			"importance": "low",
		},
		"diarrhea": {
			"importance": "low",
		},
		"headache": {
			"importance": "low",
		},
		"weightLoss": {
			"importance": "low",
		},
	},
}

func (desease *Desease) Detect(patientSymptoms []string, thresholdDisease int16) bool {
	var count int16
	count = 0
	for i := 0; i < len(patientSymptoms); i++ {
		for symptoms := range desease.Symptoms {
			if symptoms == patientSymptoms[i] {
				if desease.Symptoms[symptoms]["importance"] == "high" {
					count = count + 100
				} else if desease.Symptoms[symptoms]["importance"] == "moderate" {
					count = count + 50
				} else {
					count = count + 25
				}
			}
		}
	}
	return count >= thresholdDisease
}
