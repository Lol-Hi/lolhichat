import { AxiosError } from "axios";
import { ErrorResponse } from "../api/apiResponse";

/**
 * Extracts the error message from the error produced.
 * @param {AxiosError} err The error produced by the backend.
 * @returns {string} The error message associated with the error.
 */
export function errorMessage(err: AxiosError): string {
	if(err.response) {
		const errorData = err.response.data as ErrorResponse;
		return `HTTP Error ${err.response.status}: ${errorData.error}`;
	} else if(err.request) {
		const errorData = err.request as XMLHttpRequest;
		return`Network Error ${errorData.status}: ${errorData.statusText}`;
	}
	return `Other Error: ${err.message}`;
}

