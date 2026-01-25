import React from "react";
import { useNavigate } from "react-router-dom";

import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import Input from "@mui/material/Input";
import SearchIcon from "@mui/icons-material/Search";

/**
 * A simple search bar to be placed in the navigation bar.
 * @returns {JSX.Element}
 */
function QuickSearchBar() {
	const navigate = useNavigate();

	/**
	 * Navigate to the search page to process the search request.
	 */
	function handleSearch() {
		navigate("/search");
	}

	return (
		<div className="quickSearchBar">
			<Box component="form" onSubmit={handleSearch}>
				<Input 
					id="search"
					type="search"
					name="search"
					placeholder="Quick Search"
					color="primary"
					endAdornment={
						<IconButton type="submit">
							<SearchIcon />
						</IconButton>
					}
				/>
			</Box>
		</div>
	);
}

export default QuickSearchBar;
