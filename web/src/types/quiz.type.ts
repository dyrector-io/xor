export type QuizItem = {
	Name: string;
	CrunchbaseDescription: string;
	GithubDescription: string;
	Logo: string;
	Category: string;
	Subcategory: string;
	GithubStars: number;
};

export type QuizResponse = {
	List: Array<QuizItem>;
	Date: string;
};
