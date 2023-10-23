package request

type LoginRequest struct {
	Password string `json:"password" binding:"required"`
	Nip      string `json:"nip" binding:"required"`
}
