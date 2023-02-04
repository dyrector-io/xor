<script lang="ts">
	// @ts-nocheck
	import { score, hintNumber, guessNumber } from '../lib/store';
	import Button from './Button.svelte';
	import FuzzySet from 'fuzzyset.js';
	import type { JavascriptQuizItem } from 'src/types/quiz.type';
	import { onDestroy } from 'svelte';
	import SvelteMarkdown from 'svelte-markdown';
	import Code from './Code.svelte';

	export let nextQuestion;
	export let question: JavascriptQuizItem;
	export let index;

	let answer = '';
	let isCorrect = false;
	let isAnswered = false;

	function skip() {
		$score.splice(index, 1, '🔴');
	}

	// Reset the hint and guess store
	onDestroy(() => {
		$guessNumber = 0;
		$hintNumber = 0;
	});

	function checkQuestion() {
		$guessNumber++;

		let fuzzy = new FuzzySet([question.Name], true);
		const jsEndPattern = /.?js$/i;
		if (question.Name.match(jsEndPattern)) {
			fuzzy.add(question.Name.replace(jsEndPattern, ''));
		}

		let res = fuzzy.get(answer);
		// maybe 0.87 is enough
		if (res && res[0][0] > 0.87) {
			isCorrect = true;

			if ($guessNumber == 1) {
				$score.splice(index, 1, '🟢');
			} else if ($guessNumber > 1) {
				$score.splice(index, 1, '🟡');
			}

			nextQuestion();
		} else {
			if ($guessNumber == 3) {
				$score.splice(index, 1, '🔴');
				nextQuestion();
			}
		}

		isAnswered = true;
	}

	function onKeyDown(e) {
		if (e.key === 'Enter') {
			console.log('Enter pressed');
			checkQuestion();
		}
	}
</script>

<div>
	<div class="mb-8">
		<!-- Dashboard -->
		<span class="text-amber-300">Question:</span>
		{index + 1}/5
		<span class="text-amber-300">Attempts:</span>
		{$guessNumber}/3
	</div>

	<!-- Informations -->
	<p><span>Description:</span> {question.Description}</p>
	<p><span>Random fact:</span> {question.RandomFact}</p>
	<p>
		<span>Code example:</span><br />
		<SvelteMarkdown
			renderers={{ codespan: Code, code: Code }}
			source={question.CodeExample}
			options={{
				breaks: true,
				sanitize: false,
				langPrefix: 'language-js'
			}}
		/>
	</p>
	<p>
		<span>GitHub Stars:</span>
		{question.GithubStars.toLocaleString('en-US')} / <span>Weekly npm downloads:</span>
		{question.WeeklyDownloads.toLocaleString('en-US')}
	</p>
</div>

<form>
	<input
		on:keydown={onKeyDown}
		required
		autofocus
		type="text"
		bind:value={answer}
		class="text-black p-2 pl-2 w-1/2"
	/>
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
</style>
