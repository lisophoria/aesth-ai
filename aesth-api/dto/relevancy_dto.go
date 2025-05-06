package dto

type RelevancyDTO struct {
	Smiling	float32 `json:"smiling"`
}

type RelevancyPairDTO struct {
	Left RelevancyDTO `json:"left"`
	Right RelevancyDTO `json:"right"`
}