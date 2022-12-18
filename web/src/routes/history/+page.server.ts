import { env } from '$env/dynamic/private';
import type { HistoryItem, HistoryResponse } from 'src/types/history.type';

/** @type {import('./$types').PageServerLoad} */
export async function load(): Promise<HistoryResponse> {
	const data: Array<HistoryItem> = await fetch(`${env.API_PATH}/history`)
		.then((resp) => resp.json())
		.catch((err) => console.log(err.message));

	return { result: data };
}
