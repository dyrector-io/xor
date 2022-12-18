<script lang="ts">
	// @ts-nocheck
	import { score } from '../lib/score';
	import Button from './Button.svelte';
	import FuzzySet from 'fuzzyset.js'

	export let question;
	export let nextQuestion;
	export let index;

	let answer;
	let isCorrect = false;
	let isAnswered = false;
	let guessNumber = 0;
	let hintNumber = 0;

	function skip() {
		$score.splice(index, 1, '🔴');
	}


	function checkQuestion() {
		guessNumber++;

		let fuzzy = new FuzzySet([question.Name], true)
		let res = fuzzy.get(answer)
		if (res) {
			console.log(res[0][0])
		}
		// maybe 0.87 is enough
		if (res && res[0][0] > 0.9) {
	
			isCorrect = true;

			switch (hintNumber) {
				case 0:
				$score.splice(index, 1, '🟢');
					break;
				case 1:
				$score.splice(index, 1, '🟡');
					break;
				case 2:
				$score.splice(index, 1, '🟠');
					break;
			}

			nextQuestion();
		} else {
			if (guessNumber === 3) {
				$score.splice(index, 1, '🔴');
				nextQuestion();
			}
		}

		isAnswered = true;
	}

	function hint() {
		hintNumber++;
	}
</script>

<div>
	<h2 class="pb-2">Question #{index + 1}: {question.Name}</h2>
	<span class="text-amber-300">Attempt:</span>
	{guessNumber}/3 <span class="text-amber-300">Hints:</span>
	{hintNumber}/2
	<p><span>Logo:</span></p>
	<img class="w-2/12 blur py-4" src={question.Logo} draggable="false" />

	<p><span>GitHub:</span> {question.GithubDescription}</p>
	<p><span>Crunchbase:</span> {question.CrunchbaseDescription}</p>
	<p><span>GitHub Stars:</span> {question.GithubStars}</p>

	{#if hintNumber > 0 && !isCorrect}
		<p><span>Category:</span> {question.Category}</p>
	{/if}

	{#if hintNumber > 1 && !isCorrect}
		<p><span>SubCategory:</span> {question.Subcategory}</p>
	{/if}
</div>

<form>
	{#if guessNumber < 3}
		<input bind:value={answer} class="text-black p-2 pl-2 w-1/2" />
		{#if !isCorrect}
			{#if index < 4}
				<Button on:click={checkQuestion}>Submit</Button>
				<Button on:click={nextQuestion} on:click={skip}>Skip</Button>
			{:else}
				<Button on:click={checkQuestion}>Finish</Button>
			{/if}
			{#if hintNumber < 2}
				<Button on:click={hint}>Hint</Button>
			{/if}
		{/if}
	{/if}
	{#if (index < 4 && guessNumber === 3) || (index < 4 && isCorrect)}
		<Button on:click={nextQuestion}>Next</Button>
	{/if}
	{#if (index === 4 && guessNumber === 3) || (index === 4 && isCorrect)}
		<Button on:click={nextQuestion}>Finish</Button>
	{/if}
</form>

<div class="py-4">
	Your result: {$score.join('')}
</div>

{#if isAnswered}
	<div class="pt-8 animate-bounce">
		{#if !isCorrect}
			<span class="text-yellow-600">Wrong! Try harder!</span>
		{:else if guessNumber === 3}
			<span class="text-red-600">You missed! Answer: {question.Name}</span>
		{/if}
	</div>
{/if}

{#if isAnswered}
	<div class="pt-8 animate-bounce">
		{#if !isCorrect}
			<span class="text-yellow-600">Wrong! Try harder!</span>
		{:else if NumberOfTry === 3}
			<span class="text-red-600">You missed! Answer: {question.Name}</span>
		{/if}
	</div>
{/if}

<style>
	p {
		padding-bottom: 1rem; /* 8px */
	}
	p > span {
		font-weight: 700;
		color: rgb(59 130 246);
	}
</style>
