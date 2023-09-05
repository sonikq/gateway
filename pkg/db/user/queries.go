package user_access

const (
	getForms              = `select test_abac.get_reports($1, $2);`
	checkFormEmpty        = `select test_abac.check_form_empty($1, $2);`
	getChapterData        = `select test_abac.get_chapter_data($1, $2);`
	getFormData           = `select test_abac.get_form_data($1, $2);`
	getActions            = `select test_abac.get_actions();`
	saveChapterData       = `select test_abac.save_chapter_data($1, $2::json);`
	getChapterFormula     = `select * from test_abac.get_stat_formula($1);`
	getChapterTargetCells = `select * from test_abac.get_chapter_formula($1);`
	getCasbinSubject      = `select * from test_abac.get_casbin_subject($1, $2);`
	getCasbinObject       = `select * from test_abac.get_casbin_object($1);`
	setFormNotEmpty       = `select * from stat.set_form_not_empty($1);`
	getNSI                = `select * from stat.get_nsi($1);`
	getFormChapters       = `select * from stat.get_form_chapters($1);`
	getFormIDByChapterID  = `select test_abac.get_formID_by_chapterID($1)`
	getTemplateJSON       = `select * from test_abac.get_template_json($1);`
)
