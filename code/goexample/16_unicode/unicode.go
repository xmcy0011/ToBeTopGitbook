package main

func main() {
	const text = "第八届全国脊柱内镜大会暨中国医疗保健国际交流促进会骨科分会脊柱内镜学部2023年会暨河南省康复医学会骨科微创专委会成立大会"

	for runeStr := range text {
		println("r:", runeStr)
	}
}
