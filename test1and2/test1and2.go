package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"test1and2/model"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Student struct {
	ID           int `gorm:"primaryKey"`
	TID          string
	Name         string
	City         string
	Gender       string
	Height       float64
	Score        map[string]float64
	Constitution string
}

// 数据库的100条数据由随机数生成

func main() {
	// 连接数据库
	model.Database("root:a291905714@/congz?charset=utf8&parseTime=True&loc=Local")
	var dbData []*model.Info
	// 从数据库读取数据
	model.DB.Model(model.Info{}).Find(&dbData)
	// 从文本读取数据并合并
	// txt文件数据为8条
	// 成绩不一致取平均，其他项不一致就用数据库的数据
	readFile("./student.txt", dbData)
	// 整合数据
	finalInfo := InteData(dbData)
	// 输出合并之后的100条数据
	for _, v := range finalInfo {
		fmt.Println(v)
	}
	fmt.Println()

	// 实验一
	// 数据分析
	analysis(finalInfo)

	// 实验二
	// 题目一
	// 散点图
	outputScatter(finalInfo)
	// 实验二
	// 题目二
	// 统计并打印直方图
	outputBar(finalInfo)

	// 实现二
	// 题目三
	// z-score归一化
	zscore(finalInfo)
}

// 从文本读取数据并合并
func readFile(filePath string, dbData []*model.Info) {
	// 读取文件
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败")
	}
	// 换行符分割，逐行读取
	lines := strings.Split(string(fileBytes), "\n")
	for n, value := range lines {
		// 第一行数据不读取
		if n == 0 {
			continue
		}
		var info *model.Info
		// 分割tab 读取每一个字符串
		data := strings.Split(value, "	")
		for i, v := range data {
			switch i {
			case 0:
				// 遍历数据库得到的数据
				for _, m := range dbData {
					if v == m.Name {
						info = m
					}
				}
			case 1:
				info.TID = v
			case 5:
				// 成绩取txt和db的平均值
				c1, _ := strconv.ParseFloat(v, 64)
				info.C1 = (info.C1 + c1) / 2
			case 6:
				// 成绩取txt和db的平均值
				c2, _ := strconv.ParseFloat(v, 64)
				info.C2 = (info.C1 + c2) / 2
			case 7:
				// 成绩取txt和db的平均值
				c3, _ := strconv.ParseFloat(v, 64)
				info.C3 = (info.C1 + c3) / 2
			case 8:
				// 成绩取txt和db的平均值
				c4, _ := strconv.ParseFloat(v, 64)
				info.C4 = (info.C1 + c4) / 2
			case 9:
				// 成绩取txt和db的平均值
				c5, _ := strconv.ParseFloat(v, 64)
				info.C5 = (info.C1 + c5) / 2
			case 10:
				// 成绩取txt和db的平均值
				c6, _ := strconv.ParseFloat(v, 64)
				info.C6 = (info.C1 + c6) / 2
			case 11:
				// 成绩取txt和db的平均值
				c7, _ := strconv.ParseFloat(v, 64)
				info.C7 = (info.C1 + c7) / 2
			case 12:
				// 成绩取txt和db的平均值
				c8, _ := strconv.ParseFloat(v, 64)
				info.C8 = (info.C1 + c8) / 2
			case 13:
				// 成绩取txt和db的平均值
				c9, _ := strconv.ParseFloat(v, 64)
				info.C9 = (info.C1 + c9) / 2
			case 14:
				// 成绩取txt和db的平均值
				c10, _ := strconv.ParseFloat(v, 64)
				info.C10 = (info.C1 + c10) / 2
			}
		}
	}
}

// 整合info到student中
func InteData(dbData []*model.Info) []Student {
	var finalInfo []Student
	for _, v := range dbData {
		item := Student{
			ID:           v.ID,
			TID:          v.TID,
			Name:         v.Name,
			City:         v.City,
			Gender:       v.Gender,
			Height:       v.Height,
			Constitution: v.Constitution,
			Score:        make(map[string]float64),
		}
		item.Score["c1"] = v.C1
		item.Score["c2"] = v.C2
		item.Score["c3"] = v.C3
		item.Score["c4"] = v.C4
		item.Score["c5"] = v.C5
		item.Score["c6"] = v.C6 * 10
		item.Score["c7"] = v.C7 * 10
		item.Score["c8"] = v.C8 * 10
		item.Score["c9"] = v.C9 * 10
		item.Score["c10"] = v.C10 * 10
		finalInfo = append(finalInfo, item)
	}
	return finalInfo
}

