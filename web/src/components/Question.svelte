<script>
	// @ts-nocheck
	import { score } from '../lib/score';
	import Button from './Button.svelte';

	export let question;
	export let nextQuestion;
	export let index;

	let answer;
	let isCorrect = false;
	let isAnswered = false;
	let NumberOfTry = 0;
	let hintNumber = 0;

	function checkQuestion() {
		if (NumberOfTry === 3) {
			nextQuestion();
		}
		NumberOfTry++;

		if (answer === question.Name) {
			isCorrect = true;
			score.update((currentVal) => currentVal + 1);
		}

		isAnswered = true;
	}

	function hint() {
		// TODO Implement
		hintNumber++;
	}
</script>

<div>
	<h2 class="pb-2">Question #{index + 1}: {question.Name}</h2>
	<span class="text-amber-300">Attempt:</span> {NumberOfTry}/3 <span class="text-amber-300">Hints:</span> {hintNumber}/2
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
		{#if NumberOfTry < 3}
		<input bind:value={answer} class="text-black p-2 pl-2 w-1/2" />
		{#if !isCorrect}
			<Button on:click={checkQuestion}>Submit</Button>
			{#if index < 4}
				<Button on:click={nextQuestion}>Skip</Button>
			{/if}
			{#if hintNumber < 2}
				<Button on:click={hint}>Hint</Button>
			{/if}
		{/if}
	{/if}
	{#if (index < 4 && NumberOfTry === 3) || (index < 4 && isCorrect)}
		<Button on:click={nextQuestion}>Next</Button>
	{/if}
	{#if (index === 4 && NumberOfTry === 3) || (index === 4 && isCorrect)}
		<Button on:click={nextQuestion}>Finish</Button>
	{/if}
</form>

{#if isAnswered}
	<div class="pt-8 animate-bounce">
		{#if isCorrect}
			<span class="text-emerald-400">Correct answer!</span>
		{:else if NumberOfTry === 3}
			<span class="text-red-600">You missed! Answer: {question.Name}</span>
		{:else}
			<span class="text-yellow-600">Wrong! Try harder!</span>
		{/if}
	</div>
{/if}

<!-- {#each allAnswers as answer}
	<button class="answer" disabled={isAnswered} on:click={() => checkQuestion(answer.correct)}
		>{@html answer.answer}</button
	>
{/each} -->
<style>
	p {
		padding-bottom: 1rem; /* 8px */
	}
	p > span {
		font-weight: 700;
		color: rgb(59 130 246);
	}
</style>
