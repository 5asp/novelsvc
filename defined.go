package novelsvc

type BookSource struct {
	SourceName string	`json:"source_name" comment:"书源名称"`
	SourceURL string 	`json:"source_url" comment:"书源网址"`
	SourceKey string 	`json:"source_key" comment:"书源标识"`

	SearchURL string 	`json:"search_url" comment:"搜索网址"`
	SearchItemRule string 	`json:"search_list_rule" comment:"搜索结果规则"`
	SearchItemNameRule string 	`json:"search_item_name_rule" comment:"搜索结果名称规则"`
	SearchItemAuthorRule string 	`json:"search_item_author_rule" comment:"搜索结果作者规则"`
	SearchItemCoverRule string 	`json:"search_item_cover_rule" comment:"搜索结果封面规则"`
	SearchItemCategoryRule string 	`json:"search_item_category_rule" comment:"搜索结果分类规则"`
	SearchItemNewChapterRule string 	`json:"search_item_new_chapter_rule" comment:"搜索结果最新章节规则"`
	SearchItemURLRule string 	`json:"search_item_url_rule" comment:"搜索结果链接规则"`


	DetailBookItemRule string 	`json:"detail_book_item_rule" comment:"小说详情规则"`
	DetailBookNameRule string 	`json:"detail_book_name_rule" comment:"小说名称规则"`
	DetailBookAuthorRule string 	`json:"detail_book_author_rule" comment:"小说作者规则"`
	DetailBookCoverRule string 	`json:"detail_book_cover_rule" comment:"小说封面规则"`
	DetailBookCategoryRule string 	`json:"detail_book_category_rule" comment:"小说分类规则"`
	DetailBookDescriptionRule string 	`json:"detail_book_description_rule" comment:"小说描述规则"`
	DetailBookNewChapterRule string 	`json:"detail_book_new_chapter_rule" comment:"小说最新章节规则"`

	DetailChapterListURLRule string 	`json:"detail_chapter_url_rule" comment:"小说章节列表链接规则"`
	DetailNewChapterRule string 	`json:"detail_new_chapter_rule" comment:"小说新章节规则"`
	DetailNewChapterTitleRule string 	`json:"detail_new_chapter_rule" comment:"小说新章节名称规则"`
	DetailNewChapterURLRule string 	`json:"detail_new_chapter_rule" comment:"小说新章节名称规则"`
	DetailChapterRule string	`json:"detail_chapter_list_rule" comment:"小说章节列表规则"`
	DetailChapterTitleRule string	`json:"detail_chapter_list_rule" comment:"小说章节名称规则"`
	DetailChapterURLRule string	`json:"detail_chapter_list_rule" comment:"小说章节链接规则"`

	ContentTitleRule string	`json:"chapter_content_rule" comment:"内容标题规则"`
	ContentTextRule string	`json:"chapter_content_rule" comment:"内容正文规则"`
	ContentPreviousURLRule string	`json:"chapter_previous_url_rule" comment:"内容上一章链接规则"`
	ContentNextURLRule string	`json:"chapter_next_url_rule" comment:"内容下一章链接规则"`

	Weight int 	`json:"weight" comment:"权重"`
}

type BookContent struct {
	Title string 	`json:"title" comment:"章节标题`
	Text string 	`json:"text" comment:"章节正文`
	DetailURL string 	`json:"detail_url" comment:"详情链接`
	PreviousURL string 	`json:"previous_url" comment:"章节链接`
	NextURL string 	`json:"next_url" comment:"章节链接`
	Source string	`json:"source" comment:"搜索结果来源"`
}

type BookChapter struct {
	Title string 	`json:"title" comment:"章节标题`
	ChapterURL string 	`json:"chapter_url" comment:"章节链接`
	DetailURL string 	`json:"detail_url" comment:"详情链接`
	Source string	`json:"source" comment:"搜索结果来源"`
}
type BookInfo struct {
	//Id_	string		`json:"id" bson:"_id" comment:"小说ID"`
	Name        string `json:"name" bson:"name" comment:"小说名称"`
	Author      string `json:"author" bson:"author" comment:"小说作者"`
	Cover       string `json:"cover" bson:"cover" comment:"小说封面"`
	Category    string `json:"category" bson:"category" comment:"小说分类"`
	Description string `json:"description" bson:"description" comment:"小说描述"`
	NewChapter  string `json:"new_chapter" bson:"new_chapter" comment:"搜索结果最新章节"`
	URL         string `json:"url" bson:"url" comment:"搜索结果链接"`
	Source      string `json:"source" bson:"source" comment:"搜索结果来源"`
}

