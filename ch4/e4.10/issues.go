// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//!+
func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	// data, err := json.MarshalIndent(result, "", "    ")
	// if err != nil {
	// 	log.Fatalf("JSON marshaling failed: %s", err)
	// }
	// fmt.Printf("%s\n", data)
	for _, item := range result.Items {
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			fmt.Printf("In a Month : %-35s #%-5d %9.9s %s\n",
				item.CreatedAt.String(), item.Number, item.User.Login, item.Title)
		}
	}
	for _, item := range result.Items {
		if !item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) && item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)) {
			fmt.Printf("In a Year  : %-35s #%-5d %9.9s %s\n",
				item.CreatedAt.String(), item.Number, item.User.Login, item.Title)
		}
	}
	for _, item := range result.Items {
		if item.CreatedAt.Before(time.Now().AddDate(-1, 0, 0)) {
			fmt.Printf("1+ Years   : %-35s #%-5d %9.9s %s\n",
				item.CreatedAt.String(), item.Number, item.User.Login, item.Title)
		}
	}
}

//!-

/*
//!+textoutput
https://docs.github.com/en/github/searching-for-information-on-github/searching-issues-and-pull-requests
>.\e4.10.exe repo:versatica/mediasoup is:issue is:open
14 issues:
In a Month : 2021-02-03 15:22:22 +0000 UTC       #512         ibc Implement AV1 codec
In a Month : 2021-02-01 02:51:55 +0000 UTC       #511   penguinol Check libwebrtc nack_module2.cc (add delay before sending NACK)
In a Year  : 2020-12-21 19:42:11 +0000 UTC       #492   LewisWolf SimulcastConsumer cannot switch layers if initial tsReferenceSpatialLayer disappears
In a Year  : 2020-11-20 10:31:29 +0000 UTC       #481     jmillan Implement RED
In a Year  : 2020-11-06 19:31:26 +0000 UTC       #477         ibc Implement Cryptex for RTP header extensions encryption
In a Year  : 2020-05-25 07:29:41 +0000 UTC       #408   penguinol There are mosaics on video when switching layers
In a Year  : 2020-05-11 09:47:52 +0000 UTC       #397         ibc Move to GitHub Actions?
In a Year  : 2020-04-15 08:18:57 +0000 UTC       #383         ibc Support End-to-End Encryption with WebRTC Insertable Streams
In a Year  : 2020-04-09 11:22:26 +0000 UTC       #381         ibc Add support for multichannel OPUS
1+ Years   : 2020-02-13 07:28:25 +0000 UTC       #359         ibc Drop gyp, move to CMake
1+ Years   : 2019-11-19 08:57:47 +0000 UTC       #347         ibc [Windows] Support for Visual Studio 2019
1+ Years   : 2019-11-18 17:01:00 +0000 UTC       #345      EyalSi [Windows] Parent project GUID  not found in "mediasoup-worker" project dependency section
1+ Years   : 2019-11-18 11:43:29 +0000 UTC       #344         ibc Write our own bandwidth estimator
1+ Years   : 2019-11-14 15:36:08 +0000 UTC       #343         ibc Integrate with Fuzzit
//!-textoutput
*/
