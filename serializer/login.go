package serializer

type LoginResponse struct {
	User  UserResponse `json:"user" form:"user"`
	Token string       `json:"token" form:"token"`
}

func BuildLoginResponse(userResp UserResponse, token string) LoginResponse {
	return LoginResponse{
		User:  userResp,
		Token: token,
	}
}
