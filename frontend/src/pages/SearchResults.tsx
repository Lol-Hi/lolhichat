import React, { useState, useEffect } from "react";
import { AxiosInstance, AxiosError } from "axios";
import { useNavigate, useSearchParams } from "react-router-dom";
import ReactTimeAgo from "react-time-ago";

import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import Input from "@mui/material/Input";
import Typography from "@mui/material/Typography";
import SearchIcon from "@mui/icons-material/Search";
import { useApiClient } from "../hooks/useApiClient";

import { SearchResponse } from "../api/apiResponse";
import { errorMessage } from "../helpers/errorMessage";

function SearchResults() {
	const [searchParams, setSearchParams] = useSearchParams();
	const origQuery = searchParams.get("search");;
	const [query, setQuery] = useState(origQuery);
	const [results, setResults] = useState<React.ReactElement[]>([]);
	const [errorMsg, setErrorMsg] = useState("");

	const apiClient = useApiClient();
	const navigate = useNavigate();

	async function fetchSearch() {
		try {
			const payload = JSON.stringify({ query });
			const response = await apiClient.post<SearchResponse[]>("/search", payload);
			const contents = response.data;
			contents.forEach(res => {
				console.log(`Topic: ${res.topic} Description: ${res.description} Created at ${res.createdAt} by ${res.host}`);
			});
			setResults(contents.map(res => (
				<Box 
					id={res.urlCode} 
					onClick={() => navigate(`/thread/${res.urlCode}`)}
				>
					<Typography component="h6" variant="h6">{res.topic}</Typography>
					<p>{res.description}</p>
					<p>created by {res.host} <ReactTimeAgo date={new Date(res.createdAt)} /></p>
				</Box>	
			)));
		} catch (err) {
			setErrorMsg(errorMessage(err as AxiosError));
		}
	};

	useEffect(() => {
		if(origQuery !== "") {
			fetchSearch();
		}
	}, []);

	function handleSearch(event: React.FormEvent<HTMLFormElement>) {
		event.preventDefault();
		if(query !== null && query !== "") {
			const newSearchParams = new URLSearchParams(searchParams);
			newSearchParams.set("search", query);
			setSearchParams(newSearchParams);
			fetchSearch();
		}
	};

	if(errorMsg) {
		return (<p>Error: { errorMsg }</p>);
	}

	return (
		<div className="SearchResults">
			<Box component="form" onSubmit={handleSearch}>
				<Input 
					id="search"
					type="search"
					name="search"
					value={query}
					color="primary"
					onChange={e => setQuery(e.target.value)}
				/>
				<IconButton type="submit">
					<SearchIcon />
				</IconButton>
			</Box>
			<p>Showing { results.length } search results for "{ origQuery }"</p>
			
			<div className="threadResults">
				{ results }
			</div>
		</div>
	);
}

export default SearchResults;
