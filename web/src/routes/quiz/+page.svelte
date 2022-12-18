<script lang="ts">
	import { onMount, onDestroy } from 'svelte';

	// Store
	import { score } from '../../lib/score';
	import { results } from '../../lib/results';

	// Components
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';

	export let data: any;
	const dailyQuestions = data.result;
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
		const filledForToday = $results.find(
			(x) => x['date'] === new Date().toISOString().slice(0, 10)
		);

		if (filledForToday) {
			todayDone = true;
		}
	});

	$: if (activeQuestion === 5) {
		endOfTheQuiz = true;
		scorePoints = $score.join('');
	}

	$: if (endOfTheQuiz) {
		const todayDate = new Date().toISOString().slice(0, 10);
		$results.push({ date: todayDate, points: $score });
	}

	// TODO REMOVE LOCALSTORAGE
	// $: if (endOfTheQuiz) {
	// 	const todayDate = new Date().toISOString().slice(0, 10);
	// 	let temporary: [] = $results;
	// 	temporary.push({ date: todayDate, points: $score });
	// 	results.set(temporary);
	// }
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
