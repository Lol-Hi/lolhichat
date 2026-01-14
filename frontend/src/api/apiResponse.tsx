export interface ErrorResponse {
	error: string;
}

export interface LoginResponse {
	refreshToken: string;
	userToken: 		string;
}

export interface HomeResponse {
	username: string;
}

export interface NewThreadResponse {
	urlCode: string;
}

export interface ViewThreadResponse {
	topic: 				string;
	description: 	string;
	host: 				string;
	createdAt: 		string;
}

export interface SearchResponse {
	topic: 				string;
	description: 	string;
	host: 				string;
	urlCode: 			string;
	createdAt: 		string;
}
