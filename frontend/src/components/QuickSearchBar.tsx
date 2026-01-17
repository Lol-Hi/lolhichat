import React from "react";
import { useNavigate } from "react-router-dom";

import Box from "@mui/material/Box";
import IconButton from "@mui/material/IconButton";
import Input from "@mui/material/Input";
import SearchIcon from "@mui/icons-material/Search";

function QuickSearchBar() {
	const navigate = useNavigate();

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
