import { error } from '@sveltejs/kit';

export type Question = {
	Name: string;
	Category: string;
	Subcategory: string;
	GithubStars: number;
	GithubDescription: string;
};

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	const data = await fetch(`http://localhost:3333/quiz`)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));

	return { result: data };
}
