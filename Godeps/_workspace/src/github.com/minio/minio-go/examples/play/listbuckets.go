// +build ignore

/*
 * Minimal object storage library (C) 2015 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"

	"github.com/minio/minio-go"
)

func main() {
	config := minio.Config{
		AccessKeyID:     "",
		SecretAccessKey: "",
		Endpoint:        "https://play.minio.io:9000",
	}
	playClient, err := minio.New(config)
	if err != nil {
		log.Fatalln(err)
	}
	for bucket := range playClient.ListBuckets() {
		if bucket.Err != nil {
			log.Fatalln(bucket.Err)
		}
		log.Println(bucket.Stat)
	}
}
