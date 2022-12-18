<script>
	import '../app.css';
</script>

<nav class="p-8 lowercase">
	<a href="/">Home /</a>
	<a href="/quiz">Quiz /</a>
	<a href="/result">Results /</a>
	<a href="/history">History</a>
</nav>

<div>
	<slot/>
</div>


<!-- <footer class="text-center text-sm absolute inset-x-0 bottom-0 h-8">
	Created by <a href='https://github.com/dyrector-io/xor'>dyrector.io</a> team with love. If you wanna submit extra question please visit the repository.
</footer> -->

<style>
	:global(body) {
		font-family: 'Roboto Mono', monospace;
		color: #fefefe;
		background: rgb(12, 23, 19);
        margin: 0 auto;
		 width: 75%;
		align-content: center;
	}
</style>
