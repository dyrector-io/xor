<script lang="ts">
	import type { HistoryResponse } from 'src/types/history.type';
	export let data: HistoryResponse;
</script>

<div class="">
	<h1 class="text-2xl pb-8">History</h1>
	<div>
		{#each data.result as item}
			<li class="list-none">{item.Date}: {item.Projects}</li>
		{/each}
	</div>
</div>
