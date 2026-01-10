export interface ErrorResponse {
	error: string;
}

export interface LoginResponse {
	refreshToken: string;
	userToken: string;
}

export interface HomeResponse {
	username: string;
}

export interface NewThreadResponse {
	urlCode: string;
}
