<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown';
	const homeContent = `
👋 Hello,
<br />
Every day you get the short description of 5 npm packages and you have to guess what the projects are! You have three attempts per project.
<br />
Depending how many hints you use:
🟢 - You got the project right without any hints.
🟡 - You got the project right, but not on the first try.
🔴 - You didn't get the project right.
⚪ - Unanswered questions.
`;

	const information = `
Your data is stored client-side (browser local storage). More details [here](https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage).
The quiz is reset every noon (UTC) quiz starts on February 3 and runs for 15 days, ending on February 18.
<br />
Have Fun and Good Luck! 🍾
`;
</script>

<div class="">
	<h1 class="text-2xl pb-8">Home</h1>
	<div>
		<SvelteMarkdown
			source={homeContent}
			options={{
				breaks: true
			}}
		/>

		<a href="/quiz"
			><button class="bg-emerald-300 hover:bg-emerald-700 text-zinc-800 font-bold p-2 my-8"
				>Start</button
			></a
		>
		<SvelteMarkdown
			source={information}
			options={{
				breaks: true
			}}
		/>

		<p class="mt-4">
			contact: <a class="text-emerald-300" href="mailto:hello@dyrector.io" target="_blank"
				>hello@dyrector.io</a
			>
			/ repository:
			<a class="text-emerald-300" href="https://github.com/dyrector-io/xor/" target="_blank"
				>dyrector-io/xor</a
			>
		</p>
	</div>
</div>
