package main

import (
	"flag"
	"fmt"
	"gen_vue/module"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var output string
	flag.StringVar(&output, "out", "", "文件输出路径")
	flag.Parse()

	// 读取配置文件
	var conf = module.Conf

	f, err := os.OpenFile("config/template.vue", os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		log.Printf("template open err #%v ", err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Printf("template read err #%v ", err)
	}

	contentStr := string(content)
	// 处理文件
	//处理table列
	columns := ""
	for k, v := range conf.COLUMNS {
		if k != len(conf.COLUMNS)-1 {
			columns += "{ label: '" + v.Label + "', prop: '" + v.Prop + "', width: 'auto' },\n"
		} else {
			columns += "{ label: '" + v.Label + "', prop: '" + v.Prop + "', width: 'auto' }"
		}

	}

	//处理listQuery参数
	params := ""
	for k, p := range conf.PARAM {
		if k != len(conf.PARAM)-1 {
			params += p + ": undefined,\n"
		} else {
			params += p + ": undefined"
		}

	}

	// 处理筛选项
	filters := ""
	for _, v := range conf.FILTER {
		switch v.Type {
		case "input":
			filters += `<el-input  v-model="listQuery.` + v.Prop + `" placeholder="` + v.Prop + `" size="mini" class="filter-item" style="width: 100px;" />`
			break
		case "select":
			filters += `<el-select v-model="listQuery.` + v.Prop + `"placeholder="` + v.Prop + `"clearable style="width: 100px" size="mini" class="filter-item">
							<el-option v-for="(value, key) in statusOptions" :key="value" :label="value" :value="key"/>
						</el-select>`
			break
		case "date":
			filters += `<el-date-picker v-model="listQuery.` + v.Prop + `" type="date" size="mini" class="filter-item" value-format=" yyyy-MM-dd" placeholder="请选择日期" />`
			break
		}
	}
	contentStr = strings.Replace(contentStr, "<!--NAME-->", conf.NAME, -1)
	contentStr = strings.Replace(contentStr, "<!--COLUMNS-->", columns, -1)
	contentStr = strings.Replace(contentStr, "<!--PARAM-->", params, -1)
	contentStr = strings.Replace(contentStr, "<!--FILTER-->", filters, -1)

	// 生成输出文件
	output = strings.Trim(output, " ")
	if output == "" {
		output = conf.NAME + ".vue"
	}
	if !strings.HasSuffix(output, ".vue") {
		output += "/" + conf.NAME + ".vue"
	}

	os.MkdirAll(filepath.Dir(output), os.ModePerm)

	f, err = os.Create(output)
	defer f.Close()
	if err != nil {
		log.Printf("create file err #%v", err)
	} else {
		_, err = f.WriteString(contentStr)
		if err != nil {
			log.Printf("write content err #%v", err)
		}
	}
	fmt.Println(output)
}
