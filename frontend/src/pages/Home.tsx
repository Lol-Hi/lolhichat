import React, { useState, useEffect } from "react";
import { AxiosError } from "axios";

import { useAuth } from "../hooks/useAuth";
import { useApiClient } from "../hooks/useApiClient";

import { HomeResponse } from "../api/apiResponse";
import { errorMessage } from "../helpers/errorMessage";

/**
 * A default homepage when the user has not logged in.
 * @returns {JSX.Element}
 */
function HomeNoLogin() {
	return (
		<div>
			<h1>Welcome to Lol, Hi Chat!</h1>
			<p>
				See what others are cooking when u <a href="/login">log in</a> to LolHiChat!
			</p>
		</div>
	);
}

/**
 * A homepage that greets the user after logging in.
 * @param {string} userToken The JWT user token of the current user session.
 * @returns {JSX.Element}
 */
function HomeWithLogin(userToken: string) {
	const [username, setUsername] = useState("");
	const [errorMsg, setErrorMsg] = useState("");
	
	const apiClient = useApiClient();

	useEffect(() => {
		const fetchUsername = async () => {
			try {
				const response = await apiClient.get<HomeResponse>("/api/home");
				const content = response.data;
				setUsername(content.username);
			} catch (err) {
				setErrorMsg(errorMessage(err as AxiosError));
			}
		};
		fetchUsername();
	}, [apiClient]);

	if(errorMsg) {
		return (<p> Error: { errorMsg }</p>);
	}

	return (
		<div>
			<h1>Welcome to Lol, Hi Chat!, { username }</h1>
		</div>
	);
}

/**
 * The homepage of lolhichat.
 * 
 * Checks the authorization status of the user before deciding which homepage to display.
 * 
 * @returns {JSX.Element}
 */
function Home() {
	const { userToken } = useAuth();
	
	if(!userToken) {
		return HomeNoLogin();
	}
	return HomeWithLogin(userToken);
}

export default Home;
