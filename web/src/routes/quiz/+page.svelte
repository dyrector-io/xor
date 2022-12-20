<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_API_PATH } from '$env/static/public';

	// Store
	import { score, questionNumber } from '../../lib/store';
	import { resultStore } from '../../lib/results';

	// Components
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';
	import type { QuizItem, QuizResponse } from 'src/types/quiz.type';
	import type { ResultItem } from 'src/types/result.type';

	let quiz: QuizResponse = {
		Date: '',
		List: new Array<QuizItem>()
	};
	let endOfTheQuiz = false;
	let filledForToday = false;
	let closed = false;
	let todayResult: ResultItem | undefined;

	function nextQuestion() {
		$questionNumber++;

		if ($questionNumber === 5) {
			endOfTheQuiz = true;
			filledForToday = true;

			$resultStore = [...$resultStore, { date: quiz.Date, points: $score }]
		}
	}

	function closeModal() {
		endOfTheQuiz = false;
		closed = true;
	}

	onMount(async () => {
		quiz = await fetch(`${PUBLIC_API_PATH}/quiz`)
			.then((resp) => resp.json())
			.catch((err) => console.log(err.message));

		todayResult = $resultStore.find((e) => e['date'])
		filledForToday = todayResult?.date === quiz.Date;
	});

</script>

<h1 class="text-2xl pb-8">XOR Quiz</h1>

{#if filledForToday}
	<div class="py-4">
		You are done with the today quiz. Check your results <a href="/result">here.</a>
	</div>
{:else}
	<div>
		{#each quiz.List as question, index}
			{#if index === $questionNumber}
				<Question {nextQuestion} {question} {index} />
			{/if}
		{/each}
	</div>
{/if}

{#if quiz.Date && filledForToday && !closed}
	<Modal on:close={closeModal} todayResult={todayResult} />
{/if}
