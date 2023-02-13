<script lang="ts">
	import ipbuf from './ipbuf';

	// The input from the textarea
	let input = '';
	// Cache the last input to prevent unnecessary requests
	let lastInput = '';

	// The promise returned by the API call
	let ipbufPromise: Promise<string> = Promise.resolve('');

	function buttonClickHandler() {
		// Check if the input is empty or the same as the last input
		if (input === '' || input === lastInput) {
			return;
		}

		// Update the last input
		lastInput = input;

		// Make the request to the API
		ipbufPromise = ipbuf(input);
		// Use #await to render the promise
	}
</script>

<h1>IPBuf</h1>
<div class="input container">
	<textarea placeholder="Enter your message here" class="input" bind:value={input} />
	<button class="btn btn-primary" on:click={buttonClickHandler}>Submit</button>
</div>
<div class="output container">
	{#await ipbufPromise}
		<textarea disabled>Loading...</textarea>
	{:then value}
		<textarea disabled>{value}</textarea>
	{:catch error}
		<textarea disabled>"Error: {error.message}"</textarea>
	{/await}
</div>

<style>
	h1 {
		text-align: center;
	}

	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
	}

	.input button {
		width: 500px;
		min-width: 25%;
		max-width: 80%;

		margin: 10px;
	}

	textarea {
		width: 500px;
		min-width: 25%;
		max-width: 80%;

		height: 200px;
		min-height: 10vh;
		max-height: 40vh;
	}
</style>
