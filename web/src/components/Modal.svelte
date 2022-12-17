<script>
	// Modules
	import { fly, fade } from 'svelte/transition';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();
</script>

<div class="modal-wrapper text-black" transition:fade>
	<div class="modal" transition:fly={{ y: -100 }}>
		<button
			on:click={() => {
				dispatch('close');
			}}>Close</button
		>
		<slot />
	</div>
</div>

<style>
	.modal-wrapper {
		position: fixed;
		top: 0;
		left: 0;
		bottom: 0;
		right: 0;
		background: rgba(0, 0, 0, 0.8);
		display: flex;
		justify-content: center;
		flex-direction: column;
		align-items: center;
	}
	.modal {
		background: white;
		padding: 20px;
		border-radius: 15px;
		width: 275px;
		height: 200px;
	}
</style>
