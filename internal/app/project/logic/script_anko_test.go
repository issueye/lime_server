package logic

import (
	"os"
	"testing"
)

func TestNewJsonFromFile(t *testing.T) {
	// 创建测试用的临时JSON文件
	testJSON := `{
        "string": "test",
        "number": 123,
        "boolean": true,
        "array": ["a", "b", "c"],
        "nested": {
            "key": "value"
        },
		"RT_VERSION": {
			"#1": {
			"0000": {
				"fixed": {
				"file_version": "1.0.0.1",
				"product_version": "1.0.0.1"
				},
				"info": {
					"0409": {
						"Comments": "青柠版本管理系统",
						"CompanyName": "杨桃桃",
						"FileDescription": "青柠版本管理系统",
						"FileVersion": "1.0.0.1",
						"InternalName": "青柠版本管理系统",
						"LegalCopyright": "杨桃桃 @ 2024",
						"LegalTrademarks": "",
						"OriginalFilename": "",
						"PrivateBuild": "",
						"ProductName": "杨桃桃",
						"ProductVersion": "v1.0.0-alpha.1",
						"SpecialBuild": ""
					}
				}
			}
			}
		}
    }`

	tmpfile, err := os.CreateTemp("", "test-*.json")
	if err != nil {
		t.Fatalf("无法创建临时文件: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(testJSON)); err != nil {
		t.Fatalf("无法写入测试数据: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("无法关闭临时文件: %v", err)
	}

	tests := []struct {
		name    string
		file    string
		wantErr bool
	}{
		{
			name:    "有效的JSON文件",
			file:    tmpfile.Name(),
			wantErr: false,
		},
		{
			name:    "不存在的文件",
			file:    "non_existent.json",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			json, err := NewJsonFromFile(tt.file)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewJsonFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// 验证JSON内容是否正确读取
				if got := json.GetString("string"); got != "test" {
					t.Errorf("GetString(string) = %v, want %v", got, "test")
				}
				if got := json.GetInt("number"); got != 123 {
					t.Errorf("GetInt(number) = %v, want %v", got, 123)
				}
				if got := json.GetBool("boolean"); got != true {
					t.Errorf("GetBool(boolean) = %v, want %v", got, true)
				}
				if got := json.GetStrings("array"); len(got) != 3 || got[0] != "a" {
					t.Errorf("GetStrings(array) = %v, want %v", got, []string{"a", "b", "c"})
				}
				if got := json.GetString("nested.key"); got != "value" {
					t.Errorf("GetString(nested.key) = %v, want %v", got, "value")
				}

				got := json.GetString("RT_VERSION.#1.0000.info.0409.ProductVersion")
				t.Log("ProductVersion", got)
				if got != "v1.0.0-alpha.1" {
					t.Errorf("GetString(RT_VERSION.#1.0000.info.0409.ProductVersion) = %v, want %v", got, "v1.0.0-alpha.1")
				}
			}
		})
	}
}

func TestAnkoJsonIntegration(t *testing.T) {
	anko := NewAnko()

	script := `
        json = Json('{"name": "test", "nums": [1,2,3]}')
        name = json.GetString("name")
        nums = json.GetInts("nums")
		println("nums", nums)
		println("nums[0]", nums[0])
        return name + "," + nums[0]
    `

	result, err := anko.Execute(script)
	if err != nil {
		t.Fatalf("执行脚本失败: %v", err)
	}

	expected := "test,1"
	if result != expected {
		t.Errorf("期望得到 %s, 实际得到 %v", expected, result)
	}
}
