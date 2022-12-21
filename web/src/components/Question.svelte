<script lang="ts">
	// @ts-nocheck
	import { score, hintNumber, guessNumber } from '../lib/store';
	import Button from './Button.svelte';
	import FuzzySet from 'fuzzyset.js';
	import type { QuizItem } from 'src/types/quiz.type';
	import { onDestroy, onMount } from 'svelte';

	export let question: QuizItem;
	export let nextQuestion;
	export let index;

	let answer = '';
	let isCorrect = false;
	let isAnswered = false;
	let nameHint: string;

	function skip() {
		$score.splice(index, 1, '🔴');
	}

	function hint() {
		$hintNumber++;

		nameHint = question.Name.slice(0, 1);

		if ($hintNumber > 1) {
			if (question.Name.length > 10) {
				nameHint = question.Name.slice(0, 4);
			} else {
				nameHint = question.Name.slice(0, 2);
			}
		}
	}

	// Reset the hint and guess store
	onDestroy(() => {
		$guessNumber = 0;
		$hintNumber = 0;
	});

	function checkQuestion() {
		$guessNumber++;

		let fuzzy = new FuzzySet([question.Name], true);
		let res = fuzzy.get(answer);

		// maybe 0.87 is enough
		if (res && res[0][0] > 0.87) {
			isCorrect = true;

			switch ($hintNumber) {
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
			if ($guessNumber >= 3) {
				$score.splice(index, 1, '🔴');
				nextQuestion();
			}
		}
		isAnswered = true;
	}
</script>

<div>
	<!-- Dashboard -->
	<span class="text-amber-300">Question:</span>
	{index + 1}/5
	<span class="text-amber-300">Attempts:</span>
	{$guessNumber}/3
	<span class="text-amber-300">Hints:</span>
	{$hintNumber}/2

	<!-- Informations -->
	<p><span>Logo:</span></p>
	<img
		class:blur={$hintNumber > 0}
		class:blur-sm={$hintNumber > 1}
		class="bg-white w-2/12 blur-md mt-4 mb-8"
		src={question.Logo}
		draggable="false"
		alt="Project Logo"
	/>
	<p><span>GitHub:</span> {question.GithubDescription}</p>
	<p><span>Crunchbase:</span> {question.CrunchbaseDescription}</p>
	<p>
		<span>GitHub Stars:</span>
		{question.GithubStars} / <span>GitHub Contributors:</span>
		{question.GithubContributorsCount}
	</p>

	{#if $hintNumber > 0}
		<p>
			<span>Category:</span>
			{question.Category} / <span>SubCategory:</span>
			{question.Subcategory}
		</p>
	{/if}

	{#if $hintNumber > 0}
		<p><span>Project Name:</span> {nameHint}...</p>
	{/if}
</div>

<form>
	<input required type="text" bind:value={answer} class="text-black p-2 pl-2 w-1/2" />
	{#if index < 4}
		<Button on:click={checkQuestion}>Submit</Button>
		<Button on:click={nextQuestion} on:click={skip}>Skip</Button>
	{:else}
		{#if $guessNumber === 2}
			<Button on:click={checkQuestion}>Finish</Button>
		{:else}
			<Button on:click={checkQuestion}>Submit</Button>
		{/if}
		<Button on:click={skip} on:click={nextQuestion}>Skip</Button>
	{/if}
	<!-- Hint button logic -->
	{#if $hintNumber < 2}
		<Button on:click={hint}>Hint</Button>
	{/if}
</form>

{#if isAnswered}
	{#if !isCorrect}
		<div class="pt-8 animate-bounce">
			<span class="text-yellow-600">Wrong! Try harder!</span>
		</div>
	{/if}
{/if}

<div class="py-4">
	Your result: {$score.join('')}
</div>

<style>
	p {
		padding-bottom: 1rem; /* 8px */
	}
	p > span {
		font-weight: 700;
		color: rgb(59 130 246);
	}

	.blur {
		filter: blur(8px);
	}

	.blur-sm {
		filter: blur(4px);
	}
</style>
