import React, { useState, useEffect } from "react";
import ReactTimeAgo from "react-time-ago";
import { AxiosError } from "axios";

import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardContent from "@mui/material/CardContent";
import Fab from "@mui/material/Fab";
import FavoriteIcon from "@mui/icons-material/Favorite";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import { useAuth } from "../hooks/useAuth";
import { useApiClient } from "../hooks/useApiClient";

import { ThreadCommentResponse, LikeStatusResponse } from "../api/apiResponse";
import { errorMessage } from "../helpers/errorMessage";

function CommentCard({ tcResponse }: { tcResponse: ThreadCommentResponse }) {
	const apiClient = useApiClient();
	const { userToken } = useAuth();

	const [ liked, setLiked ] = useState(false);
	const [ numLikes, setNumLikes ] = useState(tcResponse.likes);
	const [ error, setError ] = useState(false);

	useEffect(() => {
		const fetchLikeStatus = async () => {
			setError(false);
			try {
				const response = await apiClient.get<LikeStatusResponse>(`/api/commentLiked/${tcResponse.urlCode}`);
				const content = response.data;
				setLiked(content.isLiked);
			} catch (err) {
				console.log(errorMessage(err as AxiosError));
				setError(true);
			}
		}
		if(userToken) {
			fetchLikeStatus();
		}
	}, [])

	async function handleLike() {
		if(liked) {
			setLiked(false);
			setNumLikes(numLikes - 1);
			try {
				const response = await apiClient.post(`/api/unlikeComment/${tcResponse.urlCode}`);
			} catch (err) {
				console.log(errorMessage(err as AxiosError));
				setError(true);
			}
		} else {
			setLiked(true);
			setNumLikes(numLikes + 1);
			try {
				const response = await apiClient.post(`/api/likeComment/${tcResponse.urlCode}`);
			} catch (err) {
				console.log(errorMessage(err as AxiosError));
				setError(true);
			}
		}
	}

	return (
		<Box className="commentCard" id={tcResponse.urlCode} sx={{ position: "relative" }}>
			<Card variant="outlined">
				<CardActionArea>
					<CardContent sx={{ textAlign: 'left' }}>
						<Stack direction="row" spacing={1}>
							<Typography variant="subtitle2" sx={{ fontWeight: "bold" }}>
								{tcResponse.author}
							</Typography>
							<Typography variant="body2">
								{tcResponse.content}
							</Typography>
							<Typography variant="body2" sx={{ color: "text.secondary" }}>
								posted <ReactTimeAgo date={new Date(tcResponse.createdAt)} />
							</Typography>
						</Stack>
					</CardContent>
				</CardActionArea>
			</Card>
			<Box 
				className="commentLikes"
        sx={{
          position: "absolute",
          bottom: 0,
          right: 0,
          transform: 'translate(-25%, 15%)',
        }}
			>
				<Stack direction="column">
					<Typography variant="caption">
						{ numLikes }
					</Typography>
					<Fab
						size="small"
						color={liked ? "error" : "default"}
						disabled={!userToken || error}
						aria-label="add"
						onClick={handleLike}
					>
						<FavoriteIcon />
					</Fab>
				</Stack>
			</Box>
		</Box>
	);
}

export default CommentCard;
