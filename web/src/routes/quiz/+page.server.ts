import { env } from '$env/dynamic/private';
import type { QuizItem, QuizResponse } from 'src/types/quiz.type';

/** @type {import('./$types').PageServerLoad} */
export async function load(): Promise<QuizResponse> {
	const data: Array<QuizItem> = await fetch(`${env.API_PATH}/quiz`)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));

	return { result: data };
}
