import React from "react";
import { useNavigate } from 'react-router-dom';

import Button from "@mui/material/Button";

import { useAuth } from "../hooks/useAuth";

type muiColor = 'inherit' | 'primary' | 'secondary' | 'error' | 'info' | 'success' | 'warning';

/**
 * A logout button on the navBar.
 * 
 * Navigates to the /logout link for the logout to be handled there
 * 
 * @param {object} props
 * @param {muiColor} color The colour of the logout button
 * @returns {JSX.Element}
 */
function LogoutButton({ color }: { color: muiColor }) {
	const navigate = useNavigate();
	const { logout } = useAuth();

	const handleLogout = () => {
		logout();
		navigate("/login");
	};
	
	return (
		<Button color={color} onClick={handleLogout}>Logout</Button>
	)
}

export default LogoutButton;
