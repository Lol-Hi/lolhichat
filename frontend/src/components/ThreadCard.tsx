import React from "react";
import { useNavigate } from "react-router-dom";
import ReactTimeAgo from "react-time-ago";

import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardContent from "@mui/material/CardContent";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";

import { SearchResponse } from "../api/apiResponse";

/**
 * A card to store information about the thread.
 * @param {object} props
 * @param {SearchResponse} props.threadResponse The information about the thread.
 * @returns {JSX.Element}
 */
function ThreadCard({ threadResponse }: { threadResponse: SearchResponse }) {
	const navigate = useNavigate();
	
	return (
		<Grid size={4} className="threadCard" id={threadResponse.urlCode}>
			<Card variant="outlined">
				<CardActionArea onClick={() => navigate(`/thread/${threadResponse.urlCode}`)}>
					<CardContent sx={{ textAlign: 'left' }}>
						<Typography component="h6" variant="h6">
							{threadResponse.topic}
						</Typography>
						<Typography variant="body1">
							{threadResponse.description}
						</Typography>
						<Typography variant="body2" sx={{ color: "text.secondary" }}>
							created by {threadResponse.host} <ReactTimeAgo date={new Date(threadResponse.createdAt)} />
						</Typography>
					</CardContent>
				</CardActionArea>
			</Card>
		</Grid>
	);
}

export default ThreadCard;
