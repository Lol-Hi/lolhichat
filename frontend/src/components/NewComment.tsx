import React, { useState } from "react";
import { AxiosError } from "axios";

import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import FormLabel from "@mui/material/FormLabel";
import FormHelperText from "@mui/material/FormHelperText";
import Grid from "@mui/material/Grid";
import OutlinedInput from "@mui/material/OutlinedInput";
import Typography from "@mui/material/Typography";

import DirectToLogin from "../components/DirectToLogin";
import { useAuth } from "../hooks/useAuth";
import { useApiClient } from "../hooks/useApiClient";
import { errorMessage } from "../helpers/errorMessage";

function NewCommentWithLogin(userToken: string, urlCode: string) {
	const apiClient = useApiClient();

	const [ comment, setComment ] = useState("");
	const [ errorMsg, setErrorMsg ] = useState("");
	const [ commentError, setCommentError ] = useState(false);

	async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		setErrorMsg("");
		setCommentError(false);
		
		try {
			const payload = JSON.stringify({ comment, urlCode });
			const response = await apiClient.post("/api/newComment", payload);
			window.location.reload();
		} catch (e) {
			const err = e as AxiosError;
			setErrorMsg(errorMessage(err));
			setCommentError(true);
		}
	}

	return (
		<div className="newComment">
			<Box component="form" onSubmit={handleSubmit}>
				<FormControl fullWidth>
					<OutlinedInput
						error={commentError}
						id="thread-comment"
						type="text"
						name="thread-comment"
						placeholder="What are we cooking?"
						color={commentError ? "error" : "primary"}
						value={comment}
						onChange={e => setComment(e.target.value)}
						endAdornment={
							<Button type="submit" color="primary">Post</Button>
						}
					/>
					{ commentError && (
						<FormHelperText error={true}>{ errorMsg }</FormHelperText>
					)}
				</FormControl>
			</Box>
		</div>
	);
}

function NewComment({ urlCode }: { urlCode: string  }) {
	const { userToken } = useAuth();

	if (!userToken) {
		return DirectToLogin();
	}

	return NewCommentWithLogin(userToken, urlCode);
}
export default NewComment;
