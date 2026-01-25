/**
 * Regex expression to check if the username entered is valid.
 * 
 * A username can only contain alphanumeric characters.
 * 
 * @constant {RegExp}
 */
const usernameRegex = /^[a-zA-Z0-9]+$/;

/**
 * Regex expression to check if passwored entered is valid.
 * 
 * A password is valid if it contains:
 * 	-	At least 1 lowercase letter.
 *  -	At least 1 uppercase letter.
 * 	- 	At least 1 numeric character.
 * 	-	At least 1 special character.
 * 
 * @constant {RegExp}
 */
const passwordRegex = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/;

/**
 * Minimum length of the username.
 * @constant {number}
 */
const usernameMin = 4;

/**
 * Maxmimum length of the username.
 * @constant {number}
 */
const usernameMax = 20;

/**
 * Minimum length of password.
 * @constant {number}
 */
const passwordMin = 8;

/**
 * Maximum length of number
 * @constant {number}
 */
const passwordMax = 30;

/**
 * Checks if the entered username is valid.
 * @param {string} username The entered username to be checked.
 * @returns {[boolean, string]} A boolean value of the validity of the username, and an error message on failure.
 */
export function isValidUsername(username: string): [boolean, string] {
	if(username.length < usernameMin || username.length > usernameMax) {
		return [false, `Username should be between ${usernameMin} and ${usernameMax} characters long`];
	}
	if(!usernameRegex.test(username)) {
		return [false, "Username should only contain alphanumeric characters"];
	}
	return [true, ""];
}

/**
 * Checks if the entered password is valid.
 * @param {string} password The entered password to be checked.
 * @returns {[boolean, string]} A boolean value of the validity of the password, and an error message on failure.
 */
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
