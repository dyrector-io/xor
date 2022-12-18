import { env } from '$env/dynamic/private';

export type Question = {
	Name: string;
	Category: string;
	Subcategory: string;
	GithubStars: number;
	GithubDescription: string;
};

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	const data = await fetch(env.API_PATH)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));

	return { result: data };
}
