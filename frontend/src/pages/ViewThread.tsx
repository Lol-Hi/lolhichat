import React, { useState, useEffect } from "react";
import { AxiosError } from "axios";
import { useParams, useNavigate } from "react-router-dom";
import ReactTimeAgo from "react-time-ago";

import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardHeader from "@mui/material/CardHeader";
import CardContent from "@mui/material/CardContent";
import Grid from "@mui/material/Grid";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import NewComment from "../components/NewComment";
import CommentCard from "../components/CommentCard";

import { useApiClient } from "../hooks/useApiClient";

import { ViewThreadResponse } from "../api/apiResponse";
import { errorMessage } from "../helpers/errorMessage";

function ViewThread() {
	const { urlCode } = useParams();

	const [threadInfo, setThreadInfo] = useState<React.ReactElement>(<div />);
	const [threadComments, setThreadComments] = useState<React.ReactElement[]>([]);
	const [errorMsg, setErrorMsg] = useState("");

	const apiClient = useApiClient();

	useEffect(() => {
		const fetchThread = async () => {
			try {
				const response = await apiClient.get<ViewThreadResponse>(`/viewThread/${urlCode}`);
				const content = response.data;
				setThreadInfo((
					<Card variant="outlined">
						<CardContent>
							<Typography component="h4" variant="h4" align="center">
								{content.topic}
							</Typography>
							<Typography variant="body2" sx={{ color: "text.secondary" }}>
								created by {content.host} <ReactTimeAgo date={new Date(content.createdAt)} />
							</Typography>
							<Typography variant="body1" align="left">
								{content.description}
							</Typography>
						</CardContent>
					</Card>
				));
				setThreadComments(content.comments.map(res => (
					<CommentCard tcResponse={res} />
				)));
			} catch (err) {
				setErrorMsg(errorMessage(err as AxiosError));
			}
		};
		fetchThread();
	}, [apiClient, urlCode]);

	if(urlCode === undefined) {
		return (<p>Error 404: Page not found</p>);
	}

	if(errorMsg) {
		return (<p>Error: { errorMsg }</p>);
	}

	return (
		<div className="viewThread">
			<Box className="threadInfo" sx={{ p: 2 }}>
				{ threadInfo }
			</Box>
			<Box className="draftComment" sx={{ p: 2 }}>
				<Typography component="h6" variant="h6" sx={{ textAlign: "left" }}>
					Leave your mark!
				</Typography>
				<NewComment urlCode={urlCode} />
			</Box>
			<Box className="threadComments" sx={{ p: 2 }}>
				<Typography component="h6" variant="h6" sx={{ textAlign: "left" }}>
					View all comments
				</Typography>
				<Stack spacing={2}>
					{ threadComments }
				</Stack>
			</Box>
		</div>
	);
}

export default ViewThread;
