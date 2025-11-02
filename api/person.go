package api

type PersonRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonResponse struct {
	ID        int    `json:"person_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}
