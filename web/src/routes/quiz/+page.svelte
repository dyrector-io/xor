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
	import { PUBLIC_API_PATH } from '$env/static/public';

	const HTTP_STATUS = {
		OK: 200,
		GONE: 410
	};

	let quiz: QuizResponse = {
		Date: '',
		List: new Array<JavascriptQuizItem>()
	};

	let endOfTheQuiz = false;
	let filledForToday = false;
	let closed = false;
	let gone = true;
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
		quiz = await fetch(`${PUBLIC_API_PATH}/quiz`)
			.then((resp) => {
				if (resp.status === HTTP_STATUS.GONE) {
					gone = true;
					return { Date: '', List: [] };
				} else {
					return resp.json();
				}
			})
			.catch((err) => console.log(err.message));

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

{#if gone}
	<div class="py-4">
		The quiz is ended. Please come back later or visit the <a
			class="text-emerald-300"
			href="https://github.com/dyrector-io/xor/"
			target="_blank">repository</a
		> for more information.
	</div>
{:else if filledForToday}
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
