// Copyright (C) 2016 The Android Open Source Project
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
package mediaextentions

import (
  "path/filepath"

  "android/soong/android"
  "android/soong/cc"

  "github.com/google/blueprint"
)

func init() {
  android.RegisterModuleType("mediaextentions_defaults", mediaextentionsDefaultsFactory)
}

func mediaextentionsDefaultsFactory() (blueprint.Module, []interface{}) {
  module, props := cc.DefaultsFactory()
  android.AddLoadHook(module, mediaextentionsDefaults)

  return module, props
}

func mediaextentionsDefaults(ctx android.LoadHookContext) {
  type props struct {
    Include_dirs []string
    Cflags []string
  }

  p := &props{}
  p.Cflags, p.Include_dirs = globalDefaults(ctx)

  ctx.AppendProperties(p)
}

func globalDefaults(ctx android.BaseContext) ([]string, []string) {
  var cflags []string
  var includeDirs []string

  qcom_media_dir := ctx.DeviceConfig().QTIMediaPath()
  qcom_media_include := filepath.Join(qcom_media_dir, "/mm-core/inc")
  includeDirs = append(includeDirs, qcom_media_include)

  return cflags, includeDirs
}
