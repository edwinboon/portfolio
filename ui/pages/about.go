package pages

func AboutView() string {
	about := `# About Edwin

Edwin is a passionate software developer with over 12 years of experience,
transforming his love for programming from a hobby into a thriving career.
As a self-taught developer, he thrives on continuous learning and problem-solving,
always looking for new ways to level up his skills. Currently, he's diving deep
into backend development, terminals and shells, sharpening his expertise and
filling in the gaps he missed along the way.

But coding isn't just a job for Edwin—it's a lifelong passion. When he's not
crafting efficient code, he's spending quality time with his wife and three
amazing kids, indulging in the simple joys of family life. Whether it's getting
lost in a gripping series, battling it out in a video game, or enjoying the
strategy of a great board game, he loves unwinding with immersive experiences.

And let's not forget his refined taste for good coffee and specialty beers—
because life is too short for anything less than top-tier brews!
`

	markdown := RenderMarkdown(about)
	hint := Hint("q", "back to menu")
	return StylePage.Render(markdown + hint)
}
