<script>
	// Modules
	import { fly, fade } from 'svelte/transition';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
	import Button from '../components/Button.svelte';
</script>

<div class="modal-wrapper flex flex-col  justify-center backdrop-blur-sm bg-gray-800 inset-0 fixed w-full text-black" transition:fade>
	<div class="bg-white p-8" transition:fly={{ y: -100 }}>
		
		<slot />
		<Button
			on:click={() => {
				dispatch('close');
			}}>Close</Button>
	</div>
</div>

<style>
	.modal-wrapper {
		background: rgba(0, 0, 0, 0.7);
		align-items: center;
	}
</style>
