package selectfield

import (
	"encoding/json"
	"strings"

	"github.com/quarkcms/quark-go/pkg/component/admin/component"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/when"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/rule"
	"github.com/quarkcms/quark-go/pkg/component/admin/table"
	"github.com/quarkcms/quark-go/pkg/untils"
)

type FieldNames struct {
	Label    string `json:"label"`
	Value    string `json:"value"`
	Children string `json:"children"`
}

type Option struct {
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Disabled bool        `json:"disabled,omitempty"`
}

type SelectField struct {
	ComponentKey string `json:"componentkey"` // 组件标识
	Component    string `json:"component"`    // 组件名称

	Colon         bool        `json:"colon,omitempty"`        // 配合 label 属性使用，表示是否显示 label 后面的冒号
	Extra         string      `json:"extra,omitempty"`        // 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
	HasFeedback   bool        `json:"hasFeedback,omitempty"`  // 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
	Help          string      `json:"help,omitempty"`         // 提示信息，如不设置，则会根据校验规则自动生成
	Hidden        bool        `json:"hidden,omitempty"`       // 是否隐藏字段（依然会收集和校验字段）
	InitialValue  interface{} `json:"initialValue,omitempty"` // 设置子元素默认值，如果与 Form 的 initialValues 冲突则以 Form 为准
	Label         string      `json:"label,omitempty"`        // label 标签的文本
	LabelAlign    string      `json:"labelAlign,omitempty"`   // 标签文本对齐方式
	LabelCol      interface{} `json:"labelCol,omitempty"`     // label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。你可以通过 Form 的 labelCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准
	Name          string      `json:"name,omitempty"`         // 字段名，支持数组
	NoStyle       bool        `json:"noStyle,omitempty"`      // 为 true 时不带样式，作为纯字段控件使用
	Required      bool        `json:"required,omitempty"`     // 必填样式设置。如不设置，则会根据校验规则自动生成
	Tooltip       string      `json:"tooltip,omitempty"`      // 会在 label 旁增加一个 icon，悬浮后展示配置的信息
	ValuePropName string      `json:"valuePropName"`          // 子节点的值的属性，如 Switch 的是 'checked'。该属性为 getValueProps 的封装，自定义 getValueProps 后会失效
	WrapperCol    interface{} `json:"wrapperCol"`             // 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。你可以通过 Form 的 wrapperCol 进行统一设置，不会作用于嵌套 Item。当和 Form 同时设置时，以 Item 为准

	Api            string        `json:"api,omitempty"` // 获取数据接口
	Ignore         bool          `json:"ignore"`        // 是否忽略保存到数据库，默认为 false
	Rules          []*rule.Rule  `json:"-"`             // 全局校验规则
	CreationRules  []*rule.Rule  `json:"-"`             // 创建页校验规则
	UpdateRules    []*rule.Rule  `json:"-"`             // 编辑页校验规则
	FrontendRules  []*rule.Rule  `json:"frontendRules"` // 前端校验规则，设置字段的校验逻辑
	When           *when.When    `json:"when"`          //
	WhenItem       []*when.Item  `json:"-"`             //
	ShowOnIndex    bool          `json:"-"`             // 在列表页展示
	ShowOnDetail   bool          `json:"-"`             // 在详情页展示
	ShowOnCreation bool          `json:"-"`             // 在创建页面展示
	ShowOnUpdate   bool          `json:"-"`             // 在编辑页面展示
	ShowOnExport   bool          `json:"-"`             // 在导出的Excel上展示
	ShowOnImport   bool          `json:"-"`             // 在导入Excel上展示
	Editable       bool          `json:"-"`             // 表格上是否可编辑
	Column         *table.Column `json:"-"`             // 表格列
	Callback       interface{}   `json:"-"`             // 回调函数

	AllowClear               bool                   `json:"allowClear,omitempty"`               // 可以点击清除图标删除内容
	AutoClearSearchValue     bool                   `json:"autoClearSearchValue,omitempty"`     // 是否在选中项后清空搜索框，只在 mode 为 multiple 或 tags 时有效
	AutoFocus                bool                   `json:"autoFocus,omitempty"`                // 默认获取焦点
	Bordered                 bool                   `json:"bordered,omitempty"`                 // 是否有边框
	ClearIcon                interface{}            `json:"clearIcon,omitempty"`                // 自定义的多选框清空图标
	DefaultActiveFirstOption bool                   `json:"defaultActiveFirstOption,omitempty"` // 是否默认高亮第一个选项
	DefaultOpen              bool                   `json:"defaultOpen,omitempty"`              // 是否默认展开下拉菜单
	DefaultValue             interface{}            `json:"defaultValue,omitempty"`             // 默认选中的选项
	Disabled                 bool                   `json:"disabled,omitempty"`                 // 整组失效
	PopupClassName           string                 `json:"popupClassName,omitempty"`           // 下拉菜单的 className 属性
	DropdownMatchSelectWidth interface{}            `json:"dropdownMatchSelectWidth,omitempty"` // 下拉菜单和选择器同宽。默认将设置 min-width，当值小于选择框宽度时会被忽略。false 时会关闭虚拟滚动
	DropdownStyle            interface{}            `json:"dropdownStyle,omitempty"`            // 下拉菜单的 style 属性
	FieldNames               *FieldNames            `json:"fieldNames,omitempty"`               // 自定义 options 中 label value children 的字段
	LabelInValue             bool                   `json:"labelInValue,omitempty"`             // 是否把每个选项的 label 包装到 value 中，会把 Select 的 value 类型从 string 变为 { value: string, label: ReactNode } 的格式
	ListHeight               int                    `json:"listHeight,omitempty"`               // 设置弹窗滚动高度 256
	Loading                  bool                   `json:"loading,omitempty"`                  // 加载中状态
	MaxTagCount              int                    `json:"maxTagCount,omitempty"`              // 最多显示多少个 tag，响应式模式会对性能产生损耗
	MaxTagPlaceholder        string                 `json:"maxTagPlaceholder,omitempty"`        // 隐藏 tag 时显示的内容
	MaxTagTextLength         int                    `json:"maxTagTextLength,omitempty"`         // 最大显示的 tag 文本长度
	MenuItemSelectedIcon     interface{}            `json:"menuItemSelectedIcon,omitempty"`     // 自定义多选时当前选中的条目图标
	Mode                     string                 `json:"mode,omitempty"`                     // 设置 Select 的模式为多选或标签 multiple | tags
	NotFoundContent          string                 `json:"notFoundContent,omitempty"`          // 当下拉列表为空时显示的内容
	Open                     bool                   `json:"open,omitempty"`                     // 是否展开下拉菜单
	OptionFilterProp         string                 `json:"optionFilterProp,omitempty"`         // 搜索时过滤对应的 option 属性，如设置为 children 表示对内嵌内容进行搜索。若通过 options 属性配置选项内容，建议设置 optionFilterProp="label" 来对内容进行搜索。
	OptionLabelProp          string                 `json:"optionLabelProp,omitempty"`          // 回填到选择框的 Option 的属性值，默认是 Option 的子元素。比如在子元素需要高亮效果时，此值可以设为 value。
	Options                  []*Option              `json:"options,omitempty"`                  // 可选项数据源
	Placeholder              string                 `json:"placeholder,omitempty"`              // 选择框默认文本
	Placement                string                 `json:"placement,omitempty"`                // 选择框弹出的位置 bottomLeft bottomRight topLeft topRight
	RemoveIcon               interface{}            `json:"removeIcon,omitempty"`               // 自定义的多选框清除图标
	SearchValue              string                 `json:"searchValue,omitempty"`              // 控制搜索文本
	ShowArrow                bool                   `json:"showArrow,omitempty"`                // 是否显示下拉小箭头
	ShowSearch               bool                   `json:"showSearch,omitempty"`               // 配置是否可搜索
	Size                     string                 `json:"size,omitempty"`                     // 选择框大小
	Status                   string                 `json:"status,omitempty"`                   // 设置校验状态 'error' | 'warning'
	SuffixIcon               interface{}            `json:"suffixIcon,omitempty"`               // 自定义的选择框后缀图标
	TokenSeparators          interface{}            `json:"tokenSeparators,omitempty"`          // 自动分词的分隔符，仅在 mode="tags" 时生效
	Value                    interface{}            `json:"value,omitempty"`                    // 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
	Virtual                  bool                   `json:"virtual,omitempty"`                  // 设置 false 时关闭虚拟滚动
	Style                    map[string]interface{} `json:"style,omitempty"`                    // 自定义样式
}

