import React, { useState, useEffect } from "react";
import { AxiosError } from "axios";
import { useParams } from "react-router-dom";
import ReactTimeAgo from "react-time-ago";

import Typography from "@mui/material/Typography";

import { useApiClient } from "../hooks/useApiClient";

import { ViewThreadResponse } from "../api/apiResponse";
import { errorMessage } from "../helpers/errorMessage";

function ViewThread() {
	const { urlCode } = useParams();

	// const [topic, setTopic] = useState("");
	// const [desc, setDesc] = useState("");
	// const [host, setHost] = useState("");
	const [threadInfo, setThreadInfo] = useState<React.ReactElement>(<div />);
	const [errorMsg, setErrorMsg] = useState("");

	const apiClient = useApiClient();

	useEffect(() => {
		const fetchThread = async () => {
			try {
				const response = await apiClient.get<ViewThreadResponse>(`/viewThread/${urlCode}`);
				const content = response.data;
				// setTopic(content.topic);
				// setDesc(content.description);
				// setHost(content.host);
				setThreadInfo((
					<div className="threadInfo">
						<Typography component="h4" variant="h4">Thead: { content.topic }</Typography>
						<p>{ content.description }</p>
						<p>Created by { content.host } <ReactTimeAgo date={new Date(content.createdAt)} /></p>
					</div>
				));
			} catch (err) {
				setErrorMsg(errorMessage(err as AxiosError));
			}
		};
		fetchThread();
	}, [apiClient, urlCode]);

	if(errorMsg) {
		return (<p>Error: { errorMsg }</p>);
	}

	return (
		<div className="viewThread">
			{/*
			<Typography component="h4" variant="h4">Thead: { topic }</Typography>
			<p>{ desc }</p>
			<p>Created by { host }</p>
			*/}
			{ threadInfo }
		</div>
	);
}

export default ViewThread;