func analysis(finalInfo []Student) {
	// 学生中家乡在Beijing的所有课程的平均成绩
	var num float64
	var score float64
	// 计算在北京的总分数和成绩总数
	for _, v := range finalInfo {
		if v.City == "Beijing" {
			for _, v := range v.Score {
				num += 1
				score += v
			}
		}
	}
	average := score / num
	fmt.Printf("学生中家乡在Beijing的所有课程的平均成绩%v\n", average)

	// 学生中家乡在广州，课程1在80分以上，且课程9在9分以上的男同学的数量
	var num1 int
	for _, v := range finalInfo {
		if v.City == "Guangzhou" && v.Score["c1"] > 80 && v.Score["c9"]/9 > 9 && v.Gender == "boy" {
			num1++
		}
	}
	fmt.Printf("学生中家乡在广州，课程1在80分以上，且课程9在9分以上的男同学的数量为：%v个\n", num1)

	// 比较广州和上海两地女生的平均体能测试成绩，哪个地区的更强些？
	var numG float64
	var scoreG float64
	var numS float64
	var scoreS float64
	// 计算在广州和上海的总成绩和数量
	for _, v := range finalInfo {
		if v.City == "Guangzhou" && v.Gender == "girl" {
			numG += 1
			switch v.Constitution {
			case "bad":
				scoreG += 25
			case "general":
				scoreG += 50
			case "good":
				scoreG += 75
			case "excellent":
				scoreG += 100
			}
		}
		if v.City == "Shanghai" && v.Gender == "girl" {
			numS += 1
			switch v.Constitution {
			case "bad":
				scoreS += 25
			case "general":
				scoreS += 50
			case "good":
				scoreS += 75
			case "excellent":
				scoreS += 100
			}
		}
	}
	if scoreG/numG > scoreS/numS {
		fmt.Printf("广州的女生平均体能测试成绩更高，成绩为：%v\n", scoreG/numG)
	} else {
		fmt.Printf("上海的女生平均体能测试成绩更高，成绩为：%v\n", scoreS/numS)
	}

	// 学习成绩和体能测试成绩，两者的相关性是多少？
	fmt.Printf("c1与体育成绩的相关性为：%v\n", relevance(finalInfo, "c1"))
	fmt.Printf("c2与体育成绩的相关性为：%v\n", relevance(finalInfo, "c2"))
	fmt.Printf("c3与体育成绩的相关性为：%v\n", relevance(finalInfo, "c3"))
	fmt.Printf("c4与体育成绩的相关性为：%v\n", relevance(finalInfo, "c4"))
	fmt.Printf("c5与体育成绩的相关性为：%v\n", relevance(finalInfo, "c5"))
	fmt.Printf("c6与体育成绩的相关性为：%v\n", relevance(finalInfo, "c6"))
	fmt.Printf("c7与体育成绩的相关性为：%v\n", relevance(finalInfo, "c7"))
	fmt.Printf("c8与体育成绩的相关性为：%v\n", relevance(finalInfo, "c8"))
	fmt.Printf("c9与体育成绩的相关性为：%v\n", relevance(finalInfo, "c9"))
	fmt.Printf("c10与体育成绩的相关性为：%v\n", relevance(finalInfo, "c10"))
}

// 计算平均值
func mean(finalInfo []Student, tag string) float64 {
	// 计算体能成绩平均值
	if tag == "con" {
		var num float64
		var score float64
		for _, v := range finalInfo {
			num += 1
			switch v.Constitution {
			case "bad":
				score += 25
			case "general":
				score += 50
			case "good":
				score += 75
			case "excellent":
				score += 100
			}
		}
		return score / num
	} else { // 计算成绩平均值
		var num float64
		var score float64
		for _, v := range finalInfo {
			num += 1
			score += v.Score[tag]
		}
		return score / num
	}
}

