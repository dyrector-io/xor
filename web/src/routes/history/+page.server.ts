import { env } from '$env/dynamic/private';

/** @type {import('./$types').PageServerLoad} */
export async function load() {
	const data = await fetch(`${env.API_PATH}/history`)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));

	return { result: data };
}