// 初始化组件
func New() *SelectField {
	return (&SelectField{}).Init()
}

// 初始化
func (p *SelectField) Init() *SelectField {
	p.Component = "selectField"
	p.Colon = true
	p.LabelAlign = "right"
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = true
	p.ShowOnImport = true
	p.AllowClear = true
	p.Column = (&table.Column{}).Init()
	p.SetWidth(200)
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

// 设置Key
func (p *SelectField) SetKey(key string, crypt bool) *SelectField {
	p.ComponentKey = untils.MakeKey(key, crypt)

	return p
}

// 会在 label 旁增加一个 icon，悬浮后展示配置的信息
func (p *SelectField) SetTooltip(tooltip string) *SelectField {
	p.Tooltip = tooltip

	return p
}

// Field 的长度，我们归纳了常用的 Field 长度以及适合的场景，支持了一些枚举 "xs" , "s" , "m" , "l" , "x"
func (p *SelectField) SetWidth(width interface{}) *SelectField {
	style := make(map[string]interface{})

	for k, v := range p.Style {
		style[k] = v
	}

	style["width"] = width
	p.Style = style

	return p
}

// 配合 label 属性使用，表示是否显示 label 后面的冒号
func (p *SelectField) SetColon(colon bool) *SelectField {
	p.Colon = colon
	return p
}

// 额外的提示信息，和 help 类似，当需要错误信息和提示文案同时出现时，可以使用这个。
func (p *SelectField) SetExtra(extra string) *SelectField {
	p.Extra = extra
	return p
}

// 配合 validateStatus 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *SelectField) SetHasFeedback(hasFeedback bool) *SelectField {
	p.HasFeedback = hasFeedback
	return p
}

