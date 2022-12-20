<script lang="ts">
	import { fly, fade } from 'svelte/transition';
	import { createEventDispatcher } from 'svelte';
	import { score } from '$lib/store';
	import { resultStore } from '../lib/results';
	import Button from '../components/Button.svelte';

	const dispatch = createEventDispatcher();
	export let todayResult;

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
		<h2>Proven your familiarity with the CNCF landscape, flex to others with your score.</h2>
		<div class="my-8" id="resultText">
			Today: {todayResult.date} // I tried the CNCF #XORQuiz, my results:
			{todayResult.points.join('')}
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
