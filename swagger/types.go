package swagger

// GenericResponse is a lightweight response envelope used for Swagger docs.
// Many endpoints in this project respond with {success,message,data}.
type GenericResponse struct {
	Success bool                   `json:"success" example:"true"`
	Message string                 `json:"message" example:""`
	Data    map[string]interface{} `json:"data"`
}

type GenericErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"error"`
}
