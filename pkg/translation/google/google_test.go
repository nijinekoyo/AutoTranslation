/*
 * @Author: nijineko
 * @Date: 2025-07-03 16:41:28
 * @LastEditTime: 2025-07-03 17:08:17
 * @LastEditors: nijineko
 * @Description: Google翻译实现
 * @FilePath: \AutoTranslation\pkg\translation\google\google_test.go
 */

package google

import "testing"

func TestGoogleTranslator_TranslateText(t *testing.T) {
	type args struct {
		Text           string
		SourceLanguage *string
		TargetLang     string
	}
	tests := []struct {
		name    string
		g       *GoogleTranslator
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Translate English to Chinese",
			g:    New(),
			args: args{
				Text:           "Hello, world!",
				SourceLanguage: nil,
				TargetLang:     "zh-CN",
			},
			want:    "你好世界！",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.TranslateText(tt.args.Text, tt.args.SourceLanguage, tt.args.TargetLang)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoogleTranslator.TranslateText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GoogleTranslator.TranslateText() = %v, want %v", got, tt.want)
			}
		})
	}
}
