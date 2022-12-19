<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { PUBLIC_API_PATH } from '$env/static/public';


	export const load: PageServerLoad = () => {
		console.log(PUBLIC_API_PATH) // public
	}
	// Store
	import { score } from '../../lib/score';
	import { resultStore } from '../../lib/results';
	import { get } from 'svelte/store'

	// Components
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';
	import type { QuizItem } from 'src/types/quiz.type';
	import type { PageServerLoad } from './$types';


	let dailyQuestions: Array<QuizItem> = []
	let activeQuestion = 0;
	let endOfTheQuiz = false;
	let todayDone = false;
	let scorePoints: string = "";

	function nextQuestion() {
		activeQuestion++;
	}

	function closeModal() {
		endOfTheQuiz = false;
	}

	onDestroy(() => {
		$score.splice(0,5)
	});

	onMount(async () => {
		console.log("env:", PUBLIC_API_PATH)
		dailyQuestions = await fetch(`${PUBLIC_API_PATH}/quiz`)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));

		const res = get(resultStore)
		const filledForToday = res.find(
			(x) => x['date'] === new Date().toISOString().slice(0, 10)
		);

		if (filledForToday) {
			todayDone = true;
		}
	});

	$: if (activeQuestion === 5) {
		endOfTheQuiz = true;
		todayDone = true;
		scorePoints = $score.join('');

		const todayDate = new Date().toISOString().slice(0, 10);

		const earlierResults = get(resultStore)
		resultStore.set([...earlierResults, { date: todayDate, points: $score }])
	}
</script>

<h1 class="text-2xl pb-8">XOR Quiz</h1>
{#if todayDone}
	<div class="py-4">
		You are done with the today quiz. Check your results <a href="/result">here.</a>
	</div>
{:else}
	<div>
		{#each dailyQuestions as question, index}
			{#if index === activeQuestion}
				<Question {nextQuestion} {question} {index} />
			{/if}
		{/each}
	</div>
{/if}

{#if endOfTheQuiz}
	<Modal on:close={closeModal} scorePoints={scorePoints}></Modal>
{/if}
