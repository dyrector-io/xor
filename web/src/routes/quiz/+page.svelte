<script lang="ts">
	import { onMount } from 'svelte';

	// Store
	import { score, questionNumber } from '../../lib/store';
	import { resultStore } from '../../lib/results';
	import { get } from 'svelte/store';

	// Components
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';
	import type { JavascriptQuizItem, QuizResponse } from 'src/types/quiz.type';
	import type { ResultItem } from 'src/types/result.type';

	import quizJson from './javascript-test.json';

	let quiz: QuizResponse = {
		Date: '',
		List: new Array<JavascriptQuizItem>()
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
			closed = false;

			resultStore.set([...$resultStore, { date: quiz.Date, points: $score }]);
		}
	}

	function closeModal() {
		endOfTheQuiz = false;
		closed = true;
	}

	onMount(async () => {
		// quiz = await fetch(`${PUBLIC_API_PATH}/quiz`)
		// 	.then((resp) => resp.json())
		// 	.catch((err) => console.log(err.message));

		quiz = quizJson;

		todayResult = $resultStore.find((e) => e['date'] === quiz.Date);

		if (todayResult) {
			if (todayResult?.date === quiz.Date) {
				filledForToday = true;
			} else {
				filledForToday = false;
			}
		} else {
			if ($resultStore.length === 0) {
				filledForToday = false;
			}
		}
	});

	$: todayResult = $resultStore.find((e) => e['date'] === quiz.Date);
</script>

<h1 class="text-2xl pb-8">XOR Quiz</h1>

{#if filledForToday}
	<div class="py-4">
		You are done with the today quiz. Check your results <a class="text-emerald-300" href="/result"
			>here</a
		>.

		<h2 class="my-4">Today solutions:</h2>
		{#each quiz.List as question, index}
			<li class="ml-2 list-none">#{index + 1} {question.Name} <br /></li>
		{/each}
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

{#if quiz.Date && filledForToday && !closed && todayResult}
	<Modal on:close={closeModal} {todayResult} />
{/if}
