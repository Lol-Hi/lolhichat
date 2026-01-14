import React from "react";
import { useParams } from 'react-router-dom';

import Typography from "@mui/material/Typography";

function ViewComment() {
	const { urlCode } = useParams();

	return (
		<div className="viewComment">
			<Typography component="h4" variant="h4">Comments for { urlCode }</Typography>
		</div>
	);
}

export default ViewComment;
