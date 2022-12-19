<script lang="ts">
	import SvelteMarkdown from 'svelte-markdown'

	const homeContent = `
Hello! 👋
<br />

We're Levente and Nandor,  as we were brainstorming about our company's value proposition, we checked out some CNCF projects.  Levente started asking questions to see if others are able to guess. 
Where the fun began... so we decided to convert it into a super simple online quiz.
<br />
Every day you get the value propositions of 5 projects. You have 3 attempts per project, and 2 hints for each of them.
  
  
You can get 4 types of results depending on how many hints you use:
🟢 - You got the project right without any hints.
🟡 - You got the project right using 1 hint.
🟠 - You got the project right using 2 hints.
🔴 - You didn't get the project right.
Your data is stored client-side (browser local storage). More details [here](https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage).
If you'd like us to create similar games, drop us an email at [hello@dyrector.io](mailto:hello@dyrector.io), or join us on [Discord](https://discord.gg/pZWbd4fxga).
<br />
**Copy:**
Proven your familiarity with the CNCF landscape, flex to others with your score.
**History:**
Take a look at previously included projects.`
</script>

<div class="">
	<h1 class="text-2xl pb-8">Home</h1>
	<div>
		<SvelteMarkdown source={homeContent} options={{
			breaks: true,
		}} />
	</div>
</div>
