import React, { createContext, useState, useEffect, useContext, ReactNode } from "react";

interface AuthContextType {
	userToken: string | null;
	refreshToken: string | null;
	loading: boolean;
	login: (userToken: string, refreshToken: string) => void;
	logout: () => void;
}

const AuthContext = createContext<AuthContextType>({
	userToken: "",
	refreshToken: "",
	loading: true,
	login: () => null,
	logout: () => null
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [ userToken, setUser ] = useState("");
	const [ refreshToken, setRefresh ] = useState("");
	const [ loading, setLoading ] = useState(true);

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

  const login = (userToken: string, refreshToken: string) => {
    setUser(userToken);
		setRefresh(refreshToken);
		localStorage.setItem("user", userToken);
		localStorage.setItem("refresh", refreshToken);
  };

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

export function useAuth() {
	const context = useContext(AuthContext);
	if(!context) {
		throw Error("improper placement of AuthContext");
	}
	return context;
}
