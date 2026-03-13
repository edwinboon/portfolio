package pages

func AboutView() string {
	title := StyleTitle.Render("About")

	p1 := StyleBody.Render("Edwin is a passionate software developer with over 12 years of experience,\ntransforming his love for programming from a hobby into a thriving career.\nAs a self-taught developer, he thrives on continuous learning and problem-solving,\nalways looking for new ways to level up his skills. Currently, he's diving deep\ninto backend development, terminals and shells, sharpening his expertise and\nfilling in the gaps he missed along the way.")

	p2 := StyleBody.Render("But coding isn't just a job for Edwin—it's a lifelong passion. When he's not\ncrafting efficient code, he's spending quality time with his wife and three\namazing kids, indulging in the simple joys of family life. Whether it's getting\nlost in a gripping series, battling it out in a video game, or enjoying the\nstrategy of a great board game, he loves unwinding with immersive experiences.")

	p3 := StyleBody.Render("And let's not forget his refined taste for good coffee and specialty beers—\nbecause life is too short for anything less than top-tier brews!")

	hint := Hint("q", "back to menu")

	content := title + "\n\n" + p1 + "\n\n" + p2 + "\n\n" + p3 + "\n\n" + hint
	return StylePage.Render(content)
}
