<script>
	import { page } from '$app/stores';
	export let href;
	import '../app.css';
	export const prerender = true;
</script>

<div class="flex flex-col  h-screen">
	<header>
		<nav class="p-8 lowercase">
			<a class:active={$page.url.pathname === '/'} href="/">Home</a> /
			<a class:active={$page.url.pathname === '/quiz'} href="/quiz">Quiz</a> /
			<a class:active={$page.url.pathname === '/result'} href="/result">Results</a> /
			<a class:active={$page.url.pathname === '/history'} href="/history">History</a>
		</nav>
	</header>

	<main>
		<slot />
	</main>

	<footer class="text-center text-sm mt-auto h-8">
		deployed & hosted using<a class="text-emerald-300" href="https://dyrector.io" target="_blank">dyrector.io</a>
	</footer>
</div>

<style>
	:global(body) {
		font-family: 'Roboto Mono', monospace;
		color: #fefefe;
		background-color: rgb(39 39 42);
		margin: 0 auto;
		width: 85%;
		align-content: center;
	}

	.active {
		color: rgb(110 231 183);
	}
</style>