// 配合 help 属性使用，展示校验状态图标，建议只配合 Input 组件使用
func (p *SelectField) SetHelp(help string) *SelectField {
	p.Help = help
	return p
}

// 为 true 时不带样式，作为纯字段控件使用
func (p *SelectField) SetNoStyle() *SelectField {
	p.NoStyle = true
	return p
}

// label 标签的文本
func (p *SelectField) SetLabel(label string) *SelectField {
	p.Label = label

	return p
}

// 标签文本对齐方式
func (p *SelectField) SetLabelAlign(align string) *SelectField {
	p.LabelAlign = align
	return p
}

// label 标签布局，同 <Col> 组件，设置 span offset 值，如 {span: 3, offset: 12} 或 sm: {span: 3, offset: 12}。
// 你可以通过 Form 的 labelCol 进行统一设置。当和 Form 同时设置时，以 Item 为准
func (p *SelectField) SetLabelCol(col interface{}) *SelectField {
	p.LabelCol = col
	return p
}

// 字段名，支持数组
func (p *SelectField) SetName(name string) *SelectField {
	p.Name = name
	return p
}

// 是否必填，如不设置，则会根据校验规则自动生成
func (p *SelectField) SetRequired() *SelectField {
	p.Required = true
	return p
}

// 获取前端验证规则
func (p *SelectField) GetFrontendRules(path string) *SelectField {
	var (
		frontendRules []*rule.Rule
		rules         []*rule.Rule
		creationRules []*rule.Rule
		updateRules   []*rule.Rule
	)

	uri := strings.Split(path, "/")
	isCreating := (uri[len(uri)-1] == "create") || (uri[len(uri)-1] == "store")
	isEditing := (uri[len(uri)-1] == "edit") || (uri[len(uri)-1] == "update")

	if len(p.Rules) > 0 {
		rules = rule.ConvertToFrontendRules(p.Rules)
	}
	if isCreating && len(p.CreationRules) > 0 {
		creationRules = rule.ConvertToFrontendRules(p.CreationRules)
	}
	if isEditing && len(p.UpdateRules) > 0 {
		updateRules = rule.ConvertToFrontendRules(p.UpdateRules)
	}
	if len(rules) > 0 {
		frontendRules = append(frontendRules, rules...)
	}
	if len(creationRules) > 0 {
		frontendRules = append(frontendRules, creationRules...)
	}
	if len(updateRules) > 0 {
		frontendRules = append(frontendRules, updateRules...)
	}

	p.FrontendRules = frontendRules

	return p
}

