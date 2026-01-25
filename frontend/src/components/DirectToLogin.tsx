import React from "react";

/**
 * A token component to direct users that are not logged in to the login page.
 * @returns {JSX.Element}
 */
function DirectToLogin() {
	return (
		<div>
			<p><a href="/login">Log in</a> to access more advanced features of LolHiChat!</p> 
		</div>
	);
}

export default DirectToLogin;
