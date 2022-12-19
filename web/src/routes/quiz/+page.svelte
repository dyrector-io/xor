<script lang="ts">
	import { onMount } from 'svelte';
	import { PUBLIC_API_PATH } from '$env/static/public';

	// Store
	import { score, questionNumber } from '../../lib/store';
	import { resultStore } from '../../lib/results';
	import { get } from 'svelte/store';

	// Components
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';
	import type { QuizResponse } from 'src/types/quiz.type';

	let quiz: QuizResponse
	let endOfTheQuiz = false;
	let todayDone = false;

	function nextQuestion() {
		$questionNumber++;

		if ($questionNumber === 5) {
			endOfTheQuiz = true;
			todayDone = true;

			
			resultStore.set([...$resultStore, { date: quiz.Date, points: $score }]);
		}
	}

	function closeModal() {
		endOfTheQuiz = false;
	}

	onMount(async () => {
		quiz = await fetch(`${PUBLIC_API_PATH}/quiz`)
			.then((resp) => resp.json())
			.catch((err) => console.log(err.message));

		const res = get(resultStore);
		const filledForToday = res.find((e) => e['date'] === new Date().toISOString().slice(0, 10));

		if (filledForToday) {
			todayDone = true;
		}
	});
</script>

<h1 class="text-2xl pb-8">XOR Quiz</h1>
{#if todayDone}
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

{#if endOfTheQuiz}
	<Modal on:close={closeModal} />
{/if}
