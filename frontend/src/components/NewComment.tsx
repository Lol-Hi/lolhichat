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

/**
 * A component to allow logged-in users to draft comments.
 * @param {string} userToken The JWT user token of the curent session.
 * @param {string} urlCode The url code of the curren thread.
 * @returns 
 */
function NewCommentWithLogin(userToken: string, urlCode: string) {
	const apiClient = useApiClient();

	const [ comment, setComment ] = useState("");
	const [ errorMsg, setErrorMsg ] = useState("");
	const [ commentError, setCommentError ] = useState(false);

	/**
	 * New comment handler.
	 * 
	 * Sends a new comment POST request to the backend,
	 * and refreshes the page on success.
	 * 
	 * @param {React.FormEvent<HTMLFormElement>} event The form submission event which triggered this function.
	 */
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

/**
 * A component to check if users are logged in before allowing them to draft comments.
 * @param {object} props
 * @param {string} props.urlCode The current url code of the thread. 
 * @returns {JSX.Element}
 */
function NewComment({ urlCode }: { urlCode: string  }) {
	const { userToken } = useAuth();

	if (!userToken) {
		return DirectToLogin();
	}

	return NewCommentWithLogin(userToken, urlCode);
}
export default NewComment;
