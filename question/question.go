package question

import "time"

type List []Item

type Item struct {
	Id   int
	Date string
	Url  string
}

func GetList() List {
	return List{
		{
			44538, "28.02.2022", "https://фкт-алтай.рф/qa/question/view-44538",
		},
		{
			44612, "03.03.2022", "https://фкт-алтай.рф/qa/question/view-44612",
		},
		{
			44707, "07.03.2022", "https://фкт-алтай.рф/qa/question/view-44707",
		},
		{
			44757, "17.03.2022", "https://фкт-алтай.рф/qa/question/view-44757",
		},
		{
			44883, "23.03.2022", "https://фкт-алтай.рф/qa/question/view-44883",
		},
		{
			44962, "30.03.2022", "https://фкт-алтай.рф/qa/question/view-44962",
		},
		{
			45044, "08.04.2022", "https://фкт-алтай.рф/qa/question/view-45044",
		},
		{
			35650, "13.04.2022", "https://фкт-алтай.рф/qa/question/view-35650",
		},
		{
			35298, "20.04.2022", "https://фкт-алтай.рф/qa/question/view-35298",
		},
		{
			4604, "02.05.2022", "https://фкт-алтай.рф/qa/question/view-4604",
		},
		{
			7533, "08.05.2022", "https://фкт-алтай.рф/qa/question/view-7533",
		},
		{
			23174, "18.05.2022", "https://фкт-алтай.рф/qa/question/view-23174",
		},
		{
			37945, "26.05.2022", "https://фкт-алтай.рф/qa/question/view-37945",
		},
		{
			12422, "02.06.2022", "https://фкт-алтай.рф/qa/question/view-12422",
		},
		{
			25867, "15.06.2022", "https://фкт-алтай.рф/qa/question/view-25867",
		},
		{
			14365, "24.06.2022", "https://фкт-алтай.рф/qa/question/view-14365",
		},
		{
			34312, "10.07.2022", "https://фкт-алтай.рф/qa/question/view-34312",
		},
		{
			37694, "25.07.2022", "https://фкт-алтай.рф/qa/question/view-37694",
		},
		{
			7279, "09.08.2022", "https://фкт-алтай.рф/qa/question/view-7279",
		},
		{
			2656, "04.09.2022", "https://фкт-алтай.рф/qa/question/view-2656",
		},
		{
			12734, "16.09.2022", "https://фкт-алтай.рф/qa/question/view-12734",
		},
		{
			3893, "24.09.2022", "https://фкт-алтай.рф/qa/question/view-3893",
		},
		{
			4910, "04.10.2022", "https://фкт-алтай.рф/qa/question/view-4910",
		},
		{
			3467, "16.10.2022", "https://фкт-алтай.рф/qa/question/view-3467",
		},
		{
			21294, "29.10.2022", "https://фкт-алтай.рф/qa/question/view-21294",
		},
	}
}

func GetCurrent() Item {
	return Item{
		41574,
		time.Now().String(),
		"https://фкт-алтай.рф/qa/question/view-41574",
	}
}
