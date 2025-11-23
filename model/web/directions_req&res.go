package web

type DirectionRequest struct {
	Step uint `json:"step" validate:"required"`
	Description string `json:"description_direction" validate:"required"`
	Img string `json:"Img" validate:"required"`
}

type DirectionPatchRequest struct {
	Step *uint `json:"step" `
	Description *string `json:"description_direction"`
	Img *string `json:"Img" `
}

type DirectionResponse struct {
	Step uint `json:"step" `
	Description string `json:"description_direction"`
	Img string `json:"Img" `
}