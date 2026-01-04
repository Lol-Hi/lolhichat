import React from "react";
import apiClient from "../api/apiClient";
import { useNavigate } from 'react-router-dom';

import Button from "@mui/material/Button";

import { useAuth } from "../hooks/useAuth";

type muiColor = 'inherit' | 'primary' | 'secondary' | 'error' | 'info' | 'success' | 'warning';

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
