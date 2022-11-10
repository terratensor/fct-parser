package question

import "time"

type List []Item

type Item struct {
	Id    int
	Num   string
	Date  string
	Url   string
	Count int
}

func GetList() List {
	return List{
		{
			44538, "01", "28.02.2022", "https://фкт-алтай.рф/qa/question/view-44538", 1312,
		},
		{
			44612, "02", "03.03.2022", "https://фкт-алтай.рф/qa/question/view-44612", 843,
		},
		{
			44707, "03", "07.03.2022", "https://фкт-алтай.рф/qa/question/view-44707", 827,
		},
		{
			44757, "04", "17.03.2022", "https://фкт-алтай.рф/qa/question/view-44757", 749,
		},
		{
			44883, "05", "23.03.2022", "https://фкт-алтай.рф/qa/question/view-44883", 469,
		},
		{
			44962, "06", "30.03.2022", "https://фкт-алтай.рф/qa/question/view-44962", 658,
		},
		{
			45044, "07", "08.04.2022", "https://фкт-алтай.рф/qa/question/view-45044", 1259,
		},
		{
			35650, "08", "13.04.2022", "https://фкт-алтай.рф/qa/question/view-35650", 735,
		},
		{
			35298, "09", "20.04.2022", "https://фкт-алтай.рф/qa/question/view-35298", 728,
		},
		{
			4604, "10", "02.05.2022", "https://фкт-алтай.рф/qa/question/view-4604", 770,
		},
		{
			7533, "11", "08.05.2022", "https://фкт-алтай.рф/qa/question/view-7533", 754,
		},
		{
			23174, "12", "18.05.2022", "https://фкт-алтай.рф/qa/question/view-23174", 749,
		},
		{
			37945, "13", "26.05.2022", "https://фкт-алтай.рф/qa/question/view-37945", 752,
		},
		{
			12422, "14", "02.06.2022", "https://фкт-алтай.рф/qa/question/view-12422", 793,
		},
		{
			25867, "15", "15.06.2022", "https://фкт-алтай.рф/qa/question/view-25867", 793,
		},
		{
			14365, "16", "24.06.2022", "https://фкт-алтай.рф/qa/question/view-14365", 880,
		},
		{
			34312, "17", "10.07.2022", "https://фкт-алтай.рф/qa/question/view-34312", 903,
		},
		{
			37694, "18", "25.07.2022", "https://фкт-алтай.рф/qa/question/view-37694", 761,
		},
		{
			7279, "19", "09.08.2022", "https://фкт-алтай.рф/qa/question/view-7279", 889,
		},
		{
			2656, "20", "04.09.2022", "https://фкт-алтай.рф/qa/question/view-2656", 968,
		},
		{
			12734, "21", "16.09.2022", "https://фкт-алтай.рф/qa/question/view-12734", 961,
		},
		{
			3893, "22", "24.09.2022", "https://фкт-алтай.рф/qa/question/view-3893", 1050,
		},
		{
			4910, "23", "04.10.2022", "https://фкт-алтай.рф/qa/question/view-4910", 972,
		},
		{
			3467, "24", "16.10.2022", "https://фкт-алтай.рф/qa/question/view-3467", 1004,
		},
		{
			21294, "25", "29.10.2022", "https://фкт-алтай.рф/qa/question/view-21294", 930,
		},
	}
}

func GetCurrent() Item {
	return Item{
		41574,
		"26",
		time.Now().String(),
		"https://фкт-алтай.рф/qa/question/view-41574",
		665,
	}
}
