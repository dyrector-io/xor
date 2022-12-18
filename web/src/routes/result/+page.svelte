<script lang="ts">
	import { results } from '../../lib/results';

	const resultsHistory = $results;
</script>

<div class="">
	<h1 class="text-2xl pb-8">Results</h1>

	{#each resultsHistory as result, i}
		<li class="list-none">{JSON.stringify(result.date)}: {JSON.stringify(result.points)}</li>
	{/each}
</div>
