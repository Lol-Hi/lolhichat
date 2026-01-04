import React from "react"
import { Link } from "react-router-dom";

import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";

import Logout from "../pages/Logout";
import { useAuth } from "../hooks/useAuth"

function NavBar() {
	const { user } = useAuth();
	return (
		<AppBar position="static">
			<Toolbar>
				<Typography variant="h3" component="div" sx={{ flexGrow: 1 }}>Lol, Hi Chat!</Typography>
				<Button color="inherit" href="/draft">New Post</Button>
				{ !user
					? (<Button color="inherit" component={Link} to="/login">Login</Button>)
					: (<Logout color="inherit" />)
				}
			</Toolbar>
		</AppBar>
	);
}

export default NavBar;


