import React, { useState, useEffect } from "react";
import { AxiosInstance, AxiosError } from "axios";
import { useSearchParams } from "react-router-dom";

import Box from "@mui/material/Box";
import FormControl from "@mui/material/FormControl";
import Grid from "@mui/material/Grid";
import IconButton from "@mui/material/IconButton";
import OutlinedInput from "@mui/material/OutlinedInput";
import Typography from "@mui/material/Typography";
import SearchIcon from "@mui/icons-material/Search";

import ThreadCard from "../components/ThreadCard";

import { useApiClient } from "../hooks/useApiClient";

import { SearchResponse } from "../api/apiResponse";
import { errorMessage } from "../helpers/errorMessage";

/**
 * A page to display the results from a search
 * @returns {JSX.Element}
 */
function SearchResults() {
	const [searchParams, setSearchParams] = useSearchParams();
	const origQuery = searchParams.get("search");;
	const [query, setQuery] = useState(origQuery);
	const [results, setResults] = useState<React.ReactElement[]>([]);
	const [errorMsg, setErrorMsg] = useState("");

	const apiClient = useApiClient();

	/**
	 * Search request handler. 
	 * 
	 * Sends a search POST request to the backend,
	 * and wraps each result in a ThreadCard object to be displayed on the page.
	 */
	async function fetchSearch() {
		try {
			const payload = JSON.stringify({ query });
			const response = await apiClient.post<SearchResponse[]>("/search", payload);
			const contents = response.data;
			setResults(contents.map(res => (
				<ThreadCard threadResponse={res} key={res.urlCode} />
			)));
		} catch (err) {
			setErrorMsg(errorMessage(err as AxiosError));
		}
	};

	/**
	 * Call search request handler on page load. 
	 */
	useEffect(() => {
		if(origQuery !== "") {
			fetchSearch();
		}
	}, []);

	/**
	 * Call search request handler on form submission.
	 * 
	 * @param {React.FormEvent<HTMLFormElement>} event The form submission event that triggered this function.
	 */
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
		<div className="searchResults">
			<Box component="form" onSubmit={handleSearch}>
				<FormControl sx={{ width: "75%", textAlign: "left" }}>
					<OutlinedInput fullWidth
						id="search"
						type="search"
						name="search"
						value={query}
						color="primary"
						onChange={e => setQuery(e.target.value)}
						endAdornment={
							<IconButton type="submit">
								<SearchIcon />
							</IconButton>
						}
					/>
					<Typography variant="body2" sx={{ color: "text.secondary" }}>
						Showing { results.length } search results for "{ origQuery }"
					</Typography>
				</FormControl>
			</Box>
			
			<Box className="threadResults" sx={{ p: 2 }}>
				<Grid container spacing={2}>
					{ results }
				</Grid>
			</Box>
		</div>
	);
}

export default SearchResults;
