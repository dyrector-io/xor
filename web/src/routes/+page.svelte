<script lang="ts">
	import { onMount } from 'svelte';
	const END_POINT = 'URL';

    let obj: Question
    let answer: string

    type Question = {
        question: string,
        answer: string,
        starsNumber: number,
        field: string
    }

	onMount(() => {
        getQuiz()
	});

	function getQuiz(): Question {
        const q: Question = {
            question: "Kiscica?",
            answer: "Aladar",
            starsNumber: 2121,
            field: "Service Meash"
        }

        obj = q

        return q
		// return fetch(`URL`)
		// 	.then((resp) => resp.json())
		// 	.then((json) => init(json))
		// 	.catch((err) => console.log(err.message))
	}


	function checkAnswer(): any {
	}

    function init() {
        const obj: Question = getQuiz()
        console.log(obj)
    }

</script>

<div class="">
    <h1 class="text-3xl font-bold">XOR Quiz</h1>
    <div>{JSON.stringify(obj)}</div>
    <input class="text-black" bind:value={answer}>
</div>
