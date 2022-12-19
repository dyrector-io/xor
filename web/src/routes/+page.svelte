<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown';
	const homeContent = `
Hello! 👋
<br />

We're Levente and Nandor. As we were brainstorming to find our company's value proposition, we checked out some CNCF projects. Levente started asking questions to see if others are able to guess. This where the fun began... So we decided to convert it into a super simple quiz!
<br />
Every day you get the value propositions of 5 CNCF projects and you have to guess what the projects are! You have three attempts per project, and can get two hints should you need it.
<br />
Depending how many hints you use, you <TBD what happens here?>:
🟢 - You got the project right without any hints.
🟡 - You got the project right using 1 hint.
🟠 - You got the project right using 2 hints.
🔴 - You didn't get the project right.
⚪ - Unanswered questions.
<br />
Your data is stored client-side (browser local storage). More details [here](https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage). The quiz is reset every noon (UTC).
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
		<p class="mt-4">contact: <a class="text-emerald-300" href="mailto:hello@dyrector.io" target="_blank">hello@dyrector.io</a> / repository:  <a class="text-emerald-300"  href="https://github.com/dyrector-io/xor/">dyrector-io/xor</a></p>
	</div>
</div>