// 校验规则，设置字段的校验逻辑
func (p *SelectField) SetRules(rules []*rule.Rule) *SelectField {
	p.Rules = rules

	return p
}

// 校验规则，只在创建表单提交时生效
func (p *SelectField) SetCreationRules(rules []*rule.Rule) *SelectField {
	p.CreationRules = rules

	return p
}

// 校验规则，只在更新表单提交时生效
func (p *SelectField) SetUpdateRules(rules []*rule.Rule) *SelectField {
	p.UpdateRules = rules

	return p
}

// 子节点的值的属性，如 Switch 的是 "checked"
func (p *SelectField) SetValuePropName(valuePropName string) *SelectField {
	p.ValuePropName = valuePropName
	return p
}

// 需要为输入控件设置布局样式时，使用该属性，用法同 labelCol。
// 你可以通过 Form 的 wrapperCol 进行统一设置。当和 Form 同时设置时，以 Item 为准。
func (p *SelectField) SetWrapperCol(col interface{}) *SelectField {
	p.WrapperCol = col
	return p
}

// 指定当前选中的条目，多选时为一个数组。（value 数组引用未变化时，Select 不会更新）
func (p *SelectField) SetValue(value interface{}) *SelectField {
	p.Value = value
	return p
}

// 设置默认值。
func (p *SelectField) SetDefault(value interface{}) *SelectField {
	p.DefaultValue = value
	return p
}

// 是否禁用状态，默认为 false
func (p *SelectField) SetDisabled(disabled bool) *SelectField {
	p.Disabled = disabled
	return p
}

// 是否忽略保存到数据库，默认为 false
func (p *SelectField) SetIgnore(ignore bool) *SelectField {
	p.Ignore = ignore
	return p
}

// 表单联动
func (p *SelectField) SetWhen(value ...any) *SelectField {
	w := when.New()
	i := when.NewItem()
	var operator string
	var option any

	if len(value) == 2 {
		operator = "="
		option = value[0]
		callback := value[1].(func() interface{})

		i.Body = callback()
	}

	if len(value) == 3 {
		operator = value[0].(string)
		option = value[1]
		callback := value[2].(func() interface{})

		i.Body = callback()
	}

	getOption := untils.InterfaceToString(option)

	switch operator {
	case "=":
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	case ">":
		i.Condition = "<%=String(" + p.Name + ") > '" + getOption + "' %>"
		break
	case "<":
		i.Condition = "<%=String(" + p.Name + ") < '" + getOption + "' %>"
		break
	case "<=":
		i.Condition = "<%=String(" + p.Name + ") <= '" + getOption + "' %>"
		break
	case ">=":
		i.Condition = "<%=String(" + p.Name + ") => '" + getOption + "' %>"
		break
	case "has":
		i.Condition = "<%=(String(" + p.Name + ").indexOf('" + getOption + "') !=-1) %>"
		break
	case "in":
		jsonStr, _ := json.Marshal(option)
		i.Condition = "<%=(" + string(jsonStr) + ".indexOf(" + p.Name + ") !=-1) %>"
		break
	default:
		i.Condition = "<%=String(" + p.Name + ") === '" + getOption + "' %>"
		break
	}

	i.ConditionName = p.Name
	i.ConditionOperator = operator
	i.Option = option
	p.WhenItem = append(p.WhenItem, i)
	p.When = w.SetItems(p.WhenItem)

	return p
}

