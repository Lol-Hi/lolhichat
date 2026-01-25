import React, { createContext, useState, useEffect, useContext, ReactNode } from "react";

/**
 * Representation of the return object of the useAuth hook.
 * @interface
 */
interface AuthContextType {
	/**
	 * The JWT user token of the current session.
	 * @type {string | null}
	 */
	userToken: string | null;

	/**
	 * The JWT refresh token to renew the user token.
	 * @type {string | null}
	 */
	refreshToken: string | null;

	/**
	 * A boolean value to check if the authentication of the session is still being updated.
	 * @type {boolean}
	 */
	loading: boolean;

	/**
	 * A function to update the JWT user and refresh tokens of the current session.
	 * @function
	 */
	login: (userToken: string, refreshToken: string) => void;

	/**
	 * A function to clear the current session of the saved JWT tokens.
	 * @function
	 */
	logout: () => void;
}

/**
 * The authorization context of the current session
 * @constant {AuthContextType}
 */
const AuthContext = createContext<AuthContextType>({
	userToken: "",
	refreshToken: "",
	loading: true,
	login: () => null,
	logout: () => null
});

/**
 * A component that wraps the app in an authorization context provider
 * @param {object} props
 * @param {ReactNode} props.children The contents of the app
 * @returns {JSX.Element} The contents of the app, wrapped in an authorization context provider
 */
export function AuthProvider({ children }: { children: ReactNode }) {
  	const [ userToken, setUser ] = useState("");
	const [ refreshToken, setRefresh ] = useState("");
	const [ loading, setLoading ] = useState(true);

	/**
	 * Ensures that the state of the user and refresh tokens are consistent with what is saved in localStorage before updating the loading status.
	 */
	useEffect(() => {
		const storedUser = localStorage.getItem("user");
		if (storedUser) {
			setUser(storedUser);
		}
		const storedRefresh = localStorage.getItem("refresh");
		if (storedRefresh) {
			setRefresh(storedRefresh);
		}
		setLoading(false);
	}, []);

	/**
	 * Updates the JWT user and refresh tokens of the current session
	 * @param {string} userToken The new JWT user token for the current context.
	 * @param {string} refreshToken The new JWT refresh token for the current context.
	 */
	const login = (userToken: string, refreshToken: string) => {
		setUser(userToken);
		setRefresh(refreshToken);
		localStorage.setItem("user", userToken);
		localStorage.setItem("refresh", refreshToken);
	};

	/**
	 * Clears the JWT user and refresh tokens of the current session
	 */
	const logout = () => {
		setUser("");
		setRefresh("");
		localStorage.removeItem("user");
		localStorage.removeItem("refresh");
	};

	return (
		<AuthContext.Provider value={{ userToken, refreshToken, loading, login, logout }}> 
			{ !loading && children } 
		</AuthContext.Provider>
	);
};

/**
 * A hook for the app to interact with the authorization context.
 * @returns {AuthContextType} As defined above.
 */
export function useAuth() {
	const context = useContext(AuthContext);
	if(!context) {
		throw Error("improper placement of AuthContext");
	}
	return context;
}
