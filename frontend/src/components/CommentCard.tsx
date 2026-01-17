import React from "react";
import ReactTimeAgo from "react-time-ago";

import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import CardActionArea from "@mui/material/CardActionArea";
import CardContent from "@mui/material/CardContent";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";

import { ThreadCommentResponse } from "../api/apiResponse";

function CommentCard({ tcResponse }: { tcResponse: ThreadCommentResponse }) {
	return (
		<Box className="commentCard" id={tcResponse.urlCode}>
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
		</Box>
	);
}

export default CommentCard;
