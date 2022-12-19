import { writable } from 'svelte/store';

export const score = writable(['⚪', '⚪', '⚪', '⚪', '⚪']);

export const questionNumber = writable(0);
