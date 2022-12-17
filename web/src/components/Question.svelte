<script>
	// @ts-nocheck
	import { score } from '../lib/score';

	export let question;
	export let skip;
	export let index;

	let answer;
	let isCorrect = false;
	let isAnswered = false;
	let NumberOfTry = 0;

	function checkQuestion() {
		if (NumberOfTry === 3) {
			skip();
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
		NumberOfTry++;
	}
</script>

<div>
	<p class="py-2">Question (GitHub) #{index + 1}: {question.GithubDescription}</p>
	<p class="py-2">Question (Crunchbase) #{index + 1}: {question.CrunchbaseDescription}</p>
	<p class="py-2">Question #{index + 1}: {question.Description}</p>

	<p class="py-2">GitHub Stars: {question.GithubStars}</p>
</div>

{#if NumberOfTry > 0 && !isCorrect}
	<p>
		Category {question.Category}
	</p>
{/if}

{#if NumberOfTry > 1 && !isCorrect}
	<p>
		SubCategory {question.Subcategory}
	</p>
{/if}

<form on:submit={checkQuestion}>
	{#if NumberOfTry < 3}
		<input bind:value={answer} class="text-black" />
		{#if !isCorrect}
			<button type="submit"> Submit / </button>
			<button on:click={hint}> Hint </button>
		{/if}
	{/if}
	{#if NumberOfTry === 3}
		<button on:click={skip}>Next</button>
	{/if}
</form>

{#if isAnswered}
	<h5 class:isCorrect>
		{#if isCorrect}
			Correct answer!
		{:else if NumberOfTry === 3}
			You are wrong, go on the next question. The correct answer: {question.Name}
		{:else}
			Wrong! Try harder!
		{/if}
	</h5>
{/if}
<!-- {#each allAnswers as answer}
	<button class="answer" disabled={isAnswered} on:click={() => checkQuestion(answer.correct)}
		>{@html answer.answer}</button
	>
{/each} -->
