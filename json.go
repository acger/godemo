package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

func main() {
	jsonstr := `
		{
		  "took" : 7,
		  "timed_out" : false,
		  "_shards" : {
			"total" : 5,
			"successful" : 5,
			"skipped" : 0,
			"failed" : 0
		  },
		  "hits" : {
			"total" : {
			  "value" : 2,
			  "relation" : "eq"
			},
			"max_score" : 1.4384104,
			"hits" : [
			  {
				"_index" : "acger_pair_0",
				"_type" : "_doc",
				"_id" : "1",
				"_score" : 1.4384104,
				"_source" : {
				  "skill" : "画画 砍价 么么哒 秒睡",
				  "skill_need" : "画画 天才 力量大 团结友爱",
				  "uid" : 5,
				  "boost" : 1,
				  "star" : 0,
				  "created_at" : "2022-02-24T13:34:25.173+00:00",
				  "updated_at" : "2022-02-24T13:34:25.173+00:00",
				  "deleted_at" : null
				},
				"highlight" : {
				  "skill" : [
					"<b>画画</b> 砍价 么么哒 秒睡"
				  ],
				  "skill_need" : [
					"<b>画画</b> 天才 力量大 团结友爱"
				  ]
				}
			  },
			  {
				"_index" : "acger_pair_0",
				"_type" : "_doc",
				"_id" : "3",
				"_score" : 0.2876821,
				"_source" : {
				  "skill" : "啥都没有 画画",
				  "skill_need" : "是个人 是只猫 是个蠢驴 画画",
				  "uid" : 3,
				  "boost" : 0,
				  "star" : 0,
				  "created_at" : "2022-02-24T13:37:04.104+00:00",
				  "updated_at" : "2022-02-24T16:23:55.641+00:00",
				  "deleted_at" : null
				},
				"highlight" : {
				  "skill" : [
					"啥都没有 <b>画画</b>"
				  ],
				  "skill_need" : [
					"是个人 是只猫 是个蠢驴 <b>画画</b>"
				  ]
				}
			  }
			]
		  }
		}
`


	type Shards struct {
		Total int `json:"total"`
		Successful int `json:"successful"`
		Skipped int `json:"skipped"`
		Failed int `json:"failed"`
	}
	type Total struct {
		Value int `json:"value"`
		Relation string `json:"relation"`
	}
	type Source struct {
		Skill string `json:"skill"`
		SkillNeed string `json:"skill_need"`
		UID int `json:"uid"`
		Boost int `json:"boost"`
		Star int `json:"star"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt interface{} `json:"deleted_at"`
	}
	type Hits struct {
		Index string `json:"_index"`
		Type string `json:"_type"`
		ID string `json:"_id"`
		Score float64 `json:"_score"`
		Source Source `json:"_source"`
	}

	type Hit struct {
		Total Total `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits []Hits `json:"hits"`
	}

	type AutoGenerated struct {
		Took int `json:"took"`
		TimedOut bool `json:"timed_out"`
		Shards Shards `json:"_shards"`
		Hits Hit `json:"hits"`
	}

	a := AutoGenerated{}

	jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(jsonstr), &a)

	fmt.Print(a.Hits.Hits[0].Source.SkillNeed)
}
