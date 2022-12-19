<script lang="ts">
	import type { HistoryItem } from 'src/types/history.type';
	import { onMount } from 'svelte';
	import { PUBLIC_API_PATH } from '$env/static/public';

	let history: Array<HistoryItem> = []

	onMount(async () => {
		history = await fetch(`${PUBLIC_API_PATH}/history`)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));
	});
</script>

<div class="">
	<h1 class="text-2xl pb-8">History</h1>
	<p>Take a look at previously included projects.</p>
	<div>
		{#each history as item}
			<li class="list-none">{item.Date}: {item.Projects}</li>
		{/each}
	</div>
</div>
