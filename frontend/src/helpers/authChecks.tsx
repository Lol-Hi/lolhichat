const usernameRegex = /^[a-zA-Z0-9]+$/;
const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/;
const usernameMin = 4;
const usernameMax = 20;
const passwordMin = 8;
const passwordMax = 30;

export function isValidUsername(username: string): [boolean, string] {
	if(username.length < usernameMin || username.length > usernameMax) {
		return [false, `Username should be between ${usernameMin} and ${usernameMax} characters long`];
	}
	if(!usernameRegex.test(username)) {
		return [false, "Username should only contain alphanumeric characters"];
	}
	return [true, ""];
}

export function isValidPassword(password: string): [boolean, string] {
	if(password.length < passwordMin || password.length > passwordMax) {
		return [false, `Password should be between ${passwordMin} and ${passwordMax} characters long`];
	}
	if(!passwordRegex.test(password)) {
		return [
			false, 
			`Password should contain
				lowercase letters, 
				uppercase letters, 
				numbers and
				special characters`
		];
	}
	return [true, ""];
}
