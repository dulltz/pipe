// Copyright 2020 The PipeCD Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubernetes

import (
	"fmt"
	"time"

	"github.com/kapetaniosci/pipe/pkg/app/piped/planner"
	"github.com/kapetaniosci/pipe/pkg/config"
	"github.com/kapetaniosci/pipe/pkg/model"
)

func buildProgressivePipeline(pp *config.DeploymentPipeline, now time.Time) []*model.PipelineStage {
	var stages []config.PipelineStage
	if pp != nil {
		stages = pp.Stages
	}

	var predefined bool
	if len(stages) == 0 {
		predefined = true
		stage, _ := planner.GetPredefinedStage(planner.PredefinedStageK8sUpdate)
		stages = []config.PipelineStage{stage}
	}

	out := make([]*model.PipelineStage, 0, len(stages))
	for i, s := range stages {
		id := s.Id
		if id == "" {
			id = fmt.Sprintf("stage-%d", i)
		}
		stage := &model.PipelineStage{
			Id:         id,
			Name:       s.Name.String(),
			Desc:       s.Desc,
			Index:      int32(i),
			Predefined: predefined,
			Status:     model.StageStatus_STAGE_NOT_STARTED_YET,
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		}
		out = append(out, stage)
	}

	return out
}

func buildPipeline(now time.Time) []*model.PipelineStage {
	stage, _ := planner.GetPredefinedStage(planner.PredefinedStageK8sUpdate)
	stages := []config.PipelineStage{stage}

	out := make([]*model.PipelineStage, 0, len(stages))
	for i, s := range stages {
		id := s.Id
		if id == "" {
			id = fmt.Sprintf("stage-%d", i)
		}
		stage := &model.PipelineStage{
			Id:         id,
			Name:       s.Name.String(),
			Desc:       s.Desc,
			Index:      int32(i),
			Predefined: true,
			Status:     model.StageStatus_STAGE_NOT_STARTED_YET,
			CreatedAt:  now.Unix(),
			UpdatedAt:  now.Unix(),
		}
		out = append(out, stage)
	}

	return out
}