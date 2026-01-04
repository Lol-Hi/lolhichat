import React, { createContext, useState, useEffect, useContext, ReactNode } from "react";

interface AuthContextType {
	user: string | null;
	loading: boolean;
	login: (token: string) => void;
	logout: () => void;
}

const AuthContext = createContext<AuthContextType>({
	user: "",
	loading: true,
	login: () => null,
	logout: () => null
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [ user, setUser ] = useState("");
	const [ loading, setLoading ] = useState(true);

  useEffect(() => {
    const storedUser = localStorage.getItem("user")
		console.log(`auth level: adding user ${storedUser} to local user object`);
    if (storedUser) {
			setUser(storedUser);
    }
		setLoading(false);
  }, []);

  const login = (userToken: string) => {
    setUser(userToken);
		localStorage.setItem("user", userToken);
  };

  const logout = () => {
    setUser("");
		localStorage.removeItem("user");
  };

  return (
		<AuthContext.Provider value={{ user, loading, login, logout }}> 
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
