<script lang="ts">
	import { onMount } from 'svelte';
	import { score } from '../../lib/score';
	import { results } from '../../lib/results';
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';

	let activeQuestion = 0;

	// Recieve questions from server Page
	export let data;
	const dailyQuestions = data.result;
	let endOfTheQuiz= false;
	let todayDone = false;

	function skip() {
		// if (questionNumber === 10) {
		//   $score < 7 ? (isLooseModalOpen = true) : (isWinModalOpen = true);
		// } else {
		//   activeQuestion = activeQuestion + 1;
		// }
		activeQuestion = activeQuestion + 1;
	}

	function closeModal() {
		endOfTheQuiz = false;
	}

	onMount(async () => {
		const filledForToday = $results.find((x) => x.date === new Date().toISOString().slice(0, 10));

		if (filledForToday) {
			todayDone = true
		}
	});

	$: if (activeQuestion === 5) {
		endOfTheQuiz = true;
	}

	$: if (endOfTheQuiz) {
		const todayDate = new Date().toISOString().slice(0, 10);

		let temporary: [] = $results;
		temporary.push({ date: todayDate, points: $score });
		results.set(temporary);
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
			<Question {skip} {question} {index} />
		{/if}
	{/each}
</div>

<div class="py-4">
	Your result: {$score}
</div>
{/if}


{#if endOfTheQuiz}
	<Modal on:close={closeModal}>
		<h2>You Lost!</h2>
		<p>Incididunt eiusmod culpa nisi voluptate in.</p>
	</Modal>
{/if}
