import React, { useState } from "react";
import apiClient from "../api/apiClient";
import { AxiosError } from "axios";
import { useNavigate } from 'react-router-dom';

import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import FormLabel from "@mui/material/FormLabel";
import FormControlLabel from "@mui/material/FormControlLabel";
import Input from "@mui/material/Input";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import { isValidUsername, isValidPassword } from "../helpers/authChecks";
import { ErrorResponse } from "../api/apiResponse";

function SignUp() {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	const [confPass, setConfPass] = useState("");

	const [userError, setUserError] = useState(false);
	const [passError, setPassError] = useState(false);
	const [confError, setConfError] = useState(false);
 	const [errorMsg, setErrorMsg] = useState("");

	const navigate = useNavigate();

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setErrorMsg("");
		setUserError(false);
		setPassError(false);
		setConfError(false);
		
		const [isValidUser, userErrorMsg] = isValidUsername(username);
		if(!isValidUser) {
			setErrorMsg(userErrorMsg);
			setUserError(true);
			return;
		}
		const [isValidPass, passErrorMsg] = isValidPassword(password);
		if(!isValidPass) {
			setErrorMsg(passErrorMsg);
			setPassError(true);
			return;
		}
		
		if(password !== confPass) {
			setErrorMsg("Password and Confirm Password does not match");
			setConfError(true);
			return;
		}

		try {
			const payload = JSON.stringify({username, password});
			const response = await apiClient.post("/signUp", payload);
			setTimeout(() => navigate("/login"), 500);
		} catch (e) {
			const err = e as AxiosError;
			if(err.response) {
				const errorData = err.response.data as ErrorResponse;
				setErrorMsg(`HTTP Error ${err.response.status}: ${errorData.error}`);
			} else if(err.request) {
				setErrorMsg(`Network Error: ${err.request}`);
			} else {
				setErrorMsg(`Other Error: ${err.message}`);
			}
		}
	}

	return (
		<div className="signUp">
			<Typography component="h4" variant="h4">Welcome back</Typography>
			<Box 
				component="form" 
				onSubmit={handleSubmit}
			>
				<Stack spacing={1}>
					{ errorMsg && (<Alert severity="error">{errorMsg}</Alert>) }
					<FormControl>
					<FormLabel htmlFor="username">Username</FormLabel>
						<Input
							error={userError}
							id="username"
							type="username"
							name="username"
							placeholder="Username"
							color={userError ? "error" : "primary"}
							value={username}
							onChange={e => setUsername(e.target.value)}
						/>
					</FormControl>
					<FormControl>
						<FormLabel htmlFor="password">Password</FormLabel>
						<Input
							error={passError}
							id="password"
							type="password"
							name="password"
							placeholder="Password"
							color={passError ? "error" : "primary"}
							value={password}
							onChange={e => setPassword(e.target.value)}
						/>
					</FormControl>
					<FormControl>
						<FormLabel htmlFor="password">Confirm Password</FormLabel>
						<Input
							error={confError}
							id="confPass"
							type="password"
							name="confPass"
							placeholder="Confirm Password"
							color={confError ? "error" : "primary"}
							value={confPass}
							onChange={e => setConfPass(e.target.value)}
						/>
					</FormControl>
					<Button type="submit" variant="contained">
						Sign Up
					</Button>
				</Stack>
			</Box>
		</div>
	);
}

export default SignUp;
