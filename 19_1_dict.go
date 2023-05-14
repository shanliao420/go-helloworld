package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type DictYouDao struct {
	WebTrans struct {
		WebTranslation []struct {
			Same      string `json:"@same,omitempty"`
			Key       string `json:"key"`
			KeySpeech string `json:"key-speech"`
			Trans     []struct {
				Summary struct {
					Line []string `json:"line"`
				} `json:"summary"`
				Value   string `json:"value"`
				Support int    `json:"support"`
				URL     string `json:"url"`
			} `json:"trans"`
		} `json:"web-translation"`
	} `json:"web_trans"`
	OxfordAdvanceHTML struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvanceHtml"`
	PicDict struct {
		Pic []struct {
			Image string `json:"image"`
			Host  string `json:"host"`
			URL   string `json:"url"`
		} `json:"pic"`
	} `json:"pic_dict"`
	VideoSents struct {
		SentsData []struct {
			VideoCover  string `json:"video_cover"`
			Contributor string `json:"contributor"`
			SubtitleSrt string `json:"subtitle_srt"`
			ID          int    `json:"id"`
			Video       string `json:"video"`
		} `json:"sents_data"`
		WordInfo struct {
			ReturnPhrase string   `json:"return-phrase"`
			Sense        []string `json:"sense"`
		} `json:"word_info"`
	} `json:"video_sents"`
	Simple struct {
		Query string `json:"query"`
		Word  []struct {
			Usphone      string `json:"usphone"`
			Ukphone      string `json:"ukphone"`
			Ukspeech     string `json:"ukspeech"`
			ReturnPhrase string `json:"return-phrase"`
			Usspeech     string `json:"usspeech"`
		} `json:"word"`
	} `json:"simple"`
	Phrs struct {
		Word string `json:"word"`
		Phrs []struct {
			Headword    string `json:"headword"`
			Translation string `json:"translation"`
		} `json:"phrs"`
	} `json:"phrs"`
	Oxford struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxford"`
	Syno struct {
		Synos []struct {
			Pos  string   `json:"pos"`
			Ws   []string `json:"ws"`
			Tran string   `json:"tran"`
		} `json:"synos"`
		Word string `json:"word"`
	} `json:"syno"`
	Collins struct {
		CollinsEntries []struct {
			Entries struct {
				Entry []struct {
					TranEntry []struct {
						PosEntry struct {
							Pos     string `json:"pos"`
							PosTips string `json:"pos_tips"`
						} `json:"pos_entry"`
						ExamSents struct {
							Sent []struct {
								ChnSent string `json:"chn_sent"`
								EngSent string `json:"eng_sent"`
							} `json:"sent"`
						} `json:"exam_sents"`
						Tran string `json:"tran"`
					} `json:"tran_entry"`
				} `json:"entry"`
			} `json:"entries"`
			Phonetic     string `json:"phonetic"`
			BasicEntries struct {
				BasicEntry []struct {
					Cet      string `json:"cet"`
					Headword string `json:"headword"`
				} `json:"basic_entry"`
			} `json:"basic_entries"`
			Headword string `json:"headword"`
			Star     string `json:"star"`
		} `json:"collins_entries"`
	} `json:"collins"`
	WordVideo struct {
		WordVideos []struct {
			Ad struct {
				Avatar string `json:"avatar"`
				Title  string `json:"title"`
				URL    string `json:"url"`
			} `json:"ad"`
			Video struct {
				Cover string `json:"cover"`
				Image string `json:"image"`
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"video"`
		} `json:"word_videos"`
	} `json:"word_video"`
	Webster struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"webster"`
	Discriminate struct {
		Data []struct {
			Source string `json:"source"`
			Usages []struct {
				Headword string `json:"headword"`
				Usage    string `json:"usage"`
			} `json:"usages"`
			Headwords []string `json:"headwords"`
			Tran      string   `json:"tran"`
		} `json:"data"`
		ReturnPhrase string `json:"return-phrase"`
	} `json:"discriminate"`
	WikipediaDigest struct {
		Summarys []struct {
			Summary string `json:"summary"`
			Key     string `json:"key"`
		} `json:"summarys"`
		Source struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
	} `json:"wikipedia_digest"`
	Lang string `json:"lang"`
	Ec   struct {
		WebTrans []string `json:"web_trans"`
		Special  []struct {
			Nat   string `json:"nat"`
			Major string `json:"major"`
		} `json:"special"`
		ExamType []string `json:"exam_type"`
		Source   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
		Word struct {
			Usphone  string `json:"usphone"`
			Ukphone  string `json:"ukphone"`
			Ukspeech string `json:"ukspeech"`
			Trs      []struct {
				Pos  string `json:"pos"`
				Tran string `json:"tran"`
			} `json:"trs"`
			ReturnPhrase string `json:"return-phrase"`
			Usspeech     string `json:"usspeech"`
		} `json:"word"`
	} `json:"ec"`
	Ee struct {
		Source struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
		Word struct {
			Trs []struct {
				Pos string `json:"pos"`
				Tr  []struct {
					Tran         string   `json:"tran"`
					SimilarWords []string `json:"similar-words,omitempty"`
					Examples     []string `json:"examples,omitempty"`
				} `json:"tr"`
			} `json:"trs"`
			Speech       string `json:"speech"`
			ReturnPhrase string `json:"return-phrase"`
		} `json:"word"`
	} `json:"ee"`
	BlngSentsPart struct {
		SentenceCount int `json:"sentence-count"`
		SentencePair  []struct {
			Sentence            string `json:"sentence"`
			SentenceEng         string `json:"sentence-eng"`
			SentenceTranslation string `json:"sentence-translation"`
			SpeechSize          string `json:"speech-size"`
			AlignedWords        struct {
				Src struct {
					Chars []struct {
						S      string `json:"@s"`
						E      string `json:"@e"`
						Aligns struct {
							Sc []struct {
								ID string `json:"@id"`
							} `json:"sc"`
							Tc []struct {
								ID string `json:"@id"`
							} `json:"tc"`
						} `json:"aligns"`
						ID string `json:"@id"`
					} `json:"chars"`
				} `json:"src"`
				Tran struct {
					Chars []struct {
						S      string `json:"@s"`
						E      string `json:"@e"`
						Aligns struct {
							Sc []struct {
								ID string `json:"@id"`
							} `json:"sc"`
							Tc []struct {
								ID string `json:"@id"`
							} `json:"tc"`
						} `json:"aligns"`
						ID string `json:"@id"`
					} `json:"chars"`
				} `json:"tran"`
			} `json:"aligned-words"`
			Source         string `json:"source"`
			URL            string `json:"url"`
			SentenceSpeech string `json:"sentence-speech"`
		} `json:"sentence-pair"`
		More        string `json:"more"`
		TrsClassify []struct {
			Proportion string `json:"proportion"`
			Tr         string `json:"tr"`
		} `json:"trs-classify"`
	} `json:"blng_sents_part"`
	Individual struct {
		Trs []struct {
			Pos  string `json:"pos"`
			Tran string `json:"tran"`
		} `json:"trs"`
		Idiomatic []struct {
			Colloc struct {
				En string `json:"en"`
				Zh string `json:"zh"`
			} `json:"colloc"`
		} `json:"idiomatic"`
		Level    string `json:"level"`
		ExamInfo struct {
			Year             int `json:"year"`
			QuestionTypeInfo []struct {
				Time int    `json:"time"`
				Type string `json:"type"`
			} `json:"questionTypeInfo"`
			RecommendationRate int `json:"recommendationRate"`
			Frequency          int `json:"frequency"`
		} `json:"examInfo"`
		ReturnPhrase  string `json:"return-phrase"`
		PastExamSents []struct {
			En     string `json:"en"`
			Source string `json:"source"`
			Zh     string `json:"zh"`
		} `json:"pastExamSents"`
	} `json:"individual"`
	CollinsPrimary struct {
		Words struct {
			Indexforms []string `json:"indexforms"`
			Word       string   `json:"word"`
		} `json:"words"`
		Gramcat []struct {
			Audiourl      string `json:"audiourl"`
			Pronunciation string `json:"pronunciation"`
			Senses        []struct {
				Sensenumber string `json:"sensenumber"`
				Examples    []struct {
					Sense struct {
						Lang string `json:"lang"`
						Word string `json:"word"`
					} `json:"sense"`
					Example string `json:"example"`
				} `json:"examples"`
				Definition string `json:"definition"`
				Lang       string `json:"lang"`
				Word       string `json:"word"`
			} `json:"senses"`
			Partofspeech string `json:"partofspeech"`
			Audio        string `json:"audio"`
			Forms        []struct {
				Form string `json:"form"`
			} `json:"forms"`
		} `json:"gramcat"`
	} `json:"collins_primary"`
	RelWord struct {
		Word string `json:"word"`
		Stem string `json:"stem"`
		Rels []struct {
			Rel struct {
				Pos   string `json:"pos"`
				Words []struct {
					Word string `json:"word"`
					Tran string `json:"tran"`
				} `json:"words"`
			} `json:"rel"`
		} `json:"rels"`
	} `json:"rel_word"`
	AuthSentsPart struct {
		SentenceCount int    `json:"sentence-count"`
		More          string `json:"more"`
		Sent          []struct {
			Score      float64 `json:"score"`
			Speech     string  `json:"speech"`
			SpeechSize string  `json:"speech-size"`
			Source     string  `json:"source"`
			URL        string  `json:"url"`
			Foreign    string  `json:"foreign"`
		} `json:"sent"`
	} `json:"auth_sents_part"`
	MediaSentsPart struct {
		SentenceCount int    `json:"sentence-count"`
		More          string `json:"more"`
		Query         string `json:"query"`
		Sent          []struct {
			Mediatype string `json:"@mediatype"`
			Snippets  struct {
				Snippet []struct {
					StreamURL string `json:"streamUrl"`
					Duration  string `json:"duration"`
					Swf       string `json:"swf"`
					Name      string `json:"name"`
					Source    string `json:"source"`
					Win8      string `json:"win8"`
				} `json:"snippet"`
			} `json:"snippets"`
			SpeechSize string `json:"speech-size,omitempty"`
			Eng        string `json:"eng"`
			Chn        string `json:"chn,omitempty"`
		} `json:"sent"`
	} `json:"media_sents_part"`
	ExpandEc struct {
		ReturnPhrase string `json:"return-phrase"`
		Source       struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"source"`
		Word []struct {
			TransList []struct {
				Content struct {
					DetailPos string `json:"detailPos"`
					ExamType  []struct {
						Zh string `json:"zh"`
						En string `json:"en,omitempty"`
					} `json:"examType"`
					Sents []struct {
						SentOrig   string `json:"sentOrig"`
						SourceType string `json:"sourceType"`
						SentSpeech string `json:"sentSpeech"`
						SentTrans  string `json:"sentTrans"`
						Source     string `json:"source"`
						Type       string `json:"type,omitempty"`
					} `json:"sents"`
				} `json:"content"`
				Trans string `json:"trans"`
			} `json:"transList"`
			Pos string `json:"pos"`
			Wfs []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"wfs,omitempty"`
		} `json:"word"`
	} `json:"expand_ec"`
	Etym struct {
		Etyms struct {
			En []struct {
				Source string `json:"source"`
				Word   string `json:"word"`
				Value  string `json:"value"`
				URL    string `json:"url"`
				Desc   string `json:"desc"`
			} `json:"en"`
			Zh []struct {
				Source string `json:"source"`
				Word   string `json:"word"`
				Value  string `json:"value"`
				URL    string `json:"url"`
				Desc   string `json:"desc"`
			} `json:"zh"`
		} `json:"etyms"`
		Word string `json:"word"`
	} `json:"etym"`
	Special struct {
		Summary struct {
			Sources struct {
				Source struct {
					Site string `json:"site"`
					URL  string `json:"url"`
				} `json:"source"`
			} `json:"sources"`
			Text string `json:"text"`
		} `json:"summary"`
		CoAdd   string `json:"co-add"`
		Total   string `json:"total"`
		Entries []struct {
			Entry struct {
				Major string `json:"major"`
				Trs   []struct {
					Tr struct {
						Nat      string `json:"nat"`
						ChnSent  string `json:"chnSent"`
						Cite     string `json:"cite"`
						DocTitle string `json:"docTitle"`
						EngSent  string `json:"engSent"`
						URL      string `json:"url"`
					} `json:"tr"`
				} `json:"trs"`
				Num int `json:"num"`
			} `json:"entry"`
		} `json:"entries"`
	} `json:"special"`
	Senior struct {
		EncryptedData string `json:"encryptedData"`
		Source        struct {
			Name string `json:"name"`
		} `json:"source"`
	} `json:"senior"`
	Input string `json:"input"`
	Meta  struct {
		Input           string   `json:"input"`
		GuessLanguage   string   `json:"guessLanguage"`
		IsHasSimpleDict string   `json:"isHasSimpleDict"`
		Le              string   `json:"le"`
		Lang            string   `json:"lang"`
		Dicts           []string `json:"dicts"`
	} `json:"meta"`
	Le            string `json:"le"`
	OxfordAdvance struct {
		EncryptedData string `json:"encryptedData"`
	} `json:"oxfordAdvance"`
}

type DictRequestCaiYun struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResponseCaiYun struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []interface{} `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func requestDictCaiYun(client *http.Client, targetWord string) {
	request := DictRequestCaiYun{TransType: "en2zh", Source: targetWord}
	buff, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buff)

	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "c04977ea8bd6b9302d3e3df5680c1dec")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body:", string(bodyText))
	}

	// fmt.Printf("%s\n", bodyText)

	var dictResponse DictResponseCaiYun
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(dictResponse)
	fmt.Println("--------------------------彩云-------------------------------")
	fmt.Println(dictResponse.Dictionary.Entry, "\tUS:", dictResponse.Dictionary.Prons.EnUs, "EN:", dictResponse.Dictionary.Prons.En)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
	for _, item := range dictResponse.Dictionary.WqxExample {
		fmt.Println(item[0], "\n", item[1])
	}
	fmt.Println("------------------------------------------------------------")
}

func requestDictYouDao(client *http.Client, targetWord string) {
	var data = strings.NewReader("q=" + targetWord + "&le=en&t=2&client=web&sign=40bae1f6a677590420afa01d90691444&keyfrom=webdict")
	req, err := http.NewRequest("POST", "https://dict.youdao.com/jsonapi_s?doctype=json&jsonversion=4", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "OUTFOX_SEARCH_USER_ID=-1793131797@10.109.66.219; OUTFOX_SEARCH_USER_ID_NCOO=50452448.372467056")
	req.Header.Set("Origin", "https://www.youdao.com")
	req.Header.Set("Referer", "https://www.youdao.com/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", bodyText)

	var dictYouDao DictYouDao
	err = json.Unmarshal(bodyText, &dictYouDao)
	if err != nil {
		return
	}

	// for _, item := range dictYouDao.Individual.Trs {
	//  	fmt.Println(item)
	// }

	phone := dictYouDao.Simple.Word[0]
	fmt.Println("--------------------------有道-------------------------------")

	fmt.Println(dictYouDao.Simple.Query, "\nUS:/", phone.Usphone, "/\tUK:/", phone.Ukphone, "/")
	for _, item := range dictYouDao.ExpandEc.Word {
		fmt.Print(item.Pos)
		for _, it := range item.TransList {
			fmt.Print(it.Trans)
			fmt.Print("; ")
		}
		fmt.Println()
	}
	fmt.Println("------------------------------------------------------------")

}

func main() {
	client := &http.Client{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		var query string
		if scanner.Scan() {
			query = scanner.Text()
		}
		var waitGroup sync.WaitGroup
		waitGroup.Add(2)
		go func() {
			defer waitGroup.Done()
			requestDictYouDao(client, query)
		}()
		go func() {
			defer waitGroup.Done()
			requestDictCaiYun(client, query)
		}()
		waitGroup.Wait()
	}

}
