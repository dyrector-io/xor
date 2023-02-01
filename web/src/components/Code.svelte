<script lang="ts">
	export let raw = '';
	$: codeContent = raw.replace(/`/g, '');
</script>

<pre><code class="code">{codeContent}</code></pre>

<style>
	.code {
		display: flex;
		margin-top: 10px;
		padding: 20px;
		background-color: #2a2a2a;
		color: rgb(98, 112, 114);
	}
</style>
