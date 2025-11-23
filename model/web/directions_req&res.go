package web

type DirectionRequest struct {
	Step uint `json:"step_direction" validate:"required"`
	Description string `json:"description_direction" validate:"required"`
	Image string `json:"image_direction"`
}

type DirectionPatchRequest struct {
	Step *uint `json:"step_direction" `
	Description *string `json:"description_direction"`
	Image *string `json:"image_direction" `
}

type DirectionResponse struct {
	Step uint `json:"step_direction" `
	Description string `json:"description_direction"`
	Image string `json:"image_direction" `
}