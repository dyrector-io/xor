export type QuizItem = {
	Name: string;
	CrunchbaseDescription: string;
	GithubDescription: string;
	Logo: string;
	Category: string;
	Subcategory: string;
	GithubStars: number;
	GithubContributorsCount: number;
};

export type JavascriptQuizItem = {
	Name: string;
	Description: string;
	CodeExample: string;
	RandomFact: string;
	GitHubLink: string;
	GithubStars: number;
	WeeklyDownloads: number;
};

export type QuizResponse = {
	List: Array<JavascriptQuizItem>;
	Date: string;
};
