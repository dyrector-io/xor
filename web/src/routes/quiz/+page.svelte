<script lang="ts">
	// Data
	import { score } from '../../lib/score';
	import Question from '../../components/Question.svelte';
	import Modal from '../../components/Modal.svelte';

	let activeQuestion = 0;

	// Recieve questions from server Page
	export let data;
	const dailyQuestions = data.result;
	let endOfTheQuiz: boolean = false;

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

	$: if (activeQuestion === 5) {
		endOfTheQuiz = true;
	}
</script>


	<h1 class="text-2xl pb-8">XOR Quiz</h1>
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


{#if endOfTheQuiz}
	<Modal on:close={closeModal}>
		<h2>You Lost!</h2>
		<p>Incididunt eiusmod culpa nisi voluptate in.</p>
	</Modal>
{/if}
