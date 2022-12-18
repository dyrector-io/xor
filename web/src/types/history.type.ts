export type HistoryItem = {
	Date: string;
	Projects: string;
};

export type HistoryResponse = {
	result: Array<HistoryItem>;
};
