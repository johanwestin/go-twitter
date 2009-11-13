//
// Copyright 2009 Bill Casarin <billcasarin@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package twitter

type Status interface {
  GetCreatedAt() string;
  GetCreatedAtInSeconds() int;
  GetFavorited() bool;
  GetId() int64;
  GetText() string;
  GetInReplyToScreenName() string;
  GetInReplyToStatusId() int64;
  GetInReplyToUserId() int64;
  GetNow() int;
}

// Our internal status struct
// the naming is odd so that
// json.Unmarshal can do its thing properly
type tTwitterStatus struct {
  Text string;
  Created_at string;
  CreatedAtSeconds int;
  Favorited bool;
  Id int64;
  In_reply_to_screen_name string;
  In_reply_to_status_id int64;
  In_reply_to_user_id int64;
  Error string;
  now int;
}

type tTwitterStatusDummy struct {
  Object tTwitterStatus;
}

type tTwitterTimelineDummy struct {
  Object []tTwitterStatus;
}

func (self *tTwitterStatus) GetCreatedAt() string {
  return self.Created_at;
}

func (self *tTwitterStatus) GetCreatedAtInSeconds() int {
  return self.CreatedAtSeconds;
}

func (self *tTwitterStatus) GetFavorited() bool {
  return self.Favorited;
}

func (self *tTwitterStatus) GetId() int64 {
  return self.Id;
}

func (self *tTwitterStatus) GetInReplyToScreenName() string {
  return self.In_reply_to_screen_name;
}

func (self *tTwitterStatus) GetText() string {
  return self.Text;
}

func (self *tTwitterStatus) GetInReplyToStatusId() int64 {
  return self.In_reply_to_status_id;
}

func (self *tTwitterStatus) GetInReplyToUserId() int64 {
  return self.In_reply_to_user_id;
}

func (self *tTwitterStatus) GetNow() int {
  return self.now;
}
