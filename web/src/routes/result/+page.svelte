<script lang="ts">
	import type { ResultItem } from 'src/types/result.type';
	import { onMount } from 'svelte';
	import { get } from 'svelte/store';
	import { resultStore } from '../../lib/results';

	let resultsHistory:Array<ResultItem> = get(resultStore)

	onMount(() => {
		resultStore.update(item => {
			resultsHistory = item
			return resultsHistory
		})

	})
</script>

<div class="">
	<h1 class="text-2xl pb-8">Results</h1>

	{#each resultsHistory as result}
		<li class="list-none">{result.date}: {result.points.join("")}</li>
	{/each}
</div>
