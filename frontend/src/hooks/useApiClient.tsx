import { useEffect } from "react";
import { AxiosError } from "axios";

import apiClient from "../api/apiClient";
import { ErrorResponse, LoginResponse } from "../api/apiResponse";
import { useAuth } from "../hooks/useAuth";

const ERR_EXP_TOKEN = "Expired User Token";


export function useApiClient() {
	const { userToken, refreshToken, login } = useAuth();

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

	useEffect(() => {
		const refreshInterceptor = apiClient.interceptors.response.use(
			(config) => config,
			async (error) => {
				console.log("refresh interceptor invoked");
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
							window.location.reload();
						} catch (e) {
							return Promise.reject(e);
						}
					}
				}
				return Promise.reject(error);
			}
		);
		return () => apiClient.interceptors.response.eject(refreshInterceptor);
	}, [userToken, refreshToken]);
	
	return apiClient;
}