// 计算标准差
func std(finalInfo []Student, tag string) float64 {
	average := mean(finalInfo, tag)
	// 计算体能成绩方差
	if tag == "con" {
		var num float64
		var score float64
		for _, v := range finalInfo {
			num += 1
			switch v.Constitution {
			case "bad":
				score += math.Pow(25-average, 2)
			case "general":
				score += math.Pow(50-average, 2)
			case "good":
				score += math.Pow(75-average, 2)
			case "excellent":
				score += math.Pow(100-average, 2)
			}
		}
		return math.Sqrt(score / num)
	} else { // 计算成绩方差
		var num float64
		var score float64
		for _, v := range finalInfo {
			num += 1
			score += math.Pow(v.Score[tag]-average, 2)
		}
		return math.Sqrt(score / num)
	}
}

// 计算相关性
func relevance(finalInfo []Student, tag string) float64 {
	var score float64
	for _, v := range finalInfo {
		switch v.Constitution {
		case "bad":
			score += ((25 - mean(finalInfo, "con")) / std(finalInfo, "con")) * ((v.Score[tag] - mean(finalInfo, tag)) / std(finalInfo, tag))
		case "general":
			score += ((50 - mean(finalInfo, "con")) / std(finalInfo, "con")) * ((v.Score[tag] - mean(finalInfo, tag)) / std(finalInfo, tag))
		case "good":
			score += ((75 - mean(finalInfo, "con")) / std(finalInfo, "con")) * ((v.Score[tag] - mean(finalInfo, tag)) / std(finalInfo, tag))
		case "excellent":
			score += ((100 - mean(finalInfo, "con")) / std(finalInfo, "con")) * ((v.Score[tag] - mean(finalInfo, tag)) / std(finalInfo, tag))
		}
	}
	return score

}

var (
	scatter3DColor = []string{
		"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
		"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
	}
)

// 统计并打印散点图
// 实验2 第1题
func outputScatter(finalInfo []Student) {
	items := make([]opts.Chart3DData, 0)
	for _, v := range finalInfo {
		switch v.Constitution {
		case "bad":
			items = append(items, opts.Chart3DData{Value: []interface{}{
				25,
				0,
				v.Score["c1"],
			}})
		case "general":
			items = append(items, opts.Chart3DData{Value: []interface{}{
				50,
				0,
				v.Score["c1"],
			}})
		case "good":
			items = append(items, opts.Chart3DData{Value: []interface{}{
				75,
				0,
				v.Score["c1"],
			}})
		case "excellent":
			items = append(items, opts.Chart3DData{Value: []interface{}{
				100,
				0,
				v.Score["c1"],
			}})
		}
	}
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "basic Scatter3D example"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        100,
			InRange:    &opts.VisualMapInRange{Color: scatter3DColor},
		}),
	)

	scatter3d.AddSeries("scatter3d", items)
	f, _ := os.Create("scatter.html")
	scatter3d.Render(f)
}

// 统计并打印直方图
// 实验2 第2题
func outputBar(finalInfo []Student) {
	c1 := make([]float64, 0)
	for _, v := range finalInfo {
		c1 = append(c1, v.Score["c1"])
	}
	// 计算每个区间的人数
	items := make([]opts.BarData, 0)
	for i := 0.0; i < 100; i += 5 {
		var num int
		for _, v := range c1 {
			if i == 0 {
				if i <= v && v <= i+5 {
					num++
				}
			} else {
				if i < v && v <= i+5 {
					num++
				}
			}
		}
		items = append(items, opts.BarData{Value: num})
	}
	nameItem := []string{"0-5", "6-10", "11-15", "16-20", "21-25", "26-30", "31-35", "36-40", "41-45", "46-50", "51-55", "56-60", "61-65", "66-70", "71-75", "76-80", "81-85", "86-90", "91-95", "96-100"}
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "成绩直方图"}),
	)
	bar.SetXAxis(nameItem).
		AddSeries("Category A", items)
	f, _ := os.Create("bar.html")
	bar.Render(f)
}

// z-score归一
func zscore(finalInfo []Student) {
	// 初始化二维切片
	array := make([][]float64, 10)
	for i := range array {
		array[i] = make([]float64, 0)
	}
	var tag string
	// 进行z-score归一化并存储在array中
	for i := 1; i < 11; i++ {
		tag = "c" + strconv.Itoa(i)
		for _, v := range finalInfo {
			array[i-1] = append(array[i-1], (v.Score[tag]-mean(finalInfo, tag))/std(finalInfo, tag))
		}
	}
	fmt.Println("输出归一化矩阵")
	for _, v := range array {
		fmt.Println(v)
	}
}
