/**
 * IPBuf API wrapper
 */

const url = 'https://ipbuf.azurewebsites.net/';
// const url = ' http://localhost:7071/'; // For local testing
const endpoint = url + 'api/httptriggeripbuf';

/**
 * Sends a POST request to the IPBuf API
 * @param data The data to send to the API
 * @returns Response protobuf as a hex string
 */
async function ipbuf(data: string) {
	const response = await fetch(endpoint, {
		method: 'POST',
		body: data,
		headers: { 'Content-Type': 'text/plain' }
	});

	// Convert the response to a Uint8Array
	const arrayBuffer = await response.arrayBuffer();
	const uint8Array = new Uint8Array(arrayBuffer);

	// Convert the Uint8Array to a hex string
	const hexString = Array.from(uint8Array, (byte) => {
		return ('0' + (byte & 0xff).toString(16)).slice(-2);
	}).join('');

	return '0x' + hexString;
}

export default ipbuf;