// Specify that the element should be hidden from the index view.
func (p *SelectField) HideFromIndex(callback bool) *SelectField {
	p.ShowOnIndex = !callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *SelectField) HideFromDetail(callback bool) *SelectField {
	p.ShowOnDetail = !callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *SelectField) HideWhenCreating(callback bool) *SelectField {
	p.ShowOnCreation = !callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *SelectField) HideWhenUpdating(callback bool) *SelectField {
	p.ShowOnUpdate = !callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *SelectField) HideWhenExporting(callback bool) *SelectField {
	p.ShowOnExport = !callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *SelectField) HideWhenImporting(callback bool) *SelectField {
	p.ShowOnImport = !callback

	return p
}

// Specify that the element should be hidden from the index view.
func (p *SelectField) OnIndexShowing(callback bool) *SelectField {
	p.ShowOnIndex = callback

	return p
}

// Specify that the element should be hidden from the detail view.
func (p *SelectField) OnDetailShowing(callback bool) *SelectField {
	p.ShowOnDetail = callback

	return p
}

// Specify that the element should be hidden from the creation view.
func (p *SelectField) ShowOnCreating(callback bool) *SelectField {
	p.ShowOnCreation = callback

	return p
}

// Specify that the element should be hidden from the update view.
func (p *SelectField) ShowOnUpdating(callback bool) *SelectField {
	p.ShowOnUpdate = callback

	return p
}

// Specify that the element should be hidden from the export file.
func (p *SelectField) ShowOnExporting(callback bool) *SelectField {
	p.ShowOnExport = callback

	return p
}

// Specify that the element should be hidden from the import file.
func (p *SelectField) ShowOnImporting(callback bool) *SelectField {
	p.ShowOnImport = callback

	return p
}

