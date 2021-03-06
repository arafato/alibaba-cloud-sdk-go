package facebody

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Data is a nested struct in facebody response
type Data struct {
	LandmarkCount       int       `json:"LandmarkCount" xml:"LandmarkCount"`
	Confidence          float64   `json:"Confidence" xml:"Confidence"`
	DenseFeatureLength  int       `json:"DenseFeatureLength" xml:"DenseFeatureLength"`
	Mask                int       `json:"Mask" xml:"Mask"`
	FaceCount           int       `json:"FaceCount" xml:"FaceCount"`
	FaceProbability     float64   `json:"FaceProbability" xml:"FaceProbability"`
	FaceProbabilityList []float64 `json:"FaceProbabilityList" xml:"FaceProbabilityList"`
	GenderList          []int     `json:"GenderList" xml:"GenderList"`
	FaceRectangles      []int     `json:"FaceRectangles" xml:"FaceRectangles"`
	PoseList            []float64 `json:"PoseList" xml:"PoseList"`
	Pupils              []float64 `json:"Pupils" xml:"Pupils"`
	Glasses             []int     `json:"Glasses" xml:"Glasses"`
	Expressions         []int     `json:"Expressions" xml:"Expressions"`
	RectAList           []int     `json:"RectAList" xml:"RectAList"`
	RectBList           []int     `json:"RectBList" xml:"RectBList"`
	AgeList             []int     `json:"AgeList" xml:"AgeList"`
	Thresholds          []float64 `json:"Thresholds" xml:"Thresholds"`
	DenseFeatures       []string  `json:"DenseFeatures" xml:"DenseFeatures"`
	Landmarks           []float64 `json:"Landmarks" xml:"Landmarks"`
}
