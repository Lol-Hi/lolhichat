import React from "react"
import { Link } from "react-router-dom";

import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import QuickSearchBar from "../components/QuickSearchBar";
import Logout from "../components/Logout";
import { useAuth } from "../hooks/useAuth"

/**
 * A standard navigation bar for the app.
 * 
 * Allows users to search for threads, create a new thread, and login/logout.
 * @returns {JSX.Element}
 */
function NavBar() {
	const { userToken } = useAuth();
	return (
		<div className="navBar">
			<AppBar position="static" sx={{ mb: 4 }}>
				<Toolbar>
					<Typography variant="h3" component="div" sx={{ flexGrow: 1 }}>Lol, Hi Chat!</Typography>
					<QuickSearchBar/>
					<Button color="inherit" href="/newThread">New Thread</Button>
					{ !userToken
						? (<Button color="inherit" component={Link} to="/login">Login</Button>)
						: (<Logout color="inherit" />)
					}
				</Toolbar>
			</AppBar>
		</div>
	);
}

export default NavBar;


