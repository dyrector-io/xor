<script lang="ts">
	import { fly, fade } from 'svelte/transition';
	import { createEventDispatcher } from 'svelte';
	import Button from '../components/Button.svelte';
	import type { ResultItem } from 'src/types/result.type';

	const dispatch = createEventDispatcher();
	export let todayResult: ResultItem | undefined;

	const copy = () => {
		const text = document.getElementById('resultText')!.innerText;
		navigator.clipboard.writeText(text);
	};
</script>

<div
	class="modal-wrapper flex flex-col  justify-center backdrop-blur-sm bg-gray-800 inset-0 fixed w-full text-black"
	transition:fade
>
	<div class="bg-white p-8 md:w-8/12" transition:fly={{ y: -100 }}>
		<h2>
			You have proven your familiarity with the JavaScript ecosystem, you may flex to others with
			your score. Share on socials:
		</h2>
		<div class="my-8" id="resultText">
			{todayResult?.date} JS #XORQuiz results:
			{todayResult?.points.join('')}<br />
			https://xor.dyrectorio.com
		</div>

		<Button on:click={copy}
			><a
				class="twitter-share-button"
				href="https://twitter.com/intent/tweet?text={todayResult?.date}&nbsp;JS&nbsp;ecosystem&nbsp;XOR&nbsp;Quiz&nbsp;results:&nbsp;{todayResult?.points.join(
					''
				)}&hashtags=XORQuiz,JavaScript,Quiz&url=https://xor.dyrectorio.com"
				data-size="large"
			>
				Twitter</a
			></Button
		>
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
