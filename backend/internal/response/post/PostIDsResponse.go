package response

type PostIDsResponse struct {
	IDs    []uint `json:"ids"`
	UserID []uint `json:"user_id"`
}
