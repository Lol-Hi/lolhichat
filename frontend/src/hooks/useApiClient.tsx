import { useEffect } from "react";
import { AxiosError } from "axios";

import apiClient from "../api/apiClient";
import { ErrorResponse, LoginResponse } from "../api/apiResponse";
import { useAuth } from "../hooks/useAuth";

/**
 * Error message for an expired token
 * @constant {string}
 */
const ERR_EXP_TOKEN = "Expired User Token";

/**
 * Hook to implement request and response interceptors for the api client.
 * @returns {AxiosInstance} The modified api client that implements the request and response interceptors.
 */
export function useApiClient() {
	const { userToken, refreshToken, login } = useAuth();

	/**
	 * Sets up the request interceptor to add the JWT user token to the Authorization header.
	 * @param {string} userToken The JWT user token for the current login session.
	 * @return a function to eject the request interceptor.
	 */
	useEffect(() => {
		const authInterceptor = apiClient.interceptors.request.use(
			(config) => {
				if (userToken && config.headers) {
					config.headers.Authorization = `Bearer ${userToken}`;
				}
				return config;
			},
			(error) => Promise.reject(error)
		);

		return () => apiClient.interceptors.request.eject(authInterceptor);
	}, [userToken]);

	/**
	 * Sets up the response interceptor to automatically send a token renewal request if the token has expired.
	 * @param {string} userToken The JWT user token for the current login session.
	 * @param {string} refreshToken The JWT refresh token to authenticate the renewal request.
	 * @param {(string, string) => void} login The function to update the JWT tokens of the current session.
	 */
	useEffect(() => {
		const refreshInterceptor = apiClient.interceptors.response.use(
			(config) => config,
			async (error) => {
				const err = error as AxiosError;
				if (err.config && err.response) {
					const origRequest = err.config;
					const errorData = err.response.data as ErrorResponse;
					if (err.response.status === 401 && errorData.error === ERR_EXP_TOKEN && !origRequest._retry) {
						origRequest._retry = true;
						try {
							const payload = JSON.stringify({ userToken, refreshToken });
							const response = await apiClient.post<LoginResponse>("/renew", payload);
							const contents = response.data;
							login(contents.userToken, contents.refreshToken);
							window.location.reload(); // TODO: route to the new url
						} catch (e) {
							return Promise.reject(e);
						}
					}
				}
				return Promise.reject(error);
			}
		);
		return () => apiClient.interceptors.response.eject(refreshInterceptor);
	}, [userToken, refreshToken, login]);
	
	return apiClient;
}

