import axios from "axios";
/**
 * @import {AxiosInstance} from "axios"
 */

/**
 * Sets up the axios apiClient for http requests to be posted to the backend.
 * @constant {AxiosInstance}
 */
const apiClient = axios.create({
	baseURL: "http://localhost:8080",
	headers: {
		"Content-Type": "application/json"
	}
});

export default apiClient;
