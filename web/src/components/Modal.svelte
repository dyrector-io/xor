<script lang="ts">
	import { fly, fade } from 'svelte/transition';
	import { createEventDispatcher } from 'svelte';
	import Button from '../components/Button.svelte';

	const dispatch = createEventDispatcher();
	export let scorePoints: string;

	const copy = () => {
		const text = document.getElementById('resultText').innerText;
		navigator.clipboard.writeText(text);
	};
</script>

<div
	class="modal-wrapper flex flex-col  justify-center backdrop-blur-sm bg-gray-800 inset-0 fixed w-full text-black"
	transition:fade
>
	<div class="bg-white p-8" transition:fly={{ y: -100 }}>
		<h2>Share your results on social media:</h2>
		<div id="resultText">
			Today: {new Date().toISOString().slice(0, 10)} // I tried the CNCF #XORQuiz, my results:
			<p class="pb-8">{scorePoints}</p>
		</div>

		<Button on:click={copy}>Copy</Button>
		<Button on:click={() => dispatch('close')}>Close</Button>
	</div>
</div>

<style>
	.modal-wrapper {
		background: rgba(0, 0, 0, 0.7);
		align-items: center;
	}
</style>
