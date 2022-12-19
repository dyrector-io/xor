import { writable } from 'svelte/store';

export const score = writable(['⚪', '⚪', '⚪', '⚪', '⚪']);

export const questionNumber = writable(0);
export const hintNumber = writable(0);
export const guessNumber = writable(0);
