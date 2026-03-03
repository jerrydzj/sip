package poem

var poems = []Poem{
	{
		Title:  "静夜思",
		Author: "李白",
		Text:   "床前明月光，疑是地上霜。举头望明月，低头思故乡。",
		Moods:  []string{"nostalgic", "lonely", "calm"},
	},
	{
		Title:  "春晓",
		Author: "孟浩然",
		Text:   "春眠不觉晓，处处闻啼鸟。夜来风雨声，花落知多少。",
		Moods:  []string{"joy", "calm", "hopeful"},
	},
	{
		Title:  "登鹳雀楼",
		Author: "王之涣",
		Text:   "白日依山尽，黄河入海流。欲穷千里目，更上一层楼。",
		Moods:  []string{"hopeful", "inspired", "energetic"},
	},
	{
		Title:  "相思",
		Author: "王维",
		Text:   "红豆生南国，春来发几枝。愿君多采撷，此物最相思。",
		Moods:  []string{"nostalgic", "lonely", "love"},
	},
	{
		Title:  "江雪",
		Author: "柳宗元",
		Text:   "千山鸟飞绝，万径人踪灭。孤舟蓑笠翁，独钓寒江雪。",
		Moods:  []string{"sad", "lonely", "calm"},
	},
}
