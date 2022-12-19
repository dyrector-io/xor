<script>
	import '../app.css';
	export const prerender = true
</script>

<div class="flex flex-col  h-screen">
	<header>
		<nav class="p-8 lowercase">
			<a href="/">Home /</a>
			<a href="/quiz">Quiz /</a>
			<a href="/result">Results /</a>
			<a href="/history">History</a>
		</nav>
	</header>

	<main>
		<slot />
	</main>

	<footer class="text-center text-sm mt-auto h-8">
		deployed & hosted using <a class="text-emerald-300" href="https://dyrector.io" target="_blank"
			>dyrector.io</a>
	</footer>
</div>

<style>
	:global(body) {
		font-family: 'Roboto Mono', monospace;
		color: #fefefe;
		background-color: rgb(39 39 42);
		margin: 0 auto;
		width: 75%;
		align-content: center;
	}
</style>
