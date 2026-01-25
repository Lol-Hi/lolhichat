/**
 * Representation of an error response from the backend.
 * @interface 
 */
export interface ErrorResponse {
	/**
	 * Error message of the error
	 * @type {string}
	 */
	error: string;
}

/**
 * Representation of a login response from the backend.
 * @interface
 */
export interface LoginResponse {
	/**
	 * The JWT user token for the user session.
	 * @type {string}
	 */
	refreshToken: 	string;

	/**
	 * The JWT refresh token to renew the user token after it expires.
	 * @type {string}
	 */
	userToken: 		string;
}

/**
 * Representation of a response from the backend containing information about the homepage.
 * @interface
 */
export interface HomeResponse {
	/**
	 * The username of the currently logged-in user.
	 * @type {string}
	 */
	username: string;
}

/**
 * Representation of a response from the backend for a new thread request.
 * @interface
 */
export interface NewThreadResponse {
	/**
	 * The urlCode of the newly ceated thread.
	 * @type {string}
	 */
	urlCode: string;
}

/**
 * Representation of a response from the backend for a view thread request.
 * @interface
 */
export interface ViewThreadResponse {
	/**
	 * The topic title of the thread.
	 * @type {string}
	 */
	topic: 			string;

	/**
	 * The description of the thread.
	 * @type {string}
	 */
	description: 	string;

	/**
	 * The username of the host of the thread.
	 * @type {string}
	 */
	host: 			string;

	/**
	 * The time at which the thread is created.
	 * @type {string}
	 */
	createdAt: 		string;

	/**
	 * The list of comments posted in the thread.
	 * @type {ThreadCommentResponse[]}
	 */
	comments:		ThreadCommentResponse[];
}

/**
 * Respresentation of a response from the backend for a comment in a thread.
 * @interface
 */
export interface ThreadCommentResponse {
	/**
	 * The username of the author of the comment.
	 * @type {string}
	 */
	author:			string;

	/**
	 * The content of the comment.
	 * @type {string}
	 */
	content:		string;

	/**
	 * The urlCode of the comment.
	 * @type {string}
	 */
	urlCode:		string;

	/**
	 * The number of likes received by the comment
	 * @type {number}
	 */
	likes:			number;

	/**
	 * The time at which the comment is created.
	 * @type {string}
	 */
	createdAt:		string;
}

/**
 * Representation of a response from the backend for whether the user has liked the comment
 * @interface
 */
export interface LikeStatusResponse {
	/**
	 * Boolean response to indicate if the user has liked the comment.
	 * @type {boolean}
	 */
	isLiked: boolean;
}

/**
 * Representation of a response from the backend for a search request of threads.
 * @interface
 */
export interface SearchResponse {
	/**
	 * The topic title of the thread.
	 * @type {string}
	 */
	topic: 			string;

	/**
	 * The description of the thread.
	 * @type {string}
	 */
	description: 	string;

	/**
	 * The username for the host of the thread.
	 * @type {string}
	 */
	host: 			string;

	/**
	 * The urlCode of the thread.
	 * @type {string}
	 */
	urlCode: 		string;

	/**
	 * The time at which the thread is created.
	 * @type {string}
	 */
	createdAt: 		string;
}
