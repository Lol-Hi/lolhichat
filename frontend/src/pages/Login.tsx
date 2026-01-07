import React, { useState } from "react";
import { AxiosError } from "axios";
import { useNavigate } from 'react-router-dom';

import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import FormLabel from "@mui/material/FormLabel";
import Input from "@mui/material/Input";
import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import { useAuth } from "../hooks/useAuth";
import { useApiClient } from "../hooks/useApiClient";
import { errorMessage } from "../helpers/errorMessage";
import { LoginResponse } from "../api/apiResponse";

function Login() {
	const [username, setUsername] = useState("");
	const [password, setPassword] = useState("");
	
	const [userError, setUserError] = useState(false);
	const [passError, setPassError] = useState(false);
 	const [errorMsg, setErrorMsg] = useState("");

	const navigate = useNavigate();
	const { login } = useAuth();
	const apiClient = useApiClient();

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setErrorMsg("");
		setUserError(false);
		setPassError(false);

		try {
			const payload = JSON.stringify({username, password});
			const response = await apiClient.post<LoginResponse>("/login", payload);
			const contents = response.data;
			login(contents.userToken, contents.refreshToken);
			setTimeout(() => navigate("/"), 1000);
		} catch (e) {
			const err = e as AxiosError;
			setErrorMsg(errorMessage(err));
		}
	}

	return (
		<div className="login">
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
					<Button type="submit" variant="contained">
						Log in
					</Button>
					<p>Or register for an account <Link href="/signUp">here</Link></p>
				</Stack>
			</Box>
		</div>
	);
}

export default Login;
