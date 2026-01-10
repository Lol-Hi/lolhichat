import React, { useState } from "react";
import { AxiosError } from "axios";
import { useNavigate } from "react-router-dom";

import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import FormLabel from "@mui/material/FormLabel";
import Input from "@mui/material/Input";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

import DirectToLogin from "../components/DirectToLogin";
import { useAuth } from "../hooks/useAuth";
import { useApiClient } from "../hooks/useApiClient";
import { errorMessage } from "../helpers/errorMessage";
import { NewThreadResponse } from "../api/apiResponse";

function NewThreadWithLogin(userToken: string) {
	const navigate = useNavigate();
	const apiClient = useApiClient();

	const [ topic, setTopic ] = useState("");
	const [ desc, setDesc ] = useState("");
	const [ errorMsg, setErrorMsg ] = useState("");
	const [ topicError, setTopicError ] = useState(false);
	const [ descError, setDescError ] = useState(false);

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setErrorMsg("");
		setTopicError(false);
		setDescError(false);
		
		try {
			const payload = JSON.stringify({topic, desc});
			const response = await apiClient.post<NewThreadResponse>("/api/newThread", payload);
			const contents = response.data;
			const urlCode = `/thread/${contents.urlCode}`;
			console.log(`New thread at "${urlCode}"`);
			setTimeout(() => navigate(urlCode), 1000);
		} catch (e) {
			const err = e as AxiosError;
			setErrorMsg(errorMessage(err));
		}
	}

	return (
		<div className="newThread">
			<Typography component="h4" variant="h4">chat what are we cooking today?</Typography>
			<Box
				component="form"
				onSubmit={handleSubmit}
			>
				<Stack spacing={1}>
					{ errorMsg && (<Alert severity="error">{errorMsg}</Alert>) }
					<FormControl>
						<FormLabel htmlFor="threadTopic" sx={{ textAlign: "left" }} >
							Topic
						</FormLabel>
						<Input
							error={topicError}
							id="thread-topic"
							type="text"
							name="thread-topic"
							placeholder="A title for this thread"
							color={topicError ? "error" : "primary"}
							value={topic}
							onChange={e => setTopic(e.target.value)}
						/>
					</FormControl>
					<FormControl>
						<FormLabel htmlFor="threadDesc" sx={{ textAlign: "left" }} >
							Desription
						</FormLabel>
						<TextField
							error={descError}
							id="thread-desc"
							type="text"
							name="thread-desc"
							placeholder="A short description for your thread"
							color={descError ? "error" : "primary"}
							value={desc}
							multiline={true}
							onChange={e => setDesc(e.target.value)}
						/>
					</FormControl>
					<Button type="submit" variant="contained">
						Create Thread
					</Button>
				</Stack>
			</Box>
		</div>
	);
}

function NewThread() {
	const { userToken } = useAuth();

	if (!userToken) {
		return DirectToLogin();
	}

	return NewThreadWithLogin(userToken);
}

export default NewThread;