var BookSources = map[int64]BookSource{
	1: {
		SourceName: "笔趣阁",
		SourceURL: "https://www.biquzw.la",
		SourceKey: "biquge",

		SearchURL: "https://www.biquzw.la/modules/article/search.php?searchkey=%s",
		SearchItemRule: "//tbody/tr[1]/following-sibling::tr",
		SearchItemNameRule: "./td[1]/a",
		SearchItemAuthorRule: "./td[3]",
		SearchItemCoverRule: "",
		SearchItemCategoryRule: "",
		SearchItemNewChapterRule: "./td[2]/a",
		SearchItemURLRule: "//td[1]/a",

		DetailBookItemRule: `//div[@id="wrapper"]`,
		DetailBookNameRule: `//div[@id="info"]/h1`,
		DetailBookAuthorRule: `//div[@id="info"]/p[1]`,
		DetailBookCoverRule: `//div[@id="fmimg"]/img`,
		DetailBookCategoryRule: `//div[@class="con_top"]/a[2]`,
		DetailBookDescriptionRule: `//div[@id="intro"]/p[1]`,
		DetailBookNewChapterRule: "",

		DetailChapterListURLRule: "",
		DetailNewChapterRule: "",
		DetailNewChapterTitleRule: "",
		DetailNewChapterURLRule: "",
		DetailChapterRule: `//div[@id="list"]/dl/dd`,
		DetailChapterTitleRule: `./a`,
		DetailChapterURLRule: `./a`,

		ContentTitleRule: `//div[@class="bookname"]/h1`,
		ContentTextRule: `//div[@id="content"]`,
		ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
		ContentNextURLRule: `//div[@class="bottem1"]/a[4]`,
	},
	// 3: {
	// 	SourceName: "笔趣阁【ivipxs】",
	// 	SourceURL: "https://www.ivipxs.com/",
	// 	SourceKey: "ivipxs",
	// 	SearchURL: "https://www.ivipxs.com/search.php?searchkey=%s",
	// 	SearchItemRule: `//div[@class="item"]`,
	// 	SearchItemNameRule: "//dt/a",
	// 	SearchItemAuthorRule: "//dt/span",
	// 	SearchItemCoverRule: "//img",
	// 	SearchItemCategoryRule: "",
	// 	SearchItemNewChapterRule: "",
	// 	SearchItemURLRule: `//dt/a`,

	// 	DetailBookItemRule: `//div[@id="wrapper"]`,
	// 	DetailBookNameRule: `//div[@id="info"]/h1`,
	// 	DetailBookAuthorRule: `//div[@id="info"]/p[1]`,
	// 	DetailBookCoverRule: `//div[@id="fmimg"]/img`,
	// 	DetailBookCategoryRule: `//div[@class="con_top"]/a[2]`,
	// 	DetailBookDescriptionRule: `//div[@id="intro"]/p[1]`,
	// 	DetailChapterListURLRule: "",
	// 	DetailNewChapterRule: `//table[@id="adt2"]/parent::div/preceding-sibling::dd`,
	// 	DetailNewChapterTitleRule: `./a`,
	// 	DetailNewChapterURLRule: `./a`,
	// 	DetailChapterRule: `//dt[2]/following-sibling::dd`,
	// 	DetailChapterTitleRule: `./a`,
	// 	DetailChapterURLRule: `./a`,

	// 	ContentTitleRule: `//div[@class="bookname"]/h1`,
	// 	ContentTextRule: `//div[@id="content"]`,
	// 	ContentPreviousURLRule: `//div[@class="bottem1"]/a[2]`,
	// 	ContentNextURLRule: `//div[@class="bottem1"]/a[4]`,
	// },
}