// Specify that the element should only be shown on the index view.
func (p *SelectField) OnlyOnIndex() *SelectField {
	p.ShowOnIndex = true
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on the detail view.
func (p *SelectField) OnlyOnDetail() *SelectField {
	p.ShowOnIndex = false
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on forms.
func (p *SelectField) OnlyOnForms() *SelectField {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = true
	p.ShowOnUpdate = true
	p.ShowOnExport = false
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on export file.
func (p *SelectField) OnlyOnExport() *SelectField {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = false

	return p
}

// Specify that the element should only be shown on import file.
func (p *SelectField) OnlyOnImport() *SelectField {
	p.ShowOnIndex = false
	p.ShowOnDetail = false
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = false
	p.ShowOnImport = true

	return p
}

// Specify that the element should be hidden from forms.
func (p *SelectField) ExceptOnForms() *SelectField {
	p.ShowOnIndex = true
	p.ShowOnDetail = true
	p.ShowOnCreation = false
	p.ShowOnUpdate = false
	p.ShowOnExport = true
	p.ShowOnImport = true

	return p
}

// Check for showing when updating.
func (p *SelectField) IsShownOnUpdate() bool {
	return p.ShowOnUpdate
}

// Check showing on index.
func (p *SelectField) IsShownOnIndex() bool {
	return p.ShowOnIndex
}

// Check showing on detail.
func (p *SelectField) IsShownOnDetail() bool {
	return p.ShowOnDetail
}

// Check for showing when creating.
func (p *SelectField) IsShownOnCreation() bool {
	return p.ShowOnCreation
}

// Check for showing when exporting.
func (p *SelectField) IsShownOnExport() bool {
	return p.ShowOnExport
}

// Check for showing when importing.
func (p *SelectField) IsShownOnImport() bool {
	return p.ShowOnImport
}

// 设置为可编辑列
func (p *SelectField) SetEditable(editable bool) *SelectField {
	p.Editable = editable

	return p
}

// 闭包，透传表格列的属性
func (p *SelectField) SetColumn(f func(column *table.Column) *table.Column) *SelectField {
	p.Column = f(p.Column)

	return p
}

// 当前列值的枚举 valueEnum
func (p *SelectField) GetValueEnum() map[interface{}]interface{} {
	data := map[interface{}]interface{}{}

	return data
}

// 设置回调函数
func (p *SelectField) SetCallback(closure func() interface{}) *SelectField {
	if closure != nil {
		p.Callback = closure
	}

	return p
}

// 获取回调函数
func (p *SelectField) GetCallback() interface{} {
	return p.Callback
}

// 设置属性
func (p *SelectField) SetOptions(options []*Option) *SelectField {
	p.Options = options

	return p
}

// 获取数据接口
func (p *SelectField) SetApi(api string) *SelectField {
	p.Api = api

	return p
}

// 可以点击清除图标删除内容
func (p *SelectField) SetAllowClear(allowClear bool) *SelectField {
	p.AllowClear = allowClear

	return p
}

// 是否在选中项后清空搜索框，只在 mode 为 multiple 或 tags 时有效
func (p *SelectField) SetAutoClearSearchValue(autoClearSearchValue bool) *SelectField {
	p.AutoClearSearchValue = autoClearSearchValue

	return p
}

// 默认获取焦点
func (p *SelectField) SetAutoFocus(autoFocus bool) *SelectField {
	p.AutoFocus = autoFocus

	return p
}

// 默认获取焦点
func (p *SelectField) SetBordered(bordered bool) *SelectField {
	p.Bordered = bordered

	return p
}

// 自定义的多选框清空图标
func (p *SelectField) SetClearIcon(clearIcon interface{}) *SelectField {
	p.ClearIcon = clearIcon

	return p
}

// 是否默认高亮第一个选项
func (p *SelectField) SetDefaultActiveFirstOption(defaultActiveFirstOption bool) *SelectField {
	p.DefaultActiveFirstOption = defaultActiveFirstOption

	return p
}

// 是否默认展开下拉菜单
func (p *SelectField) SetDefaultOpen(defaultOpen bool) *SelectField {
	p.DefaultOpen = defaultOpen

	return p
}

// 下拉菜单的 className 属性
func (p *SelectField) SetPopupClassName(popupClassName string) *SelectField {
	p.PopupClassName = popupClassName

	return p
}

// 下拉菜单和选择器同宽。默认将设置 min-width，当值小于选择框宽度时会被忽略。false 时会关闭虚拟滚动
func (p *SelectField) SetDropdownMatchSelectWidth(dropdownMatchSelectWidth interface{}) *SelectField {
	p.DropdownMatchSelectWidth = dropdownMatchSelectWidth

	return p
}

// 下拉菜单的 style 属性
func (p *SelectField) SetDropdownStyle(dropdownStyle interface{}) *SelectField {
	p.DropdownStyle = dropdownStyle

	return p
}

// 自定义 options 中 label value children 的字段
func (p *SelectField) SetFieldNames(fieldNames *FieldNames) *SelectField {
	p.FieldNames = fieldNames

	return p
}

// 是否把每个选项的 label 包装到 value 中，会把 Select 的 value 类型从 string 变为 { value: string, label: ReactNode } 的格式
func (p *SelectField) SetLabelInValue(labelInValue bool) *SelectField {
	p.LabelInValue = labelInValue

	return p
}

// 设置弹窗滚动高度 256
func (p *SelectField) SetListHeight(listHeight int) *SelectField {
	p.ListHeight = listHeight

	return p
}

// 加载中状态
func (p *SelectField) SetLoading(loading bool) *SelectField {
	p.Loading = loading

	return p
}

// 最多显示多少个 tag，响应式模式会对性能产生损耗
func (p *SelectField) SetMaxTagCount(maxTagCount int) *SelectField {
	p.MaxTagCount = maxTagCount

	return p
}

// 隐藏 tag 时显示的内容
func (p *SelectField) SetMaxTagPlaceholder(maxTagPlaceholder string) *SelectField {
	p.MaxTagPlaceholder = maxTagPlaceholder

	return p
}

// 最大显示的 tag 文本长度
func (p *SelectField) SetMaxTagTextLength(maxTagTextLength int) *SelectField {
	p.MaxTagTextLength = maxTagTextLength

	return p
}

// 自定义多选时当前选中的条目图标
func (p *SelectField) SetMenuItemSelectedIcon(menuItemSelectedIcon interface{}) *SelectField {
	p.MenuItemSelectedIcon = menuItemSelectedIcon

	return p
}

// 设置 Select 的模式为多选或标签 multiple | tags
func (p *SelectField) SetMode(mode string) *SelectField {
	p.Mode = mode

	return p
}

// 当下拉列表为空时显示的内容
func (p *SelectField) SetNotFoundContent(notFoundContent string) *SelectField {
	p.NotFoundContent = notFoundContent

	return p
}

// 是否展开下拉菜单
func (p *SelectField) SetOpen(open bool) *SelectField {
	p.Open = open

	return p
}

// 搜索时过滤对应的 option 属性，如设置为 children 表示对内嵌内容进行搜索。若通过 options 属性配置选项内容，建议设置 optionFilterProp="label" 来对内容进行搜索。
func (p *SelectField) SetOptionFilterProp(optionFilterProp string) *SelectField {
	p.OptionFilterProp = optionFilterProp

	return p
}

// 回填到选择框的 Option 的属性值，默认是 Option 的子元素。比如在子元素需要高亮效果时，此值可以设为 value。
func (p *SelectField) SetOptionLabelProp(optionLabelProp string) *SelectField {
	p.OptionLabelProp = optionLabelProp

	return p
}

// 选择框默认文本
func (p *SelectField) SetPlaceholder(placeholder string) *SelectField {
	p.Placeholder = placeholder

	return p
}

// 选择框弹出的位置 bottomLeft bottomRight topLeft topRight
func (p *SelectField) SetPlacement(placement string) *SelectField {
	p.Placement = placement

	return p
}

// 自定义的多选框清除图标
func (p *SelectField) SetRemoveIcon(removeIcon interface{}) *SelectField {
	p.RemoveIcon = removeIcon

	return p
}

// 控制搜索文本
func (p *SelectField) SetSearchValue(searchValue string) *SelectField {
	p.SearchValue = searchValue

	return p
}

// 是否显示下拉小箭头
func (p *SelectField) SetShowArrow(showArrow bool) *SelectField {
	p.ShowArrow = showArrow

	return p
}

// 配置是否可搜索
func (p *SelectField) SetShowSearch(showSearch bool) *SelectField {
	p.ShowSearch = showSearch

	return p
}

// 选择框大小
func (p *SelectField) SetSize(size string) *SelectField {
	p.Size = size

	return p
}

// 设置校验状态 'error' | 'warning'
func (p *SelectField) SetStatus(status string) *SelectField {
	p.Status = status

	return p
}

// 自定义的选择框后缀图标
func (p *SelectField) SetSuffixIcon(suffixIcon interface{}) *SelectField {
	p.SuffixIcon = suffixIcon

	return p
}

// 自动分词的分隔符，仅在 mode="tags" 时生效
func (p *SelectField) SetTokenSeparators(tokenSeparators interface{}) *SelectField {
	p.TokenSeparators = tokenSeparators

	return p
}

// 设置 false 时关闭虚拟滚动
func (p *SelectField) SetVirtual(virtual bool) *SelectField {
	p.Virtual = virtual

	return p
}